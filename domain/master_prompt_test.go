package domain

import (
	"encoding/json"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestMasterPrompt_UpdateValueByID(t *testing.T) {
	// Setup test data with descriptive IDs
	mp := &MasterPrompt{
		StylePreset: StylePreset{
			Values: []Style{
				{ID: "style_id", Name: "OldStyleName", Weight: 10},
			},
		},
		RulePreset: RulePreset{
			Values: []Rule{
				{ID: "rule_id", Name: "OldRuleName", Weight: 20},
			},
		},
		TeamPreset: TeamPreset{
			Values: []Person{
				{ID: "person_id", Name: "OldPersonName", Features: []Feature{{
					ID:     "feature_id",
					Name:   "OldFeatureName",
					Weight: 30,
				}}},
			},
		},
	}

	updatedName := "UpdatedName"
	updatedWeight := 99

	// Test 1: Update Style
	mp.UpdateValueByID("style_id", &updatedName, &updatedWeight)
	if mp.StylePreset.Values[0].Name != updatedName {
		t.Errorf("Style name - expected: %s, got: %s", updatedName, mp.StylePreset.Values[0].Name)
	}
	if mp.StylePreset.Values[0].Weight != updatedWeight {
		t.Errorf("Style weight - expected: %d, got: %d", updatedWeight, mp.StylePreset.Values[0].Weight)
	}

	// Test 2: Update Rule
	mp.UpdateValueByID("rule_id", &updatedName, &updatedWeight)
	if mp.RulePreset.Values[0].Name != updatedName {
		t.Errorf("Rule name - expected: %s, got: %s", updatedName, mp.RulePreset.Values[0].Name)
	}
	if mp.RulePreset.Values[0].Weight != updatedWeight {
		t.Errorf("Rule weight - expected: %d, got: %d", updatedWeight, mp.RulePreset.Values[0].Weight)
	}

	// Test 3: Update Person
	mp.UpdateValueByID("person_id", &updatedName, &updatedWeight)
	if mp.TeamPreset.Values[0].Name != updatedName {
		t.Errorf("Person name - expected: %s, got: %s", updatedName, mp.TeamPreset.Values[0].Name)
	}

	// Test 4: Update Feature
	mp.UpdateValueByID("feature_id", &updatedName, &updatedWeight)
	if mp.TeamPreset.Values[0].Features[0].Name != updatedName {
		t.Errorf("Feature name - expected: %s, got: %s", updatedName, mp.TeamPreset.Values[0].Features[0].Name)
	}
	if mp.TeamPreset.Values[0].Features[0].Weight != updatedWeight {
		t.Errorf("Feature weight - expected: %d, got: %d", updatedWeight, mp.TeamPreset.Values[0].Features[0].Weight)
	}
}

func TestMasterPrompt_RemoveFeatureByID(t *testing.T) {
	// Setup test data with descriptive IDs
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "feature1_1", Name: "Feature1", Weight: 10},
						{ID: "feature1_2", Name: "Feature2", Weight: 10},
					},
				},
				{
					ID:   "person2",
					Name: "Person2",
					Features: []Feature{
						{ID: "feature2_1", Name: "Feature3", Weight: 10},
						{ID: "feature2_2", Name: "Feature4", Weight: 10},
					},
				},
			},
		},
		StylePreset: StylePreset{
			Values: []Style{
				{ID: "style1", Name: "Style1", Weight: 10},
				{ID: "style2", Name: "Style2", Weight: 10},
				{ID: "style3", Name: "Style3", Weight: 100},
			},
		},
		RulePreset: RulePreset{
			Values: []Rule{
				{ID: "rule1", Name: "Rule1", Weight: 10},
				{ID: "rule2", Name: "Rule2", Weight: 10},
				{ID: "rule3", Name: "Rule3", Weight: 100},
			},
		},
	}

	// SECTION 1: Test removing features from Person1
	mp.RemoveFeatureByID("feature1_1")
	if len(mp.TeamPreset.Values[0].Features) != 1 {
		t.Errorf("Person1 features - expected: 1 feature, got: %d", len(mp.TeamPreset.Values[0].Features))
	}

	mp.RemoveFeatureByID("feature1_2")
	if len(mp.TeamPreset.Values[0].Features) != 0 {
		t.Errorf("Person1 features - expected: 0 features, got: %d", len(mp.TeamPreset.Values[0].Features))
	}

	// SECTION 2: Test removing features from Person2
	mp.RemoveFeatureByID("feature2_1")
	if len(mp.TeamPreset.Values[1].Features) != 1 {
		t.Errorf("Person2 features - expected: 1 feature, got: %d", len(mp.TeamPreset.Values[1].Features))
	}

	mp.RemoveFeatureByID("feature2_2")
	if len(mp.TeamPreset.Values[1].Features) != 0 {
		t.Errorf("Person2 features - expected: 0 features, got: %d", len(mp.TeamPreset.Values[1].Features))
	}

	// SECTION 3: Test removing styles
	mp.RemoveFeatureByID("style1")
	expectedStyleCount := 2
	if len(mp.StylePreset.Values) != expectedStyleCount {
		t.Errorf("Styles - expected: %d styles, got: %d", expectedStyleCount, len(mp.StylePreset.Values))
	}

	mp.RemoveFeatureByID("style2")
	expectedStyleCount = 1
	if len(mp.StylePreset.Values) != expectedStyleCount {
		t.Errorf("Styles - expected: %d style, got: %d", expectedStyleCount, len(mp.StylePreset.Values))
	}

	mp.RemoveFeatureByID("style3")
	expectedStyleCount = 0
	if len(mp.StylePreset.Values) != expectedStyleCount {
		t.Errorf("Styles - expected: %d styles, got: %d", expectedStyleCount, len(mp.StylePreset.Values))
	}

	// SECTION 4: Test removing rules
	mp.RemoveFeatureByID("rule1")
	expectedRuleCount := 2
	if len(mp.RulePreset.Values) != expectedRuleCount {
		t.Errorf("Rules - expected: %d rules, got: %d", expectedRuleCount, len(mp.RulePreset.Values))
	}

	mp.RemoveFeatureByID("rule2")
	expectedRuleCount = 1
	if len(mp.RulePreset.Values) != expectedRuleCount {
		t.Errorf("Rules - expected: %d rule, got: %d", expectedRuleCount, len(mp.RulePreset.Values))
	}

	mp.RemoveFeatureByID("rule3")
	expectedRuleCount = 0
	if len(mp.RulePreset.Values) != expectedRuleCount {
		t.Errorf("Rules - expected: %d rules, got: %d", expectedRuleCount, len(mp.RulePreset.Values))
	}
}

