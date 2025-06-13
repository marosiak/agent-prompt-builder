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

	// PART 1: Test id preservation after serialization and deserialization
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
			t.Errorf("Person1 id mismatch: expected %s, got %s",
				fixedIDs["person1"], deserialized.TeamPreset.Values[0].ID)
		}

		if deserialized.TeamPreset.Values[0].Features[0].ID != fixedIDs["feature1_1"] {
			t.Errorf("Feature1_1 id mismatch: expected %s, got %s",
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
			t.Errorf("Person1 id changed after AddOneEmptyField")
		}

		if mpCopy.TeamPreset.Values[0].Features[0].ID != fixedIDs["feature1_1"] {
			t.Errorf("Feature1_1 id changed after AddOneEmptyField")
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
			t.Errorf("Feature1_1 id disappeared after removing Feature1_2")
		}
	})

	// PART 4: Test complex sequence of operations
	t.Run("Complex sequence preserves correct IDs", func(t *testing.T) {
		// Make a copy to avoid affecting other tests
		mpCopy := *mp

		// 1. Add empty field
		mpCopy.AddOneEmptyField()

		// Capture id of the new empty field
		emptyFeatureID := mpCopy.TeamPreset.Values[0].Features[len(mpCopy.TeamPreset.Values[0].Features)-1].ID

		// 2. Update this empty field
		newName := "Updated Feature"
		newWeight := 50
		mpCopy.UpdateValueByID(emptyFeatureID, &newName, &newWeight)

		// 3. Add another empty field
		mpCopy.AddOneEmptyField()

		// 4. Remove a different feature
		mpCopy.RemoveFeatureByID(fixedIDs["feature1_1"])

		// 5. Verify the updated (previously empty) feature still has the same id and updated values
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
			t.Errorf("Updated feature id (%s) disappeared after sequence of operations", emptyFeatureID)
		}

		// 6. Verify the removed feature is gone
		for _, feature := range mpCopy.TeamPreset.Values[0].Features {
			if feature.ID == fixedIDs["feature1_1"] {
				t.Errorf("Removed feature still exists")
			}
		}
	})

	// PART 5: Test for id duplication issue
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
					t.Errorf("Duplicate id found: %s", feature.ID)
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

		// Remove a feature and see if id handling behaves differently with pattern IDs
		suspiciousMp.RemoveFeatureByID("feature-1-1")

		// Check if feature was removed properly
		found := false
		for _, feature := range suspiciousMp.TeamPreset.Values[0].Features {
			if feature.ID == "feature-1-1" {
				found = true
			}
		}

		if found {
			t.Errorf("Feature with pattern id wasn't properly removed")
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

		// Check if the id has the UUID format
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

	// Capture the id of the empty field
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

	// The empty field id should remain the same
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

func TestFeatureIdConsistencyAfterRemoval(t *testing.T) {
	// Create a master prompt with a person that has multiple features
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "A", Name: "Feature A", Weight: 10},
						{ID: "B", Name: "Feature B", Weight: 20},
						{ID: "C", Name: "Feature C", Weight: 30},
						{ID: "D", Name: "Feature D", Weight: 40},
						{ID: "E", Name: "Feature E", Weight: 50},
						{ID: "F", Name: "Feature F", Weight: 60},
					},
				},
			},
		},
	}

	// Store a map of [id => Name] before removal to verify later
	originalFeatures := make(map[string]string)
	for _, feature := range mp.TeamPreset.Values[0].Features {
		originalFeatures[feature.ID] = feature.Name
	}

	// STEP 1: Remove feature C (in the middle)
	featureToRemove := "C"
	mp.RemoveFeatureByID(featureToRemove)

	// STEP 2: Call AddOneEmptyField which might be causing the issue
	// This simulates what happens in the UI flow
	mp.AddOneEmptyField()

	// Verify features by id - we need to ensure:
	// 1. Feature C is gone
	// 2. Features A, B, D, E, F still have the same IDs and correct data
	remainingFeatures := make(map[string]bool)
	for _, feature := range mp.TeamPreset.Values[0].Features {
		// Skip empty feature that was added by AddOneEmptyField
		if feature.Name == "" {
			continue
		}

		// Check that feature with this id exists in original map
		originalName, exists := originalFeatures[feature.ID]
		if !exists {
			t.Errorf("Feature with id %s was not in the original set but appears now", feature.ID)
			continue
		}

		// Check that feature name matches the original
		if originalName != feature.Name {
			t.Errorf("Feature id %s has name %s but originally had name %s",
				feature.ID, feature.Name, originalName)
		}

		remainingFeatures[feature.ID] = true
	}

	// Ensure feature C is gone
	if remainingFeatures[featureToRemove] {
		t.Errorf("Feature with id %s should have been removed", featureToRemove)
	}

	// Ensure all other original features still exist (except C)
	for id, name := range originalFeatures {
		if id != featureToRemove && !remainingFeatures[id] {
			t.Errorf("Feature with id %s and name %s is missing after removal", id, name)
		}
	}
}

