package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type PageViewComponent struct {
	app.Compo
	CurrentIndex int
	Pages        []app.UI
	Class        string
}

func (p *PageViewComponent) Render() app.UI {
	if p.CurrentIndex < 0 {
		p.CurrentIndex = 0
	}
	if p.CurrentIndex >= len(p.Pages) {
		p.CurrentIndex = len(p.Pages) - 1
	}

	return app.Div().Class(fmt.Sprintf("page-view %s", p.Class)).Body(
		p.Pages[p.CurrentIndex],
	)
}
