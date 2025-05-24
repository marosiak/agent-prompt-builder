package views

import (
	"github.com/marosiak/agent-prompt-builder/ui/components"
	"sort"

	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// RulesEditorComponent handles the rules presets and editing
type RulesEditorComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
}

func (r *RulesEditorComponent) Render() app.UI {
	var presetsToSelect []components.OptionData
	for name, template := range domain.RulesPresetsMap {
		presetsToSelect = append(presetsToSelect, components.OptionData{
			Label: name,
			Value: template,
		})
	}

	sort.Slice(presetsToSelect, func(i, j int) bool {
		return len(presetsToSelect[i].Label) < len(presetsToSelect[j].Label)
	})

	return &components.CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("📜 Rules").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("It's great to give positive rules instead negative, like 'do this' instead 'don't do this' "),
				),
				&components.DropdownComponent[domain.RulePreset]{
					OptionDataList: presetsToSelect,
					Text:           "Preset",
					Icon:           &components.SVGIcon{IconData: components.SlidersIcon, Color: "black", IconSize: components.IconSizeMedium, OpacityPercent: 55},
					Position:       components.DropdownPositionLeft,
					OnClick: func(ctx app.Context, value domain.RulePreset) {
						r.MasterPrompt.RulePreset = value
						state.SetMasterPromptWithEmptyField(ctx, &r.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Rule"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(func() []domain.Rule {
					rules := make([]domain.Rule, len(r.MasterPrompt.RulePreset.Values))
					copy(rules, r.MasterPrompt.RulePreset.Values)
					return rules
				}()).Slice(func(i int) app.UI {
					rule := r.MasterPrompt.RulePreset.Values[i]
					return &components.WeightedItemComponent{
						ID:        rule.ID,
						Name:      rule.Name,
						Weight:    rule.Weight,
						ShowTrash: rule.Name != "",
						OnDelete: func(ctx app.Context, id string) {
							r.MasterPrompt.RemoveFeatureByID(id)
							state.SetMasterPromptWithEmptyField(ctx, &r.MasterPrompt)
						},
						OnUpdate: func(ctx app.Context, id string, name *string, weight *int) {
							r.MasterPrompt.UpdateValueByID(id, name, weight)
							state.SetMasterPromptWithEmptyField(ctx, &r.MasterPrompt)
						},
					}
				}),
			),
		},
		Class: "mr-4 w-full",
	}
}
