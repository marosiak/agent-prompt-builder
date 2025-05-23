package views

import (
	"github.com/marosiak/agent-prompt-builder/actions"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	. "github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log/slog"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type MainView struct {
	app.Compo
	updateAvailable    bool
	MasterPrompt       *domain.MasterPrompt
	CurrentPageIndex   int
	ctx                app.Context
	KeystrokePressedAt *time.Time

	TeamPresetsModal *TeamPresetsModal
}

func (m *MainView) OnAppUpdate(ctx app.Context) {
	m.updateAvailable = ctx.AppUpdateAvailable()
}

func (m *MainView) OnMount(ctx app.Context) {
	m.ctx = ctx
	ctx.Dispatch(func(ctx app.Context) {
		masterPromptCopy := state.GetMasterPrompt(ctx)
		m.MasterPrompt = &masterPromptCopy
		m.CurrentPageIndex = state.GetCurrentPageIndex(ctx)
	})

	var tmpMasterPrompt domain.MasterPrompt
	ctx.ObserveState(state.MasterPromptKey(), &tmpMasterPrompt).OnChange(func() {
		m.MasterPrompt = &tmpMasterPrompt
	})

	var tmpPageIndex int
	ctx.ObserveState(state.PageStateKey(), &tmpPageIndex).OnChange(func() {
		m.CurrentPageIndex = tmpPageIndex
	})

	ctx.Handle(actions.ActionOpenTeamPresetsModal, m.HandleOpenTeamPresetsModal)
	ctx.Handle(actions.AddPerson, m.HandleAddPerson)
	ctx.Handle(actions.RemovePerson, m.HandleRemovePerson)
}

func (m *MainView) HandleOpenTeamPresetsModal(ctx app.Context, a app.Action) {
	if m.TeamPresetsModal == nil {
		m.TeamPresetsModal = &TeamPresetsModal{
			masterPrompt: m.MasterPrompt,
		}
	} else {
		m.TeamPresetsModal.masterPrompt = m.MasterPrompt
		m.TeamPresetsModal.modal.Show()
	}

	//})
}

func (m *MainView) OnNav(ctx app.Context) {
	m.ctx = ctx
}

// Todo: Divide into more files
func (m *MainView) Render() app.UI {
	if m.MasterPrompt == nil {
		return app.Div().Text("Loading...").Class("flex flex-col items-center justify-center h-screen")
	}

	renderedMasterPrompt, err := m.MasterPrompt.String()
	if err != nil {
		slog.Error("rendering master prompt", slog.Any("err", err))
	}

	onClickBreadCrumb := func(ctx app.Context, index int) {
		state.SetCurrentPageIndex(ctx, index)
	}
	pagesAmount := 6

	return app.Div().Class("bg-base-200 h-full min-h-dvh vw-100 p-12").Attr("data-theme", "cupcake").Class("").Body(
		app.Div().Class("w-1 h-1"), // no idea why it gets broken without it, doesnt look very harmfully
		&NavbarComponent{
			Class:          "py-2 mb-12",
			StartComponent: app.P().Text("Prompt Composer").Class("font-bold text-accent-content text-md hidden xl:inline"),
			CenterComponent: &BreadcrumbsComponent{
				Breadcrumbs: []Breadcrumb{
					{Title: "Introduction", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 0) }, Active: m.CurrentPageIndex == 0, Completed: true},
					{Title: "Template", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 1) }, Active: m.CurrentPageIndex == 1, Completed: m.CurrentPageIndex >= 1},
					{Title: "Style", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 2) }, Active: m.CurrentPageIndex == 2, Completed: m.CurrentPageIndex >= 2},
					{Title: "Rules", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 3) }, Active: m.CurrentPageIndex == 3, Completed: m.CurrentPageIndex >= 3},
					{Title: "Team", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 4) }, Active: m.CurrentPageIndex == 4, Completed: m.CurrentPageIndex >= 4},
					{Title: "Prompt ready", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 5) }, Active: m.CurrentPageIndex == 5, Completed: m.CurrentPageIndex >= 5},
				},
			},
			EndComponent: app.Div().Class("flex flex-row hidden xl:inline").Body(
				m.renderShareButton("Share workspace"),
			),
		},

		app.If(m.updateAvailable, func() app.UI {
			return app.Button().
				Text("Update app!").
				Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded").
				OnClick(m.onUpdateClick)
		}),
		&PageViewComponent{
			CurrentIndex: m.CurrentPageIndex,
			Pages: []app.UI{
				m.renderWelcomeCard(),
				m.renderMasterPromptTemplate(),
				m.renderStyle(),
				m.renderRules(),
				m.renderTeam(),
				app.Div().Body(
					m.renderOutputSection(renderedMasterPrompt),
					m.renderSharingFeature(),
				),
			},
		},

		app.If(m.TeamPresetsModal != nil, func() app.UI {
			return m.TeamPresetsModal
		}),

		&KeyListenerComponent{
			IgnoreInsideTextFields: true,
			KeyPressedAt:           m.KeystrokePressedAt,
			OnKeyUp: func(key string) {
				if key == "ArrowLeft" {
					if m.CurrentPageIndex > 0 {
						state.SetCurrentPageIndex(m.ctx, m.CurrentPageIndex-1)
					}
				}
				if key == "ArrowRight" {
					if m.CurrentPageIndex < pagesAmount-1 {
						state.SetCurrentPageIndex(m.ctx, m.CurrentPageIndex+1)
					}
				}
			},
		},
	)
}

