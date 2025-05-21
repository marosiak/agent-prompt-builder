package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log/slog"
)

type ModalComponent struct {
	app.Compo

	ID    string
	Title string
	Subtitle string
	Body     []app.UI
	Class    string

	ForceShowOnMount bool
}

func (m *ModalComponent) OnMount(ctx app.Context) {
	if m.ID == "" {
		slog.Error("Modal ID is empty")
		return
	}
}

func (m *ModalComponent) Show() {
	modal := app.Window().Get("document").Call("getElementById", m.ID)
	if !modal.IsNull() {
		modal.Call("showModal")
	} else {
		slog.Error("Modal not found", "ID", m.ID)
	}
}

func (m *ModalComponent) Hide() {
	modal := app.Window().Get("document").Call("getElementById", m.ID)
	if !modal.IsNull() {
		modal.Call("close")
	}
}
func (m *ModalComponent) Render() app.UI {
	openAttr := "open"
	if !m.ForceShowOnMount {
		openAttr = ""
	}
	return app.Dialog().ID(m.ID).Attr(openAttr, "true").Class("modal").Body(
		app.Div().Class(fmt.Sprintf("modal-box %s", m.Class)).Body(
			app.Form().Method("dialog").Body(
				app.Button().Class("btn btn-sm btn-circle absolute right-2 top-2").Type("button").OnClick(func(ctx app.Context, e app.Event) {
					m.Hide()
				}).Text("✕")),
			app.H3().Class("text-lg font-bold").Text(m.Title),
			app.P().Class("py-4").Text(m.Subtitle),
			app.Span().Body(m.Body...),
		),
		app.Form().Method("dialog").Class("modal-backdrop").Body(
			app.Button().Text("close"),
		),
	)
}
