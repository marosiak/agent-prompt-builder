package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type NavbarComponent struct {
	app.Compo
	StartComponent  app.UI
	CenterComponent app.UI
	EndComponent    app.UI
	Class           string
}

func (n *NavbarComponent) Render() app.UI {
	return app.Div().Class(fmt.Sprintf("navbar bg-base-100 shadow-xl p-8 rounded-xl w-full h-auto %s", n.Class)).Body(
		app.Div().Class("navbar-start").Body(
			n.StartComponent,
		),
		app.Div().Class("navbar-center").Body(
			n.CenterComponent,
		),
		app.Div().Class("navbar-end").Body(
			n.EndComponent,
		),
	)
}