func (m *MainView) renderScrollToBottomButton() app.HTMLButton {
	return app.Button().
		Class("btn btn-circle btn-xl btn-accent bottom-5 right-5 fixed shadow-xl").
		OnClick(func(ctx app.Context, e app.Event) {
			var outputSection = app.Window().Get("document").Call("getElementById", "output-section")

			if outputSection == nil {
				slog.Error("Unable to find output-section")
				return
			}

			outputSection.Call("scrollIntoView", map[string]interface{}{
				"behavior": "smooth",
				"block":    "start",
				"inline":   "nearest",
			})
		}).
		Body(
			&SVGIcon{
				IconData:       AngleDownIcon,
				Color:          "black",
				OpacityPercent: 45,
			},
		)
}

func (m *MainView) renderWelcomeCard() *CardComponent {
	return &CardComponent{
		Body: []app.UI{
			app.H2().Text("Tutorial").Class("text-xl font-bold mb-4"),
			app.P().Text("This is a tool to help you generate a master prompt for your LLM agent.").Class("text-md opacity-80 mb-1"),
			app.P().Text("Data is stored in your browser, so you won't lose anything after refresh").Class("text-md opacity-80 mb-12"),
			app.P().Text("Navigate by clicking links at top").Class("text-md opacity-80 mb-2"),
			app.P().Text("You can also use arrows via your keyboard!").Class("text-md opacity-80 font-bold mb-6"),
			&StepsComponent{
				IsVertical: true,
				Steps: []Step{
					{Title: "Think about your needs", Active: true},
					{Title: "Create 📜 Rules", Active: true},
					{Title: "Define 🤌🏻Style guidelines", Active: true},
					{Title: "Assign virtual 👨‍💻 Team", Active: true},
					{Title: "Copy output", Active: true},
					{Title: "Paste into Chat GPT / other LLM", Active: true},
				},
			},
		},
	}
}

func (m *MainView) renderOutputSection(renderedMasterPrompt string) *CardComponent {
	return &CardComponent{
		Class: "mb-4",
		Body: []app.UI{

			app.Div().Class("mb-4 flex flex-row").ID("output-section").Body(

				&SVGIcon{IconData: MagicWandIcon, Color: "green", IconSize: IconSizeLarge, OpacityPercent: 30},
				app.H2().Text("Your prompt").Class("ml-2 text-xl font-bold mb-1"),
			),
			app.P().Class("text-md opacity-80 mb-6").Text("Paste it into your chat gpt space, copilot or any other LLM"),

			app.Button().Class("btn btn-primary flex flex-row align-center mb-4").Body(
				&SVGIcon{IconData: CopyIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 25},
				app.P().Class("text-md").Text("Copy"),
			).OnClick(m.copyOutputPressed()),
			app.Textarea().Class("textarea textarea-bordered h-80 w-full").Text(renderedMasterPrompt).Placeholder("There should be your prompt"),
		},
	}
}

func (m *MainView) renderSharingFeature() *CardComponent {
	return &CardComponent{
		Body: []app.UI{
			app.H2().Text("Sharing workspace").Class("text-xl font-bold mb-4"),
			app.P().Class("text-md opacity-80 mb-4").Text("You can share your workspace with others by sending them a link. Just click copy and send it to your mate, you can also store it somewhere in notes and manage versions this way."),
			m.renderShareButton("Copy link"),
		},
	}
}

func (m *MainView) renderShareButton(text string) app.HTMLButton {
	return app.Button().Class("btn btn-secondary flex flex-row align-center").Body(
		&SVGIcon{IconData: LinkIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 30},
		app.P().Class("text-md").Text(text),
	).OnClick(m.copyLinkPressed())
}

func (m *MainView) copyOutputPressed() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		out, err := m.MasterPrompt.String()
		if err != nil {
			return
		}

		app.Window().Get("navigator").Get("clipboard").Call("writeText", out)
	}
}