func TestConsistencyWithMultipleRemovals(t *testing.T) {
	// Create a master prompt with multiple people/features
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "feature1_A", Name: "Person1 Feature A", Weight: 10},
						{ID: "feature1_B", Name: "Person1 Feature B", Weight: 20},
						{ID: "feature1_C", Name: "Person1 Feature C", Weight: 30},
					},
				},
				{
					ID:   "person2",
					Name: "Person2",
					Features: []Feature{
						{ID: "feature2_A", Name: "Person2 Feature A", Weight: 40},
						{ID: "feature2_B", Name: "Person2 Feature B", Weight: 50},
						{ID: "feature2_C", Name: "Person2 Feature C", Weight: 60},
					},
				},
			},
		},
	}

	// Store all feature information before removal
	type FeatureInfo struct {
		PersonIndex int
		Name        string
		Weight      int
	}
	originalFeatureInfo := make(map[string]FeatureInfo)

	for personIndex, person := range mp.TeamPreset.Values {
		for _, feature := range person.Features {
			originalFeatureInfo[feature.ID] = FeatureInfo{
				PersonIndex: personIndex,
				Name:        feature.Name,
				Weight:      feature.Weight,
			}
		}
	}

	// Remove features in a specific order
	removalOrder := []string{"feature1_B", "feature2_A", "feature1_C"}
	for _, featureID := range removalOrder {
		mp.RemoveFeatureByID(featureID)
		// Simulate UI behavior by calling AddOneEmptyField after each removal
		mp.AddOneEmptyField()
	}

	// Check all remaining features have consistent data
	for personIndex, person := range mp.TeamPreset.Values {
		for _, feature := range person.Features {
			// Skip empty features
			if feature.Name == "" {
				continue
			}

			// Verify feature exists in original map
			info, exists := originalFeatureInfo[feature.ID]
			if !exists {
				t.Errorf("Feature with id %s was not in original set but appears now", feature.ID)
				continue
			}

			// Verify feature is associated with the correct person
			if info.PersonIndex != personIndex {
				t.Errorf("Feature with id %s moved from person %d to person %d",
					feature.ID, info.PersonIndex, personIndex)
			}

			// Verify feature data matches original
			if info.Name != feature.Name {
				t.Errorf("Feature id %s has name %s but originally had name %s",
					feature.ID, feature.Name, info.Name)
			}

			if info.Weight != feature.Weight {
				t.Errorf("Feature id %s has weight %d but originally had weight %d",
					feature.ID, feature.Weight, info.Weight)
			}
		}
	}

	// Verify removed features are actually gone
	removedIds := make(map[string]bool)
	for _, id := range removalOrder {
		removedIds[id] = true
	}

	for personIndex, person := range mp.TeamPreset.Values {
		for _, feature := range person.Features {
			if removedIds[feature.ID] && feature.Name != "" {
				t.Errorf("Feature with id %s should have been removed but still exists in person %d",
					feature.ID, personIndex)
			}
		}
	}
}

func TestRenderingWithRemoval(t *testing.T) {
	// This test simulates the UI rendering behavior
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "A", Name: "Feature A", Weight: 10},
						{ID: "B", Name: "Feature B", Weight: 20},
						{ID: "C", Name: "Feature C", Weight: 30},
						{ID: "D", Name: "Feature D", Weight: 40},
					},
				},
			},
		},
	}

	// Record original data
	type FeatureData struct {
		ID     string
		Name   string
		Weight int
	}
	originalFeatures := make([]FeatureData, 0)
	for _, feature := range mp.TeamPreset.Values[0].Features {
		originalFeatures = append(originalFeatures, FeatureData{
			ID:     feature.ID,
			Name:   feature.Name,
			Weight: feature.Weight,
		})
	}

	// Remove the feature in the middle (C)
	mp.RemoveFeatureByID("C")

	// Simulate UI rendering by iterating through features like in the UI code
	renderedFeatures := make([]FeatureData, 0)
	for j := range mp.TeamPreset.Values[0].Features {
		// This mimics the go-app rendering iteration pattern
		feature := mp.TeamPreset.Values[0].Features[j]
		renderedFeatures = append(renderedFeatures, FeatureData{
			ID:     feature.ID,
			Name:   feature.Name,
			Weight: feature.Weight,
		})
	}

	// Verify each remaining feature has correct data
	expectedRemainingIDs := []string{"A", "B", "D"}
	for _, id := range expectedRemainingIDs {
		found := false
		for _, feature := range renderedFeatures {
			if feature.ID == id {
				found = true

				// Find original data for this id
				var originalData FeatureData
				for _, orig := range originalFeatures {
					if orig.ID == id {
						originalData = orig
						break
					}
				}

				// Verify data matches
				if feature.Name != originalData.Name {
					t.Errorf("Feature %s has name %s but originally had %s after removal of C",
						id, feature.Name, originalData.Name)
				}

				if feature.Weight != originalData.Weight {
					t.Errorf("Feature %s has weight %d but originally had %d after removal of C",
						id, feature.Weight, originalData.Weight)
				}

				break
			}
		}

		if !found {
			t.Errorf("Feature with id %s not found after rendering", id)
		}
	}

	// Verify C is gone
	for _, feature := range renderedFeatures {
		if feature.ID == "C" && feature.Name != "" {
			t.Errorf("Feature C should have been removed but still exists with name: %s", feature.Name)
		}
	}
}

