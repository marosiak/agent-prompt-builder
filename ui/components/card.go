package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type CardComponent struct {
	app.Compo
	Body    []app.UI
	Class   string
	Padding int
}

func (c *CardComponent) Render() app.UI {
	if c.Padding == 0 {
		c.Padding = 8
	}

	if c.Padding == -1 {
		c.Padding = 0 // well.. ptr is difficult to use, and didn't want to create a lot of abstraction
	}
	class := fmt.Sprintf("%s rounded-3xl p-%d shadow-sm border-1 border-black/15 mt-2 mb-2 bg-base-100 gap-2", c.Class, c.Padding)
	return app.Div().Class(class).Body(c.Body...)
}