func (m *MainView) copyLinkPressed() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		recipeBase64, err := m.MasterPrompt.ToBase64()
		if err != nil {
			slog.Error("encoding master prompt to base64", slog.Any("err", err))
			return
		}

		link := ctx.Page().URL()
		if link.Path[len(link.Path)-1] != '/' {
			link.Path += "/"
		}

		link.Path += "import"
		query := link.Query()           // Get a copy of the query parameters
		query.Set("data", recipeBase64) // Modify the query parameters
		link.RawQuery = query.Encode()

		// set clipboard
		app.Window().Get("navigator").Get("clipboard").Call("writeText", link.String())

		//alert
		app.Window().Get("alert").Invoke("Link copied to clipboard! You can now share it with your friends. or save for later")

	}
}

func (m *MainView) renderTeam() *CardComponent {
	return &CardComponent{
		Body: []app.UI{
			app.Div().Class("flex flex-row justify-between mb-4").Body(
				app.Div().Class("flex flex-col").Body(
					app.H2().Text("👨‍💻 Team").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("Describe your team members, their roles and features"),
				),
				app.Button().Class("btn btn-primary flex align-center mb-4").Body(
					&SVGIcon{IconData: MagicWandIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 30},
					app.P().Class("text-md").Text("Compose team"),
				).OnClick(func(ctx app.Context, e app.Event) {
					ctx.NewActionWithValue(actions.ActionOpenTeamPresetsModal, nil)
				}),
			),

			app.Div().Class("flex flex-col justify-stretch pr-0  w-full pb-6").
				Body(
					app.Range(m.MasterPrompt.TeamPreset.Values).Slice(func(i int) app.UI {
						var member = m.MasterPrompt.TeamPreset.Values[i]

						var handleNameChange = func(ctx app.Context, e app.Event) {
							newName := ctx.JSSrc().Get("value").String()
							m.MasterPrompt.TeamPreset.Values[i].Name = newName
							state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
						}

						var handleRoleChange = func(ctx app.Context, e app.Event) {
							newRole := ctx.JSSrc().Get("value").String()
							m.MasterPrompt.TeamPreset.Values[i].Role = newRole
							state.SetMasterPrompt(ctx, m.MasterPrompt)
						}
						emojisList := []OptionData{}

						for _, emoji := range domain.EmojiList {
							emojisList = append(emojisList, OptionData{
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

						return &CardComponent{
							Class: "mr-4 w-full rounded-3xl p-8 shadow-md border-1 border-black/15 mt-2 mb-2 bg-base-100 gap-2 mb-8",
							Body: []app.UI{
								app.Div().Class("flex flex-row align-center mb-6").Body(

									&DropdownComponent[string]{
										OptionDataList: emojisList,
										Text:           member.EmojiIcon,
										OnClick:        handleEmojiChange,
										Class:          "mr-2",
									},

									app.Input().Class("input input-md w-full mb-1").Type("text").Placeholder("Put team member name here").
										Value(member.Name).OnChange(handleNameChange).OnKeyUp(handleNameChange),
									app.If(member.Name != "", func() app.UI {
										return app.Button().ID("btn-feature-"+member.ID).Class("btn btn-error ml-2").OnClick(
											func(ctx app.Context, e app.Event) {
												theId := ctx.JSSrc().Get("id").String()
												theId = strings.ReplaceAll(theId, "btn-feature-", "")
												m.MasterPrompt.RemoveTeamMemberByID(theId)
												state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
											}).Body(
											&SVGIcon{IconData: TrashIcon, OpacityPercent: 90},
											app.P().Class("text-md text-white").Text("Remove person"),
										)
									}),
								),

								app.IfSlice(member.Name != "", func() []app.UI {
									return []app.UI{
										app.Input().Class("input input-md w-full mb-2").Type("text").Placeholder("Role = developer, customer, marketing specialist and etc..").
											Value(member.Role).OnChange(handleRoleChange).OnKeyUp(handleRoleChange),

										&Spacer{
											Size: SpacerSizeMedium,
										},
										app.Div().Class("flex flex-row justify-between mb-6").Body(
											app.H3().Class("text-xl opacity-80").Text("Features"),
											app.H3().Class("text-xl opacity-80").Text("Weight"),
										),
										app.Range(member.Features).Slice(func(j int) app.UI {
											return m.renderWeightControlledName(member.Features[j].ID, member.Features[j].Name, member.Features[j].Weight)
										}),
									}
								}),
							},
						}
					})),
		},
	}
}

func (m *MainView) renderStyle() *CardComponent {
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
						stylePreset := value
						m.MasterPrompt.StylePreset = stylePreset
						state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Style hint"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(func() []domain.Style {
					styles := make([]domain.Style, len(m.MasterPrompt.StylePreset.Values))
					copy(styles, m.MasterPrompt.StylePreset.Values)
					return styles
				}()).Slice(func(i int) app.UI {
					style := m.MasterPrompt.StylePreset.Values[i]
					return m.renderWeightControlledName(style.ID, style.Name, style.Weight)
				}),
			),
		},
		Class: "w-full",
	}
}