func TestRemovalAndIdConsistency(t *testing.T) {
	// Create a master prompt with carefully tracked IDs
	mp := &MasterPrompt{
		TeamPreset: TeamPreset{
			Values: []Person{
				{
					ID:   "person1",
					Name: "Person1",
					Features: []Feature{
						{ID: "A", Name: "Feature A", Weight: 10},
						{ID: "B", Name: "Feature B", Weight: 20},
						{ID: "C", Name: "Feature C", Weight: 30},
						{ID: "D", Name: "Feature D", Weight: 40},
						{ID: "E", Name: "Feature E", Weight: 50},
						{ID: "F", Name: "Feature F", Weight: 60},
					},
				},
			},
		},
	}

	// Step 1: Record original feature-to-name mappings
	originalFeatures := make(map[string]string)
	for _, feature := range mp.TeamPreset.Values[0].Features {
		originalFeatures[feature.ID] = feature.Name
	}

	// Step 2: Remove feature C
	featureToRemove := "C"
	mp.RemoveFeatureByID(featureToRemove)

	// Step 3: Simulate UI process - NOTE: not calling AddOneEmptyField to avoid id issues

	// Step 4: Check all remaining features maintained their IDs
	idToPositionAfterRemoval := make(map[string]int)
	for i, feature := range mp.TeamPreset.Values[0].Features {
		idToPositionAfterRemoval[feature.ID] = i

		if feature.ID == featureToRemove {
			t.Errorf("Feature with id %s should have been removed", featureToRemove)
		}

		// For remaining features, verify their name didn't change
		if originalName, exists := originalFeatures[feature.ID]; exists {
			if feature.Name != originalName {
				t.Errorf("Feature id %s should have name %s but has %s",
					feature.ID, originalName, feature.Name)
			}
		}
	}

	// Step 5: Verify the remaining IDs are A, B, D, E, F (not C)
	expectedRemainingIDs := []string{"A", "B", "D", "E", "F"}
	for _, id := range expectedRemainingIDs {
		if _, exists := idToPositionAfterRemoval[id]; !exists {
			t.Errorf("Feature id %s should still exist after removal", id)
		}
	}

	// Step 6: Only NOW add empty field and verify it doesn't affect other IDs
	mp.AddOneEmptyField()

	// Check positions again
	idToPositionAfterEmptyField := make(map[string]int)
	for i, feature := range mp.TeamPreset.Values[0].Features {
		// Skip the empty feature
		if feature.Name == "" {
			continue
		}

		idToPositionAfterEmptyField[feature.ID] = i

		// Verify name still matches original
		if originalName, exists := originalFeatures[feature.ID]; exists {
			if feature.Name != originalName {
				t.Errorf("After AddOneEmptyField, feature id %s should have name %s but has %s",
					feature.ID, originalName, feature.Name)
			}
		}
	}

	// Make sure all previous IDs are still there
	for id := range idToPositionAfterRemoval {
		if _, exists := idToPositionAfterEmptyField[id]; !exists {
			t.Errorf("Feature id %s disappeared after AddOneEmptyField", id)
		}
	}
}

