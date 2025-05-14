package views

import (
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log/slog"
	"sort"
	"strconv"
)

type MainView struct {
	app.Compo
	updateAvailable bool
	MasterPrompt    *domain.MasterPrompt
}

func (m *MainView) OnAppUpdate(ctx app.Context) {
	m.updateAvailable = ctx.AppUpdateAvailable() // Reports that an app update is available.
}

func (m *MainView) OnMount(ctx app.Context) {
	ctx.Dispatch(func(ctx app.Context) {
		masterPromptCopy := state.GetMasterPrompt(ctx)
		m.MasterPrompt = &masterPromptCopy
	})
}

// Todo: Divide into more files
func (m *MainView) Render() app.UI {
	if m.MasterPrompt == nil {
		return app.Div().Text("Loading...")
	}

	renderedMasterPrompt, err := m.MasterPrompt.String()
	if err != nil {
		slog.Error("Error rendering master prompt", err)
	}

	return app.Div().Class("p-24").Body(
		app.If(m.updateAvailable, func() app.UI {
			return app.Button().
				Text("Update app!").
				Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded").
				OnClick(m.onUpdateClick)
		}),

		app.H1().Text("Master prompt generator").Class("text-2xl font-bold mb-4"),
		app.P().Text("This is a tool to help you generate a master prompt for your LLM agent.").Class("text-md opacity-80 mb-1"),
		app.P().Text("Data is stored in your browser, so you won't lose anything after refresh").Class("text-md opacity-80 mb-12"),

		&components.CardComponent{
			Body: []app.UI{
				app.H2().Text("Sharing workspace").Class("text-xl font-bold mb-4"),
				app.P().Class("text-md opacity-80 mb-4").Text("You can share your workspace with others by sending them a link. Just click copy and send it to your mate, you can also store it somewhere in notes and manage versions this way."),
				app.Button().Class("btn btn-primary").Text("Copy link").OnClick(m.copyLinkPressed()),
			},
		},

		m.renderMasterPromptTemplate(),
		app.Div().Class("flex flex-row justify-stretch mb-6 w-full").Body(
			m.renderRules(),
			m.renderStyle(),
		),
		m.renderTeam(),

		&components.CardComponent{
			Body: []app.UI{

				app.H2().Text("Output").Class("text-xl font-bold mb-4"),
				app.P().Class("text-md opacity-80 mb-12").Text("Paste it into your chat gpt space, copilot or any other LLM"),
				app.Textarea().Class("textarea textarea-bordered h-80 w-full").Text(renderedMasterPrompt).Placeholder("There should be your prompt"),
			},
		},
	)
}

func (m *MainView) copyLinkPressed() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		recipeBase64, err := m.MasterPrompt.ToBase64()
		if err != nil {
			slog.Error("Error encoding master prompt to base64", err)
			return
		}

		link := ctx.Page().URL()
		link.Path += "/import"
		query := link.Query()           // Get a copy of the query parameters
		query.Set("data", recipeBase64) // Modify the query parameters
		link.RawQuery = query.Encode()

		// set clipboard
		app.Window().Get("navigator").Get("clipboard").Call("writeText", link.String())

	}
}

