package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type CollapsableComponent struct {
	app.Compo
	Title string
	Body  []app.UI
}

func (c *CollapsableComponent) Render() app.UI {
	return app.Div().TabIndex(0).Class("collapse collapse-arrow bg-base-100 border-base-300 border").Body(
		app.Div().Class("collapse-title font-semibold").Body(
			app.Text(c.Title),
		),
		app.Div().Class("collapse-content text-sm").Body(
			c.Body...,
		),
	)
}
