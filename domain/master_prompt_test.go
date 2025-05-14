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