func TestMasterPrompt_RemoveTeamMemberByID(t *testing.T) {
	// Setup test data with descriptive IDs
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "alice",
					Name: "Alice",
					Features: []Feature{
						{ID: "alice_feature", Name: "Alice's Feature", Weight: 10},
					},
				},
				{
					ID:   "bob",
					Name: "Bob",
					Features: []Feature{
						{ID: "bob_feature", Name: "Bob's Feature", Weight: 20},
					},
				},
				{
					ID:   "charlie",
					Name: "Charlie",
					Features: []Feature{
						{ID: "charlie_feature", Name: "Charlie's Feature", Weight: 30},
					},
				},
			},
		},
	}

	// Test case 1: Remove an existing team member
	initialCount := len(mp.TeamPreset.Values)
	mp.RemoveTeamMemberByID("bob")

	// Check if team member was removed
	expectedCount := initialCount - 1
	if len(mp.TeamPreset.Values) != expectedCount {
		t.Errorf("Team count - expected: %d members, got: %d", expectedCount, len(mp.TeamPreset.Values))
	}

	// Verify the correct member was removed
	for _, person := range mp.TeamPreset.Values {
		if person.ID == "bob" {
			t.Errorf("Bob should have been removed but still exists in the team")
		}
	}

	// Test case 2: Remove a non-existent team member
	currentCount := len(mp.TeamPreset.Values)
	mp.RemoveTeamMemberByID("non-existent-id")

	// Check that no team member was removed
	if len(mp.TeamPreset.Values) != currentCount {
		t.Errorf("Team count - expected: %d members (no change), got: %d",
			currentCount, len(mp.TeamPreset.Values))
	}

	// Test case 3: Remove all remaining team members
	mp.RemoveTeamMemberByID("alice")
	mp.RemoveTeamMemberByID("charlie")

	// Check if team is empty
	if len(mp.TeamPreset.Values) != 0 {
		t.Errorf("Team count - expected: 0 members, got: %d", len(mp.TeamPreset.Values))
	}

	// Test case 4: Try to remove from an empty team
	mp.RemoveTeamMemberByID("alice")

	// Should still be empty without errors
	if len(mp.TeamPreset.Values) != 0 {
		t.Errorf("Team count - expected: 0 members, got: %d", len(mp.TeamPreset.Values))
	}
}