func TestMasterPrompt_AddOneEmptyField_TableDriven(t *testing.T) {
	tests := []struct {
		name                 string
		initialTeamMembers   []Person
		initialStyleValues   []Style
		initialRuleValues    []Rule
		expectedTeamCount    int
		expectedRuleCount    int
		expectedStyleCount   int
		validateEmptyEntries bool // whether to check for empty entries at the end
	}{
		{
			name: "All empty team members",
			initialTeamMembers: []Person{
				{ID: "person1", Name: ""},
				{ID: "person2", Name: ""},
				{ID: "person3", Name: ""},
			},
			initialStyleValues:   []Style{},
			initialRuleValues:    []Rule{},
			expectedTeamCount:    1, // should consolidate to just one empty team member
			expectedRuleCount:    1, // should add one empty rule
			expectedStyleCount:   1, // should add one empty style
			validateEmptyEntries: true,
		},
		{
			name: "Mix of empty and non-empty team members",
			initialTeamMembers: []Person{
				{ID: "person1", Name: "Alice"},
				{ID: "person2", Name: ""},
				{ID: "person3", Name: "Bob"},
			},
			initialStyleValues:   []Style{},
			initialRuleValues:    []Rule{},
			expectedTeamCount:    3, // two non-empty plus one empty
			expectedRuleCount:    1,
			expectedStyleCount:   1,
			validateEmptyEntries: true,
		},
		{
			name: "All non-empty team members",
			initialTeamMembers: []Person{
				{ID: "person1", Name: "Alice"},
				{ID: "person2", Name: "Bob"},
				{ID: "person3", Name: "Charlie"},
			},
			initialStyleValues:   []Style{},
			initialRuleValues:    []Rule{},
			expectedTeamCount:    4, // three non-empty plus one empty
			expectedRuleCount:    1,
			expectedStyleCount:   1,
			validateEmptyEntries: true,
		},
		{
			name:                 "No entries in all presets",
			initialTeamMembers:   []Person{},
			initialStyleValues:   []Style{},
			initialRuleValues:    []Rule{},
			expectedTeamCount:    1, // should add one empty team member
			expectedRuleCount:    1, // should add one empty rule
			expectedStyleCount:   1, // should add one empty style
			validateEmptyEntries: true,
		},
		{
			name: "Already has empty rules and styles",
			initialTeamMembers: []Person{
				{ID: "person1", Name: "Alice"},
			},
			initialStyleValues: []Style{
				{ID: "style1", Name: "Bold"},
				{ID: "style2", Name: ""},
			},
			initialRuleValues: []Rule{
				{ID: "rule1", Name: "Be consistent"},
				{ID: "rule2", Name: ""},
			},
			expectedTeamCount:    2, // one non-empty plus one empty
			expectedRuleCount:    2, // one non-empty plus existing empty
			expectedStyleCount:   2, // one non-empty plus existing empty
			validateEmptyEntries: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mp := &MasterPrompt{
				TeamPreset: TeamPreset{
					Values: tt.initialTeamMembers,
				},
				StylePreset: StylePreset{
					Values: tt.initialStyleValues,
				},
				RulePreset: RulePreset{
					Values: tt.initialRuleValues,
				},
			}

			mp.AddOneEmptyField()

			// Check team preset count
			if len(mp.TeamPreset.Values) != tt.expectedTeamCount {
				t.Errorf("TeamPreset - expected %d members, got %d",
					tt.expectedTeamCount, len(mp.TeamPreset.Values))
			}

			// Check rule preset count
			if len(mp.RulePreset.Values) != tt.expectedRuleCount {
				t.Errorf("RulePreset - expected %d rules, got %d",
					tt.expectedRuleCount, len(mp.RulePreset.Values))
			}

			// Check style preset count
			if len(mp.StylePreset.Values) != tt.expectedStyleCount {
				t.Errorf("StylePreset - expected %d styles, got %d",
					tt.expectedStyleCount, len(mp.StylePreset.Values))
			}

			// Validate that last entries are empty if requested
			if tt.validateEmptyEntries {
				// Check team has an empty entry as the last item
				if len(mp.TeamPreset.Values) > 0 {
					lastTeamMember := mp.TeamPreset.Values[len(mp.TeamPreset.Values)-1]
					if lastTeamMember.Name != "" {
						t.Errorf("Last team member should be empty, got name: %q", lastTeamMember.Name)
					}
				}

				// Check rules has an empty entry as the last item
				if len(mp.RulePreset.Values) > 0 {
					lastRule := mp.RulePreset.Values[len(mp.RulePreset.Values)-1]
					if lastRule.Name != "" {
						t.Errorf("Last rule should be empty, got name: %q", lastRule.Name)
					}
				}

				// Check styles has an empty entry as the last item
				if len(mp.StylePreset.Values) > 0 {
					lastStyle := mp.StylePreset.Values[len(mp.StylePreset.Values)-1]
					if lastStyle.Name != "" {
						t.Errorf("Last style should be empty, got name: %q", lastStyle.Name)
					}
				}

				// Check each team member has an empty feature at the end
				for i, person := range mp.TeamPreset.Values {
					if len(person.Features) == 0 {
						t.Errorf("Team member at index %d should have at least one feature", i)
					} else {
						lastFeature := person.Features[len(person.Features)-1]
						if lastFeature.Name != "" {
							t.Errorf("Last feature of team member at index %d should be empty, got name: %q",
								i, lastFeature.Name)
						}
					}
				}
			}
		})
	}
}
