package views

import (
	. "github.com/marosiak/agent-prompt-builder/ui/components"
	"sort"

	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// TemplateEditorComponent handles the master prompt template editing
type TemplateEditorComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
	currentText  string // mutable during lifecycle, loaded OnMount
}

func (t *TemplateEditorComponent) OnMount(ctx app.Context) {
	t.currentText = string(t.MasterPrompt.Template)
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

	isValid, errDetails := t.MasterPrompt.Template.IsValid()

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
						t.currentText = value
						state.SetMasterPromptWithEmptyField(ctx, &t.MasterPrompt)
					},
				},
			),
			app.If(!isValid, func() app.UI {
				return t.renderErrors(*errDetails)
			}),
			app.Textarea().Class("textarea textarea-bordered h-80 w-full").
				Placeholder("Enter your master prompt template here").
				Text(t.currentText).
				OnChange(onMasterPromptTemplateChanged).OnKeyUp(onMasterPromptTemplateChanged),
		},
	}
}

func (t *TemplateEditorComponent) renderErrors(errDetails domain.PromptTemplateValidation) app.HTMLDiv {
	return app.Div().Class("flex flex-col gap-2 mb-4").Body(
		app.If(errDetails.RulesPlaceholderMissing, func() app.UI {
			return t.renderStatus(true, "Rules placeholder is missing, insert $$rules$$ into template")
		}),

		app.If(errDetails.StylePlaceholderMissing, func() app.UI {
			return t.renderStatus(true, "Style placeholder is missing, insert $$style$$ into template")
		}),

		app.If(errDetails.TeamPlaceholderMissing, func() app.UI {
			return t.renderStatus(true, "Team placeholder is missing, insert $$team$$ into template")
		}),
	)
}

func (t *TemplateEditorComponent) renderStatus(isError bool, text string) app.UI {
	color := StatusColorSuccess
	class := ""
	if isError {
		color = StatusColorError
		class = "text-error-content"
	}

	return &StatusComponent{
		Color: color,
		Size:  StatusSizeLarge,
		Label: app.P().Class("text-md " + class).Text(text),
	}
}