func TestRemoveFeatureByID_WithEmptyField(t *testing.T) {
	// Setup test data with more features to test various removal scenarios
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "f1_1", Name: "Feature1-1", Weight: 10},
						{ID: "f1_2", Name: "Feature1-2", Weight: 20},
						{ID: "f1_3", Name: "Feature1-3", Weight: 30},
					},
				},
				{
					ID:   "person2",
					Name: "Person2",
					Features: []Feature{
						{ID: "f2_1", Name: "Feature2-1", Weight: 40},
						{ID: "f2_2", Name: "Feature2-2", Weight: 50},
					},
				},
				{
					ID:   "person3",
					Name: "Person3",
					Features: []Feature{
						{ID: "f3_1", Name: "Feature3-1", Weight: 60},
						{ID: "f3_2", Name: "Feature3-2", Weight: 70},
						{ID: "f3_3", Name: "Feature3-3", Weight: 80},
						{ID: "f3_4", Name: "Feature3-4", Weight: 90},
					},
				},
			},
		},
	}

	// Define test cases for different feature removal patterns
	tests := []struct {
		name            string
		featureIDs      []string
		checkAfterwards func(t *testing.T, mp *MasterPrompt)
	}{
		{
			name:       "Remove 1 of 3 features from Person1",
			featureIDs: []string{"f1_2"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// Person1: should have 2 regular + 1 empty
				if len(mp.TeamPreset.Values[0].Features) != 3 {
					t.Errorf("Person1 - expected 3 features (2 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[0].Features))
				}
				
				// Verify the correct feature was removed
				for _, f := range mp.TeamPreset.Values[0].Features {
					if f.ID == "f1_2" && f.Name != "" {
						t.Errorf("Feature f1_2 should have been removed")
					}
				}
				
				// Verify last feature is empty
				lastFeature := mp.TeamPreset.Values[0].Features[len(mp.TeamPreset.Values[0].Features)-1]
				if lastFeature.Name != "" {
					t.Errorf("Last feature should be empty, got name: %q", lastFeature.Name)
				}
			},
		},
		{
			name:       "Remove 2 of 3 remaining features from Person1",
			featureIDs: []string{"f1_1", "f1_3"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// Person1: should have 0 regular + 1 empty
				if len(mp.TeamPreset.Values[0].Features) != 1 {
					t.Errorf("Person1 - expected 1 feature (0 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[0].Features))
				}
				
				 // Verify that feature is empty
				if mp.TeamPreset.Values[0].Features[0].Name != "" {
					t.Errorf("Feature should be empty, got name: %q", 
						mp.TeamPreset.Values[0].Features[0].Name)
				}
			},
		},
		{
			name:       "Remove all features (2 of 2) from Person2",
			featureIDs: []string{"f2_1", "f2_2"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// Person2: should have 0 regular + 1 empty
				if len(mp.TeamPreset.Values[1].Features) != 1 {
					t.Errorf("Person2 - expected 1 feature (0 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[1].Features))
				}
				
				// Verify that feature is empty
				if mp.TeamPreset.Values[1].Features[0].Name != "" {
					t.Errorf("Feature should be empty, got name: %q", 
						mp.TeamPreset.Values[1].Features[0].Name)
				}
				
				// Person1: should still have 0 regular + 1 empty
				if len(mp.TeamPreset.Values[0].Features) != 1 {
					t.Errorf("Person1 - expected 1 feature (0 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[0].Features))
				}
			},
		},
		{
			name:       "Remove 2 of 4 features from Person3",
			featureIDs: []string{"f3_2", "f3_4"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// Person3: should have 2 regular + 1 empty
				if len(mp.TeamPreset.Values[2].Features) != 3 {
					t.Errorf("Person3 - expected 3 features (2 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[2].Features))
				}
				
				// Verify correct features were removed
				for _, f := range mp.TeamPreset.Values[2].Features {
					if (f.ID == "f3_2" || f.ID == "f3_4") && f.Name != "" {
						t.Errorf("Feature %s should have been removed", f.ID)
					}
				}
				
				// Verify last feature is empty
				lastFeature := mp.TeamPreset.Values[2].Features[len(mp.TeamPreset.Values[2].Features)-1]
				if lastFeature.Name != "" {
					t.Errorf("Last feature should be empty, got name: %q", lastFeature.Name)
				}
			},
		},
		{
			name:       "Remove remaining features from Person3",
			featureIDs: []string{"f3_1", "f3_3"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// Person3: should have 0 regular + 1 empty
				if len(mp.TeamPreset.Values[2].Features) != 1 {
					t.Errorf("Person3 - expected 1 feature (0 regular + 1 empty), got %d", 
						len(mp.TeamPreset.Values[2].Features))
				}
				
				// Verify that feature is empty
				if mp.TeamPreset.Values[2].Features[0].Name != "" {
					t.Errorf("Feature should be empty, got name: %q", 
						mp.TeamPreset.Values[2].Features[0].Name)
				}
			},
		},
		{
			name:       "Remove non-existent features (should not change anything)",
			featureIDs: []string{"nonexistent1", "nonexistent2"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// All persons should still have the same feature counts as before
				if len(mp.TeamPreset.Values[0].Features) != 1 {
					t.Errorf("Person1 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[0].Features))
				}
				if len(mp.TeamPreset.Values[1].Features) != 1 {
					t.Errorf("Person2 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[1].Features))
				}
				if len(mp.TeamPreset.Values[2].Features) != 1 {
					t.Errorf("Person3 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[2].Features))
				}
			},
		},
		{
			name:       "Remove mix of existent and non-existent features",
			featureIDs: []string{"nonexistent1", "f1_1", "nonexistent2"},
			checkAfterwards: func(t *testing.T, mp *MasterPrompt) {
				// All persons should still have the same feature counts as before
				// (f1_1 was already removed in a previous test case)
				if len(mp.TeamPreset.Values[0].Features) != 1 {
					t.Errorf("Person1 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[0].Features))
				}
				if len(mp.TeamPreset.Values[1].Features) != 1 {
					t.Errorf("Person2 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[1].Features))
				}
				if len(mp.TeamPreset.Values[2].Features) != 1 {
					t.Errorf("Person3 - expected 1 feature, got %d", 
						len(mp.TeamPreset.Values[2].Features))
				}
			},
		},
	}

	// Execute test cases sequentially
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Remove each feature in the list
			for _, id := range tt.featureIDs {
				mp.RemoveFeatureByID(id)
			}
			
			// Add empty field after all removals
			mp.AddOneEmptyField()
			
			// Run the check function
			tt.checkAfterwards(t, mp)
		})
	}
}

func TestFeatureIDConsistency(t *testing.T) {
	// Create a master prompt with known IDs for better tracking
	fixedIDs := map[string]string{
		"person1":     "test-person-1",
		"person2":     "test-person-2",
		"feature1_1":  "test-feature-1-1",
		"feature1_2":  "test-feature-1-2",
		"feature2_1":  "test-feature-2-1",
		"emptyField1": "test-empty-1",
		"emptyField2": "test-empty-2",
	}

	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   fixedIDs["person1"],
					Name: "Person1",
					Features: []Feature{
						{ID: fixedIDs["feature1_1"], Name: "Feature1-1", Weight: 10},
						{ID: fixedIDs["feature1_2"], Name: "Feature1-2", Weight: 20},
					},
				},
				{
					ID:   fixedIDs["person2"],
					Name: "Person2",
					Features: []Feature{
						{ID: fixedIDs["feature2_1"], Name: "Feature2-1", Weight: 30},
					},
				},
			},
		},
	}

	// PART 1: Test ID preservation after serialization and deserialization
	t.Run("Serialization preserves IDs", func(t *testing.T) {
		// Serialize
		serialized, err := json.Marshal(mp)
		if err != nil {
			t.Fatalf("Failed to serialize: %v", err)
		}

		// Deserialize
		var deserialized MasterPrompt
		err = json.Unmarshal(serialized, &deserialized)
		if err != nil {
			t.Fatalf("Failed to deserialize: %v", err)
		}

		// Check all IDs are preserved
		if deserialized.TeamPreset.Values[0].ID != fixedIDs["person1"] {
			t.Errorf("Person1 ID mismatch: expected %s, got %s", 
				fixedIDs["person1"], deserialized.TeamPreset.Values[0].ID)
		}
		
		if deserialized.TeamPreset.Values[0].Features[0].ID != fixedIDs["feature1_1"] {
			t.Errorf("Feature1_1 ID mismatch: expected %s, got %s", 
				fixedIDs["feature1_1"], deserialized.TeamPreset.Values[0].Features[0].ID)
		}
	})

	// PART 2: Test AddOneEmptyField behavior with IDs
	t.Run("AddOneEmptyField preserves existing IDs", func(t *testing.T) {
		// Make a copy to avoid affecting other tests
		mpCopy := *mp
		
		// Store original feature counts
		person1FeatureCount := len(mpCopy.TeamPreset.Values[0].Features)
		person2FeatureCount := len(mpCopy.TeamPreset.Values[1].Features)
		
		// Add an empty field
		mpCopy.AddOneEmptyField()
		
		// Check existing IDs are preserved
		if mpCopy.TeamPreset.Values[0].ID != fixedIDs["person1"] {
			t.Errorf("Person1 ID changed after AddOneEmptyField")
		}
		
		if mpCopy.TeamPreset.Values[0].Features[0].ID != fixedIDs["feature1_1"] {
			t.Errorf("Feature1_1 ID changed after AddOneEmptyField")
		}
		
		// Check new empty features were added
		if len(mpCopy.TeamPreset.Values[0].Features) != person1FeatureCount+1 {
			t.Errorf("Expected %d features for Person1, got %d", 
				person1FeatureCount+1, len(mpCopy.TeamPreset.Values[0].Features))
		}
		
		if len(mpCopy.TeamPreset.Values[1].Features) != person2FeatureCount+1 {
			t.Errorf("Expected %d features for Person2, got %d", 
				person2FeatureCount+1, len(mpCopy.TeamPreset.Values[1].Features))
		}
		
		// Check last feature is empty
		lastFeature1 := mpCopy.TeamPreset.Values[0].Features[len(mpCopy.TeamPreset.Values[0].Features)-1]
		if lastFeature1.Name != "" {
			t.Errorf("Expected empty name for new feature, got %q", lastFeature1.Name)
		}
	})

	// PART 3: Test RemoveFeatureByID behavior with regards to IDs
	t.Run("RemoveFeatureByID preserves other feature IDs", func(t *testing.T) {
		// Make a copy to avoid affecting other tests
		mpCopy := *mp
		
		// Remove a feature
		mpCopy.RemoveFeatureByID(fixedIDs["feature1_2"])
		
		// Check feature was removed
		for _, feature := range mpCopy.TeamPreset.Values[0].Features {
			if feature.ID == fixedIDs["feature1_2"] {
				t.Errorf("Feature should have been removed but still exists")
			}
		}
		
		// Check other feature IDs are preserved
		found := false
		for _, feature := range mpCopy.TeamPreset.Values[0].Features {
			if feature.ID == fixedIDs["feature1_1"] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Feature1_1 ID disappeared after removing Feature1_2")
		}
	})

	// PART 4: Test complex sequence of operations
	t.Run("Complex sequence preserves correct IDs", func(t *testing.T) {
		// Make a copy to avoid affecting other tests
		mpCopy := *mp
		
		// 1. Add empty field
		mpCopy.AddOneEmptyField()
		
		// Capture ID of the new empty field
		emptyFeatureID := mpCopy.TeamPreset.Values[0].Features[len(mpCopy.TeamPreset.Values[0].Features)-1].ID
		
		// 2. Update this empty field
		newName := "Updated Feature"
		newWeight := 50
		mpCopy.UpdateValueByID(emptyFeatureID, &newName, &newWeight)
		
		// 3. Add another empty field
		mpCopy.AddOneEmptyField()
		
		// 4. Remove a different feature
		mpCopy.RemoveFeatureByID(fixedIDs["feature1_1"])
		
		// 5. Verify the updated (previously empty) feature still has the same ID and updated values
		found := false
		for _, feature := range mpCopy.TeamPreset.Values[0].Features {
			if feature.ID == emptyFeatureID {
				found = true
				if feature.Name != newName || feature.Weight != newWeight {
					t.Errorf("Updated feature lost its values: got name=%q, weight=%d", 
						feature.Name, feature.Weight)
				}
				break
			}
		}
		if !found {
			t.Errorf("Updated feature ID (%s) disappeared after sequence of operations", emptyFeatureID)
		}
		
		// 6. Verify the removed feature is gone
		for _, feature := range mpCopy.TeamPreset.Values[0].Features {
			if feature.ID == fixedIDs["feature1_1"] {
				t.Errorf("Removed feature still exists")
			}
		}
	})

	// PART 5: Test for ID duplication issue
	t.Run("No duplicate IDs created", func(t *testing.T) {
		// Make a copy to avoid affecting other tests
		mpCopy := *mp
		
		// Add multiple empty fields
		for i := 0; i < 5; i++ {
			mpCopy.AddOneEmptyField()
		}
		
		// Collect all feature IDs
		allIDs := make(map[string]bool)
		for _, person := range mpCopy.TeamPreset.Values {
			for _, feature := range person.Features {
				if allIDs[feature.ID] {
					t.Errorf("Duplicate ID found: %s", feature.ID)
				}
				allIDs[feature.ID] = true
			}
		}
	})

	// PART 6: Test for non-UUID IDs (in case there's a pattern)
	t.Run("Pattern-based UUIDs", func(t *testing.T) {
		// Create a master prompt with suspicious pattern IDs
		suspiciousMp := &MasterPrompt{
			TeamPreset: TeamPreset{
				Values: []Person{
					{
						ID:   "person-1",
						Name: "Person1",
						Features: []Feature{
							{ID: "feature-1-1", Name: "Feature1", Weight: 10},
							{ID: "feature-1-2", Name: "Feature2", Weight: 20},
						},
					},
					{
						ID:   "person-2",
						Name: "Person2",
						Features: []Feature{
							{ID: "feature-2-1", Name: "Feature3", Weight: 30},
						},
					},
				},
			},
		}
		
		// Remove a feature and see if ID handling behaves differently with pattern IDs
		suspiciousMp.RemoveFeatureByID("feature-1-1")
		
		// Check if feature was removed properly
		found := false
		for _, feature := range suspiciousMp.TeamPreset.Values[0].Features {
			if feature.ID == "feature-1-1" {
				found = true
			}
		}
		
		if found {
			t.Errorf("Feature with pattern ID wasn't properly removed")
		}
	})

	// PART 7: Test UUID generation in AddOneEmptyField
	t.Run("UUID generation in AddOneEmptyField", func(t *testing.T) {
			// Create a fresh master prompt
		freshMp := &MasterPrompt{
			TeamPreset: TeamPreset{
				Values: []Person{
					{
						ID:       "test-person",
						Name:     "TestPerson",
						Features: []Feature{},
					},
				},
			},
		}
		
		// Add empty field and verify it generates a valid UUID
		freshMp.AddOneEmptyField()
		
		// Check if the ID has the UUID format
		if len(freshMp.TeamPreset.Values[0].Features) == 0 {
			t.Errorf("No feature was added")
			return
		}
		
		featureID := freshMp.TeamPreset.Values[0].Features[0].ID
		
		// Check if it's a valid UUID
		_, err := uuid.Parse(featureID)
		if err != nil {
			t.Errorf("Expected valid UUID, got %s which produced error: %v", featureID, err)
		}
		
		// Check if it has the expected format (8-4-4-4-12)
		parts := strings.Split(featureID, "-")
		if len(parts) != 5 || 
		   len(parts[0]) != 8 || 
		   len(parts[1]) != 4 || 
		   len(parts[2]) != 4 || 
		   len(parts[3]) != 4 || 
		   len(parts[4]) != 12 {
			t.Errorf("UUID does not have the expected format: %s", featureID)
		}
	})
}

func TestEmptyFieldAndIDConsistency(t *testing.T) {
	// Create a master prompt with one person
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "feature1", Name: "Feature1", Weight: 10},
					},
				},
			},
		},
	}
	
	// CASE 1: Test that AddOneEmptyField adds exactly one empty field
	// when there isn't one already
	
	// First call should add an empty field
	mp.AddOneEmptyField()
	
	// Capture the ID of the empty field
	if len(mp.TeamPreset.Values[0].Features) != 2 { // Original + 1 empty
		t.Errorf("Expected 2 features (1 original + 1 empty), got %d", 
			len(mp.TeamPreset.Values[0].Features))
	}
	
	emptyFeatureID1 := mp.TeamPreset.Values[0].Features[len(mp.TeamPreset.Values[0].Features)-1].ID
	
	// Second call shouldn't add another empty field since there's already one
	mp.AddOneEmptyField()
	
	// Should still have 2 features (1 original + 1 empty)
	if len(mp.TeamPreset.Values[0].Features) != 2 {
		t.Errorf("Expected 2 features (1 original + 1 empty), got %d", 
			len(mp.TeamPreset.Values[0].Features))
	}
	
	// The empty field ID should remain the same
	emptyFeatureID2 := mp.TeamPreset.Values[0].Features[len(mp.TeamPreset.Values[0].Features)-1].ID
	if emptyFeatureID1 != emptyFeatureID2 {
		t.Errorf("Empty feature IDs should be the same when already present, got %s and %s", 
			emptyFeatureID1, emptyFeatureID2)
	}
	
	// CASE 2: Test behavior after removing a feature
	
	// Remove the original feature
	mp.RemoveFeatureByID("feature1")
	
	// Should now have 1 empty feature
	if len(mp.TeamPreset.Values[0].Features) != 1 {
		t.Errorf("Expected 1 feature (the empty one), got %d", 
			len(mp.TeamPreset.Values[0].Features))
	}
	
	// Add an empty field, shouldn't add another since one already exists
	mp.AddOneEmptyField()
	
	// Should still have just 1 empty feature
	if len(mp.TeamPreset.Values[0].Features) != 1 {
		t.Errorf("Expected 1 feature (empty), got %d", 
			len(mp.TeamPreset.Values[0].Features))
	}
	
	// The feature should be empty
	if mp.TeamPreset.Values[0].Features[0].Name != "" {
		t.Errorf("Feature should be empty, got name: %q", 
			mp.TeamPreset.Values[0].Features[0].Name)
	}
}

