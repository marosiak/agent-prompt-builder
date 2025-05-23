package views

import (
	"log/slog"
	"reflect"
	"time"

	"github.com/marosiak/agent-prompt-builder/actions"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
	. "github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
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
}

func (m *MainView) OnNav(ctx app.Context) {
	m.ctx = ctx
}

func (m *MainView) Render() app.UI {
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
		app.If(m.updateAvailable, func() app.UI {
			return app.Button().
				Text("Update app!").
				Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded").
				OnClick(m.onUpdateClick)
		}),

		// Main content pages
		&PageViewComponent{
			CurrentIndex: m.CurrentPageIndex,
			Pages: []app.UI{
				&WelcomeCardComponent{},
				&TemplateEditorComponent{MasterPrompt: *m.MasterPrompt},
				&StyleEditorComponent{MasterPrompt: *m.MasterPrompt},
				&RulesEditorComponent{MasterPrompt: *m.MasterPrompt},
				&TeamEditorComponent{MasterPrompt: *m.MasterPrompt},
				&OutputViewComponent{MasterPrompt: *m.MasterPrompt},
			},
		},

		// Team presets modal
		app.If(m.TeamPresetsModal != nil, func() app.UI {
			return m.TeamPresetsModal
		}),

		// Keyboard navigation
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
	}
}
