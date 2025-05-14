package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type CardComponent struct {
	app.Compo
	Body  []app.UI
	Class string
}

func (c *CardComponent) Render() app.UI {
	return app.Div().Class(c.Class + " rounded-3xl p-8 shadow-sm border-1 border-black/15 mt-2 mb-2").Body(c.Body...)
}
