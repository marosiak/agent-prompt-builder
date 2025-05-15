package domain

import "testing"

func TestMasterPrompt_UpdateValueByID(t *testing.T) {
	mp := &MasterPrompt{
		StylePreset: StylePreset{
			Values: []Style{
				{ID: "1", Name: "OldName", Weight: 10},
			},
		},
		RulePreset: RulePreset{
			Values: []Rule{
				{ID: "2", Name: "OldRule", Weight: 20},
			},
		},
		TeamPreset: TeamPreset{
			Values: []Person{
				{ID: "3", Name: "OldTeam", Features: []Feature{{
					ID:     "4",
					Name:   "OldFeature",
					Weight: 30,
				}}},
			},
		},
	}

	newName := "NewName"
	newWeight := 99

	// Test StylePreset
	mp.UpdateValueByID("1", &newName, &newWeight)
	if mp.StylePreset.Values[0].Name != newName {
		t.Errorf("expected Name %s, got %s", newName, mp.StylePreset.Values[0].Name)
	}
	if mp.StylePreset.Values[0].Weight != newWeight {
		t.Errorf("expected Weight %d, got %d", newWeight, mp.StylePreset.Values[0].Weight)
	}

	// Test RulePreset
	mp.UpdateValueByID("2", &newName, &newWeight)
	if mp.RulePreset.Values[0].Name != newName {
		t.Errorf("expected Name %s, got %s", newName, mp.RulePreset.Values[0].Name)
	}
	if mp.RulePreset.Values[0].Weight != newWeight {
		t.Errorf("expected Weight %d, got %d", newWeight, mp.RulePreset.Values[0].Weight)
	}

	// Test TeamPreset
	mp.UpdateValueByID("3", &newName, &newWeight)
	if mp.TeamPreset.Values[0].Name != newName {
		t.Errorf("expected Name %s, got %s", newName, mp.TeamPreset.Values[0].Name)
	}

	mp.UpdateValueByID("4", &newName, &newWeight)
	if mp.TeamPreset.Values[0].Features[0].Name != newName {
		t.Errorf("expected Name %s, got %s", newName, mp.TeamPreset.Values[0].Name)
	}
	if mp.TeamPreset.Values[0].Features[0].Weight != newWeight {
		t.Errorf("expected Feature Weight %d, got %d", newWeight, mp.TeamPreset.Values[0].Features[0].Weight)
	}
}