func (m *MainView) renderTeam() *components.CardComponent {
	var presetsToSelect []components.OptionData
	for name, template := range domain.TeamPresetsMap {
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
					app.H2().Text("👨‍💻 Team").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("Describe your team members, their roles and features"),
				),
				&components.DropdownComponent[domain.TeamPreset]{
					OptionDataList: presetsToSelect,
					Text:           "Select preset",
					OnClick: func(ctx app.Context, value domain.TeamPreset) {
						teamPreset := value
						m.MasterPrompt.TeamPreset = teamPreset
						state.SetMasterPrompt(ctx, m.MasterPrompt)
					},
				},
			),
			app.Range(m.MasterPrompt.TeamPreset.Values).Slice(func(i int) app.UI {
				var member = m.MasterPrompt.TeamPreset.Values[i]

				var handleNameChange = func(ctx app.Context, e app.Event) {
					newName := ctx.JSSrc().Get("value").String()
					m.MasterPrompt.TeamPreset.Values[i].Name = newName
					state.SetMasterPrompt(ctx, m.MasterPrompt)
				}

				var handleRoleChange = func(ctx app.Context, e app.Event) {
					newRole := ctx.JSSrc().Get("value").String()
					m.MasterPrompt.TeamPreset.Values[i].Role = newRole
					state.SetMasterPrompt(ctx, m.MasterPrompt)
				}
				emojisList := []components.OptionData{}

				for _, emoji := range domain.EmojiList {
					emojisList = append(emojisList, components.OptionData{
						Label: emoji,
						Value: emoji,
					})
				}

				handleEmojiChange := func(ctx app.Context, value string) {
					if value == "" {
						return
					}

					m.MasterPrompt.TeamPreset.Values[i].EmojiIcon = value
					state.SetMasterPrompt(ctx, m.MasterPrompt)
				}

				return &components.CardComponent{
					Body: []app.UI{
						app.Div().Class("flex flex-row align-center mb-6").Body(

							&components.DropdownComponent[string]{
								OptionDataList: emojisList,
								Text:           member.EmojiIcon,
								OnClick:        handleEmojiChange,
								Class:          "mr-2",
							},

							app.Input().Class("input input-md w-full mb-1").Type("text").Placeholder("Put team member name here").
								Value(member.Name).OnChange(handleNameChange).OnKeyUp(handleNameChange),
							app.If(member.Name != "", func() app.UI {
								return app.Button().Class("btn btn-error ml-2").Text("Remove person").OnClick(
									func(ctx app.Context, e app.Event) {
										m.MasterPrompt.RemoveTeamMemberByID(member.ID)
										state.SetMasterPrompt(ctx, m.MasterPrompt)
									})
							}),
						),

						app.IfSlice(member.Name != "", func() []app.UI {
							return []app.UI{
								app.Input().Class("input input-md w-full mb-12").Type("text").Placeholder("Role = developer, customer, marketing specialist and etc..").
									Value(member.Role).OnChange(handleRoleChange).OnKeyUp(handleRoleChange),

								app.Div().Class("flex flex-row justify-between mb-6").Body(
									app.H3().Class("text-xl").Text("Features"),
									app.H3().Class("text-xl").Text("Weight"),
								),
								app.Range(member.Features).Slice(func(j int) app.UI {
									return m.renderWeightControlledName(member.Features[j].ID, member.Features[j].Name, member.Features[j].Weight)
								}),
							}
						}),
					},
					Class: "mb-8 shadow-2xl",
				}
			}),
		},
	}
}

func (m *MainView) renderStyle() *components.CardComponent {
	var presetsToSelect []components.OptionData
	for name, template := range domain.StylePresetsMap {
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
					app.H2().Text("🤌🏻Style").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("More formal? More friendly? Structured in certain way? Describe it here"),
				),
				&components.DropdownComponent[domain.StylePreset]{
					OptionDataList: presetsToSelect,
					Text:           "Select preset",
					OnClick: func(ctx app.Context, value domain.StylePreset) {
						stylePreset := value
						m.MasterPrompt.StylePreset = stylePreset
						state.SetMasterPrompt(ctx, m.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Style hint"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(m.MasterPrompt.StylePreset.Values).Slice(func(i int) app.UI {
					return m.renderWeightControlledName(m.MasterPrompt.StylePreset.Values[i].ID, m.MasterPrompt.StylePreset.Values[i].Name, m.MasterPrompt.StylePreset.Values[i].Weight)
				}),
			),
		},
		Class: "w-full",
	}
}

