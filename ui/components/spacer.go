package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Spacer struct {
	app.Compo
	Size  int
	Class string
}

func (s *Spacer) Render() app.UI {
	if s.Size == 0 {
		s.Size = 2
	}
	return app.Div().Class(fmt.Sprintf("w-full h-%d border-b border-dashed border-b-base-300 m-2 %s", s.Size, s.Class))
}
