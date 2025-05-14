package state

import (
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"makerworld-analytics/domain"
)

func Key() string {
	return "master-prompt-recipe"
}

func SetMasterPrompt(ctx app.Context, masterPrompt *domain.MasterPrompt) {
	addOneEmptyField(masterPrompt)
	ctx.SetState(Key(), masterPrompt).Persist()
}

func DelMasterPrompt(ctx app.Context) {
	ctx.DelState(Key())
}

func GetMasterPrompt(ctx app.Context) *domain.MasterPrompt {
	var masterPrompt *domain.MasterPrompt
	ctx.GetState(Key(), &masterPrompt)

	if masterPrompt == nil {
		masterPrompt = getDefaultMasterPrompt()
		SetMasterPrompt(ctx, masterPrompt)
	}

	addOneEmptyField(masterPrompt) // adds empty field for each feature, like empty rule, empty style, empty team member - so it's easier to manage UI
	return masterPrompt
}

// possibly could be stored in domain
func getDefaultMasterPrompt() *domain.MasterPrompt {
	return &domain.MasterPrompt{
		Template: domain.TestTemplate,
		StylePreset: domain.StylePreset{
			Values: []domain.Style{
				{
					ID:     uuid.New().String(),
					Name:   "",
					Weight: 100,
				},
			},
		},
		RulePreset: domain.RulePreset{
			Values: []domain.Rule{
				{
					ID:     uuid.New().String(),
					Name:   "",
					Weight: 100,
				},
			},
		},
		TeamPreset: domain.TeamPreset{
			Values: []domain.Person{
				{
					ID:        uuid.New().String(),
					Name:      "",
					EmojiIcon: "👨‍💻",
					Role:      "",
					Features: []domain.Feature{
						{
							ID:     uuid.New().String(),
							Name:   "",
							Weight: 100,
						},
					},
				},
			},
		},
	}
}

func addOneEmptyField(prompt *domain.MasterPrompt) {
	// TODO: There could be constructor for Rule, Style, Person - so it would take less space and be easier to read, anyway not a big deal
	// --- Rules ---
	if n := len(prompt.RulePreset.Values); n == 0 ||
		prompt.RulePreset.Values[n-1].Name != "" {

		prompt.RulePreset.Values = append(prompt.RulePreset.Values, domain.Rule{
			ID:     uuid.New().String(),
			Name:   "",
			Weight: 100,
		})
	}

	// --- Styles ---
	if n := len(prompt.StylePreset.Values); n == 0 ||
		prompt.StylePreset.Values[n-1].Name != "" {

		prompt.StylePreset.Values = append(prompt.StylePreset.Values, domain.Style{
			ID:     uuid.New().String(),
			Name:   "",
			Weight: 100,
		})
	}

	if n := len(prompt.TeamPreset.Values); n == 0 ||
		prompt.TeamPreset.Values[n-1].Name != "" {

		prompt.TeamPreset.Values = append(prompt.TeamPreset.Values, domain.Person{
			ID:        uuid.New().String(),
			Name:      "",
			EmojiIcon: "👨‍💻",
			Role:      "",
			Features: []domain.Feature{{
				ID:     uuid.New().String(),
				Name:   "",
				Weight: 100,
			}},
		})
	}

	for i := range prompt.TeamPreset.Values {
		feat := prompt.TeamPreset.Values[i].Features
		if len(feat) == 0 || feat[len(feat)-1].Name != "" {
			prompt.TeamPreset.Values[i].Features = append(feat, domain.Feature{
				ID:     uuid.New().String(),
				Name:   "",
				Weight: 100,
			})
		}
	}
}
