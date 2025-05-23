package components

import (
	"sort"

	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TemplateEditorComponent handles the master prompt template editing
type TemplateEditorComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
}

func (t *TemplateEditorComponent) Render() app.UI {
	var presetsToSelect []OptionData
	for name, template := range domain.AllMasterTemplatesMap {
		presetsToSelect = append(presetsToSelect, OptionData{
			Label: name,
			Value: string(template),
		})
	}

	sort.Slice(presetsToSelect, func(i, j int) bool {
		return len(presetsToSelect[i].Label) < len(presetsToSelect[j].Label)
	})

	onMasterPromptTemplateChanged := func(ctx app.Context, e app.Event) {
		newTemplate := ctx.JSSrc().Get("value").String()
		t.MasterPrompt.Template = domain.MasterPromptTemplate(newTemplate)
		state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
	}

	return &CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("Prompt template").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("This is the template that will be used to generate the master prompt."),
				),
				&DropdownComponent[string]{
					OptionDataList: presetsToSelect,
					Text:           "Preset",
					Icon:           &SVGIcon{IconData: SlidersIcon, Color: "black", IconSize: IconSizeMedium, OpacityPercent: 55},
					Position:       DropdownPositionLeft,
					OnClick: func(ctx app.Context, value string) {
						if value == "" {
							return
						}

						t.MasterPrompt.Template = domain.MasterPromptTemplate(value)
						state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
					},
				},
			),

			app.Textarea().Class("textarea textarea-bordered h-80 w-full").
				Placeholder("Enter your master prompt template here").
				Text(t.MasterPrompt.Template).
				OnChange(onMasterPromptTemplateChanged).OnKeyUp(onMasterPromptTemplateChanged),
		},
	}
}
