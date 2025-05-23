package components

import (
	"sort"

	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// StyleEditorComponent handles the style presets and editing
type StyleEditorComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
}

func (s *StyleEditorComponent) Render() app.UI {
	var presetsToSelect []OptionData
	for name, template := range domain.StylePresetsMap {
		presetsToSelect = append(presetsToSelect, OptionData{
			Label: name,
			Value: template,
		})
	}

	sort.Slice(presetsToSelect, func(i, j int) bool {
		return len(presetsToSelect[i].Label) < len(presetsToSelect[j].Label)
	})

	return &CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("🤌🏻Style").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("More formal? More friendly? Structured in certain way? Describe it here"),
				),
				&DropdownComponent[domain.StylePreset]{
					OptionDataList: presetsToSelect,
					Text:           "Preset",
					Icon:           &SVGIcon{IconData: SlidersIcon, Color: "black", IconSize: IconSizeMedium, OpacityPercent: 55},
					Position:       DropdownPositionLeft,
					OnClick: func(ctx app.Context, value domain.StylePreset) {
						s.MasterPrompt.StylePreset = value
						state.SetMasterPromptWithEmptyField(ctx, &s.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Style hint"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(func() []domain.Style {
					styles := make([]domain.Style, len(s.MasterPrompt.StylePreset.Values))
					copy(styles, s.MasterPrompt.StylePreset.Values)
					return styles
				}()).Slice(func(i int) app.UI {
					style := s.MasterPrompt.StylePreset.Values[i]
					return &WeightedItemComponent{
						ID:        style.ID,
						Name:      style.Name,
						Weight:    style.Weight,
						ShowTrash: style.Name != "",
						OnDelete: func(ctx app.Context, id string) {
							s.MasterPrompt.RemoveFeatureByID(id)
							state.SetMasterPromptWithEmptyField(ctx, &s.MasterPrompt)
						},
						OnUpdate: func(ctx app.Context, id string, name *string, weight *int) {
							s.MasterPrompt.UpdateValueByID(id, name, weight)
							state.SetMasterPromptWithEmptyField(ctx, &s.MasterPrompt)
						},
					}
				}),
			),
		},
		Class: "w-full",
	}
}