func (m *MainView) renderRules() *CardComponent {
	var presetsToSelect []OptionData
	for name, template := range domain.RulesPresetsMap {
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
					app.H2().Text("📜 Rules").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-80 mb-12").Text("It's great to give positive rules instead negative, like 'do this' instead 'don't do this' "),
				),
				&DropdownComponent[domain.RulePreset]{
					OptionDataList: presetsToSelect,
					Text:           "Preset",
					Icon:           &SVGIcon{IconData: SlidersIcon, Color: "black", IconSize: IconSizeMedium, OpacityPercent: 55},
					Position:       DropdownPositionLeft,
					OnClick: func(ctx app.Context, value domain.RulePreset) {
						rulesPreset := value
						m.MasterPrompt.RulePreset = rulesPreset
						state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
					},
				},
			),

			app.Div().Class("flex flex-row justify-between mb-6").Body(
				app.H3().Class("text-xl").Text("Rule"),
				app.H3().Class("text-xl").Text("Weight"),
			),
			app.Div().Class("w-full").Body(
				app.Range(func() []domain.Rule {
					rules := make([]domain.Rule, len(m.MasterPrompt.RulePreset.Values))
					copy(rules, m.MasterPrompt.RulePreset.Values)
					return rules
				}()).Slice(func(i int) app.UI {
					rule := m.MasterPrompt.RulePreset.Values[i]
					return m.renderWeightControlledName(rule.ID, rule.Name, rule.Weight)
				}),
			),
		},
		Class: "mr-4 w-full",
	}
}

func (m *MainView) renderMasterPromptTemplate() *CardComponent {
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
		m.MasterPrompt.Template = domain.MasterPromptTemplate(newTemplate)
		state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
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

						m.MasterPrompt.Template = domain.MasterPromptTemplate(value)
						state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
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

				return app.Button().ID("btn-feature-" + id).Class("btn btn-circle btn-error mr-2").OnClick(func(ctx app.Context, e app.Event) {
					theId := ctx.JSSrc().Get("id").String()
					theId = strings.ReplaceAll(theId, "btn-feature-", "")

					m.MasterPrompt.RemoveFeatureByID(theId)
					state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
				}).Body(
					&SVGIcon{
						IconData:       TrashIcon,
						Color:          "white",
						OpacityPercent: 90,
					})
			}),
			app.Div().Class("flex flex-col w-full").Body(
				app.If(name == "", func() app.UI {
					return &Spacer{Size: SpacerSizeMedium}
				}),
				app.Input().Type("text").Placeholder("Create new feature / rule").Class("input input-md w-full").Value(name).OnChange(
					func(capturedID string) func(ctx app.Context, e app.Event) {
						return func(ctx app.Context, e app.Event) {
							newName := ctx.JSSrc().Get("value").String()
							m.MasterPrompt.UpdateValueByID(capturedID, &newName, nil)
							state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
						}
					}(id)),
			),
		),

		app.If(name != "", func() app.UI {
			handleSliderInput := func(capturedID string) func(ctx app.Context, e app.Event) {
				return func(ctx app.Context, e app.Event) {
					newWeight, err := strconv.Atoi(ctx.JSSrc().Get("value").String())
					if err != nil {
						slog.Error("converting weight to int", slog.Any("err", err))
						return
					}
					m.MasterPrompt.UpdateValueByID(capturedID, nil, &newWeight)
					state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
				}
			}(id)

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

func (m *MainView) HandleAddPerson(ctx app.Context, action app.Action) {
	value := action.Value
	if reflect.TypeOf(value) == reflect.TypeOf(domain.Person{}) {
		person := action.Value.(domain.Person)
		m.MasterPrompt.TeamPreset.Values = append(m.MasterPrompt.TeamPreset.Values, person)
		state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
		m.TeamPresetsModal.masterPrompt = m.MasterPrompt
	} else {
		slog.Error("Invalid type for value", "type", reflect.TypeOf(value))
		return
	}
}
func (m *MainView) HandleRemovePerson(context app.Context, action app.Action) {
	value := action.Value
	if reflect.TypeOf(value) == reflect.TypeOf(domain.Person{}) {
		person := action.Value.(domain.Person)

		m.MasterPrompt.RemoveTeamMemberByID(person.ID)

		state.SetMasterPromptWithEmptyField(context, m.MasterPrompt)
		m.TeamPresetsModal.masterPrompt = m.MasterPrompt
	} else {
		slog.Error("Invalid type for value", "type", reflect.TypeOf(value))
		return
	}
}