func (m *MainView) renderRules() *components.CardComponent {
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
					Text:           "Select preset",
					OnClick: func(ctx app.Context, value domain.RulePreset) {
						rulesPreset := value
						m.MasterPrompt.RulePreset = rulesPreset
						state.SetMasterPrompt(ctx, m.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Rule"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(m.MasterPrompt.RulePreset.Values).Slice(func(i int) app.UI {
					return m.renderWeightControlledName(m.MasterPrompt.RulePreset.Values[i].ID, m.MasterPrompt.RulePreset.Values[i].Name, m.MasterPrompt.RulePreset.Values[i].Weight)
				}),
			),
		},
		Class: "mr-4 w-full",
	}
}

func (m *MainView) renderMasterPromptTemplate() *components.CardComponent {
	var presetsToSelect []components.OptionData
	for name, template := range domain.AllMasterTemplatesMap {
		presetsToSelect = append(presetsToSelect, components.OptionData{
			Label: name,
			Value: string(template),
		})
	}

	sort.Slice(presetsToSelect, func(i, j int) bool {
		return len(presetsToSelect[i].Label) < len(presetsToSelect[j].Label)
	})

	onMasterPromptTemplateChanged := func(ctx app.Context, e app.Event) {
		newTemplate := ctx.JSSrc().Get("value").String()
		m.MasterPrompt.Template = domain.MasterPromptTemplate(newTemplate)
		state.SetMasterPrompt(ctx, m.MasterPrompt)
	}

	return &components.CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("Master prompt template").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("This is the template that will be used to generate the master prompt."),
				),
				&components.DropdownComponent[string]{
					OptionDataList: presetsToSelect,
					Text:           "Select preset",
					OnClick: func(ctx app.Context, value string) {
						if value == "" {
							return
						}

						m.MasterPrompt.Template = domain.MasterPromptTemplate(value)
						state.SetMasterPrompt(ctx, m.MasterPrompt)
					},
				},
			),

			app.Textarea().Class("textarea textarea-bordered h-80 w-full").
				Placeholder("Enter your master prompt template here").
				Text(m.MasterPrompt.Template).
				OnChange(onMasterPromptTemplateChanged).OnKeyUp(onMasterPromptTemplateChanged),
		},
	}
}

func (m *MainView) renderWeightControlledName(id string, name string, weight int) app.UI {
	return app.Div().Class("flex flex-row w-full justify-between items-center mb-2").Body(
		app.Div().Class("flex flex-row w-full").Body(
			app.If(name != "", func() app.UI {
				return app.Button().Class("btn btn-error mr-2").Text("X").OnClick(
					func(ctx app.Context, e app.Event) {
						m.MasterPrompt.RemoveFeatureByID(id)
						state.SetMasterPrompt(ctx, m.MasterPrompt)
					},
				)
			}),
			app.Input().Type("text").Placeholder("Put rule here").Class("input input-md w-full").Value(name).OnChange(
				func(ctx app.Context, e app.Event) {
					newName := ctx.JSSrc().Get("value").String()
					m.MasterPrompt.UpdateValueByID(id, &newName, nil)
					state.SetMasterPrompt(ctx, m.MasterPrompt)
				}),
		),

		app.If(name != "", func() app.UI {
			handleSliderInput := func(ctx app.Context, e app.Event) {
				newWeight, err := strconv.Atoi(ctx.JSSrc().Get("value").String())
				if err != nil {
					slog.Error("Error converting weight to int", err)
					return
				}
				m.MasterPrompt.UpdateValueByID(id, nil, &newWeight)
				state.SetMasterPrompt(ctx, m.MasterPrompt)
			}
			return app.Input().Class("w-96 ml-4").Type("range").Min(0).Max(100).Class("range").Value(weight).
				OnChange(handleSliderInput).
				OnPaste(handleSliderInput).
				OnDrag(handleSliderInput).OnMouseMove(handleSliderInput)
		}),
	)
}

func (m *MainView) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
