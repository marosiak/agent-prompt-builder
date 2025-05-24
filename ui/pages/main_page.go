package pages

import (
	"github.com/marosiak/agent-prompt-builder/ui/views"
	"log/slog"
	"reflect"
	"time"

	"github.com/marosiak/agent-prompt-builder/actions"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	. "github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type MainPage struct {
	app.Compo
	AppUpdateAvailable bool
	MasterPrompt       *domain.MasterPrompt
	CurrentPageIndex   int
	ctx                app.Context
	KeystrokePressedAt *time.Time

	TeamPresetsModal *views.TeamPresetsModal
}

func (m *MainPage) OnAppUpdate(ctx app.Context) {
	m.AppUpdateAvailable = ctx.AppUpdateAvailable()
}

func (m *MainPage) OnMount(ctx app.Context) {
	m.ctx = ctx
	masterPromptCopy := state.GetMasterPrompt(ctx)
	m.MasterPrompt = &masterPromptCopy
	m.CurrentPageIndex = state.GetCurrentPageIndex(ctx)

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
	ctx.Handle(actions.ShareWorkspace, m.HandleShareWorkspace)
	ctx.Handle(actions.RemovePerson, m.HandleRemovePerson)
}

func (m *MainPage) HandleOpenTeamPresetsModal(ctx app.Context, a app.Action) {
	if m.TeamPresetsModal == nil {
		m.TeamPresetsModal = &views.TeamPresetsModal{
			MasterPrompt: m.MasterPrompt,
		}
	} else {
		m.TeamPresetsModal.MasterPrompt = m.MasterPrompt
		m.TeamPresetsModal.Modal.Show()
	}
}

func (m *MainPage) OnNav(ctx app.Context) {
	m.ctx = ctx
}

func (m *MainPage) Render() app.UI {
	if m.MasterPrompt == nil {
		return app.Div().Text("Loading...").Class("flex flex-col items-center justify-center h-screen")
	}

	pagesAmount := 6

	return app.Div().Class("bg-base-200 h-full min-h-dvh vw-100 p-12").Attr("data-theme", "cupcake").Body(
		app.Div().Class("w-1 h-1"), // Placeholder div for layout stability

		// Navigation component
		&NavigationComponent{
			CurrentPageIndex: m.CurrentPageIndex,
		},

		// Update notification
		app.If(m.AppUpdateAvailable, func() app.UI {
			return &ModalComponent{
				ID:       "update-app-modal",
				Title:    "New update available!",
				Subtitle: "Just click button and it will be installed in friction of seconds!",
				Body: []app.UI{
					app.Div().Class("flex flex-row-reverse").Body(
						&ButtonComponent{
							DefaultState: ButtonState{
								Text: "Update app",
								Icon: &SVGIcon{
									IconData:       RefreshIcon,
									Color:          "#277780",
									OpacityPercent: 40,
								},
							},
							OnClick: m.onUpdateClick,
							Class:   "mt-4",
						},
					),
				},
				ForceShowOnMount: true,
			}
		}),

		// Main content pages
		&PageViewComponent{
			CurrentIndex: m.CurrentPageIndex,
			Pages: []app.UI{
				&views.WelcomeCardComponent{},
				&views.TemplateEditorComponent{MasterPrompt: *m.MasterPrompt},
				&views.StyleEditorComponent{MasterPrompt: *m.MasterPrompt},
				&views.RulesEditorComponent{MasterPrompt: *m.MasterPrompt},
				&views.TeamEditorComponent{MasterPrompt: *m.MasterPrompt},
				&views.OutputViewComponent{MasterPrompt: *m.MasterPrompt},
			},
		},

		// Team presets modal
		app.If(m.TeamPresetsModal != nil, func() app.UI {
			return m.TeamPresetsModal
		}),

		// Keyboard navigation
		// Todo: implementing actions may be better, and ideally this component shouldn't be component, just some piece of code ran in different goroutine
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

func (m *MainPage) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}

func (m *MainPage) HandleAddPerson(ctx app.Context, action app.Action) {
	value := action.Value
	if reflect.TypeOf(value) == reflect.TypeOf(domain.Person{}) {
		person := action.Value.(domain.Person)
		m.MasterPrompt.TeamPreset.Values = append(m.MasterPrompt.TeamPreset.Values, person)
		state.SetMasterPromptWithEmptyField(ctx, m.MasterPrompt)
		m.TeamPresetsModal.MasterPrompt = m.MasterPrompt
	} else {
		slog.Error("Invalid type for value", "type", reflect.TypeOf(value))
	}
}

func (m *MainPage) HandleShareWorkspace(ctx app.Context, action app.Action) {
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
	query := link.Query()
	query.Set("data", recipeBase64)
	link.RawQuery = query.Encode()

	app.Window().Get("navigator").Get("clipboard").Call("writeText", link.String())
}

func (m *MainPage) HandleRemovePerson(context app.Context, action app.Action) {
	value := action.Value
	if reflect.TypeOf(value) == reflect.TypeOf(domain.Person{}) {
		person := action.Value.(domain.Person)
		m.MasterPrompt.RemoveTeamMemberByID(person.ID)
		state.SetMasterPromptWithEmptyField(context, m.MasterPrompt)
		m.TeamPresetsModal.MasterPrompt = m.MasterPrompt
	} else {
		slog.Error("Invalid type for value", "type", reflect.TypeOf(value))
	}
}