func TestMasterPrompt_RemoveFeatureByID(t *testing.T) {
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "p1",
					Name: "Person1",
					Features: []Feature{
						{ID: "p1_1", Name: "Feature1", Weight: 10},
						{ID: "p1_2", Name: "Feature1", Weight: 10},
					},
				},
				{
					ID:   "p2",
					Name: "Person2",
					Features: []Feature{
						{ID: "p2_1", Name: "Feature1", Weight: 10},
						{ID: "p2_2", Name: "Feature1", Weight: 10},
					},
				},
			},
		},
		StylePreset: StylePreset{
			Values: []Style{
				{ID: "2", Name: "Style1", Weight: 10},
				{ID: "3", Name: "Style1", Weight: 10},
				{ID: "4", Name: "", Weight: 100},
			},
		},
		RulePreset: RulePreset{
			Values: []Rule{
				{ID: "5", Name: "Rule1", Weight: 10},
				{ID: "6", Name: "Rule1", Weight: 10},
				{ID: "7", Name: "", Weight: 100},
			},
		},
	}

	// test removing all features one by one and check if they're being removed for real, ommit removing person it's in different test

	mp.RemoveFeatureByID("p1_1")
	if len(mp.TeamPreset.Values[0].Features) != 1 {
		t.Errorf("expected 1 feature, got %d", len(mp.TeamPreset.Values[0].Features))
	}

	mp.RemoveFeatureByID("p1_2")
	if len(mp.TeamPreset.Values[0].Features) != 0 {
		t.Errorf("expected 0 feature, got %d", len(mp.TeamPreset.Values[0].Features))
	}

	mp.RemoveFeatureByID("p2_1")
	if len(mp.TeamPreset.Values[1].Features) != 1 {
		t.Errorf("expected 1 feature, got %d", len(mp.TeamPreset.Values[1].Features))
	}

	mp.RemoveFeatureByID("p2_2")
	if len(mp.TeamPreset.Values[1].Features) != 0 {
		t.Errorf("expected 0 feature, got %d", len(mp.TeamPreset.Values[1].Features))
	}

	// test removing style

	mp.RemoveFeatureByID("2")
	if len(mp.StylePreset.Values) != 2 {
		t.Errorf("expected 2 styles, got %d", len(mp.StylePreset.Values))
	}

	mp.RemoveFeatureByID("3")
	if len(mp.StylePreset.Values) != 1 {
		t.Errorf("expected 1 style, got %d", len(mp.StylePreset.Values))
	}

	mp.RemoveFeatureByID("4")
	if len(mp.StylePreset.Values) != 0 {
		t.Errorf("expected 0 styles, got %d", len(mp.StylePreset.Values))
	}

	// test removing rule
	mp.RemoveFeatureByID("5")
	if len(mp.RulePreset.Values) != 2 {
		t.Errorf("expected 2 rules, got %d", len(mp.RulePreset.Values))
	}

	mp.RemoveFeatureByID("6")
	if len(mp.RulePreset.Values) != 1 {
		t.Errorf("expected 1 rule, got %d", len(mp.RulePreset.Values))
	}
	mp.RemoveFeatureByID("7")
	if len(mp.RulePreset.Values) != 0 {
		t.Errorf("expected 0 rules, got %d", len(mp.RulePreset.Values))
	}
}

func TestMasterPrompt_RemoveTeamMemberByID(t *testing.T) {
	// Setup a master prompt with multiple team members
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "p1",
					Name: "Person1",
					Features: []Feature{
						{ID: "f1", Name: "Feature1", Weight: 10},
					},
				},
				{
					ID:   "p2",
					Name: "Person2",
					Features: []Feature{
						{ID: "f2", Name: "Feature2", Weight: 20},
					},
				},
				{
					ID:   "p3",
					Name: "Person3",
					Features: []Feature{
						{ID: "f3", Name: "Feature3", Weight: 30},
					},
				},
			},
		},
	}

	// Test case 1: Remove a team member that exists
	originalCount := len(mp.TeamPreset.Values)
	mp.RemoveTeamMemberByID("p2")
	
	// Check if team member was removed
	if len(mp.TeamPreset.Values) != originalCount-1 {
		t.Errorf("expected %d team members, got %d", originalCount-1, len(mp.TeamPreset.Values))
	}

	// Check if correct team member was removed
	for _, person := range mp.TeamPreset.Values {
		if person.ID == "p2" {
			t.Errorf("expected team member with ID 'p2' to be removed, but it still exists")
		}
	}

	// Test case 2: Remove a team member that doesn't exist
	currentCount := len(mp.TeamPreset.Values)
	mp.RemoveTeamMemberByID("non-existent-id")
	
	// Check that no team member was removed
	if len(mp.TeamPreset.Values) != currentCount {
		t.Errorf("expected %d team members, got %d (nothing should be removed for non-existent ID)", 
		         currentCount, len(mp.TeamPreset.Values))
	}

	// Test case 3: Remove all remaining team members
	mp.RemoveTeamMemberByID("p1")
	mp.RemoveTeamMemberByID("p3")
	
	// Check if team is empty
	if len(mp.TeamPreset.Values) != 0 {
		t.Errorf("expected 0 team members, got %d", len(mp.TeamPreset.Values))
	}

	// Test case 4: Try to remove from an empty team
	mp.RemoveTeamMemberByID("p1")
	
	// Should still be empty without errors
	if len(mp.TeamPreset.Values) != 0 {
		t.Errorf("expected 0 team members, got %d", len(mp.TeamPreset.Values))
	}
}
