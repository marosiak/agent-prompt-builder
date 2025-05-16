package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Spacer struct {
	app.Compo
	Size  SpacerSize
	Class string
}

type SpacerSize int

const (
	SpacerSizeNone   SpacerSize = 0
	SpacerSizeSmall  SpacerSize = 4
	SpacerSizeMedium SpacerSize = 8
	SpacerSizeLarge  SpacerSize = 12
	SpacerSizeHuge  SpacerSize = 16
)

func (s *Spacer) Render() app.UI {
	if s.Size == 0 {
		s.Size = 2
	}

	if s.Size == SpacerSizeNone {
		s.Size = SpacerSizeMedium
	}

	return app.Div().Class(fmt.Sprintf("w-full h-%d flex justify-center items-center", s.Size+2)).Body(
		app.Div().Class("w-full h-1 border-b border-dashed border-b-base-300"),
	)
	//return app.Div().Class(fmt.Sprintf("w-full h-%d border-b border-dashed border-b-base-300 mr-%d ml-%d mt-%d mb-%d %s",
	//	s.Size, s.MarginHorizontal, s.MarginHorizontal, s.MarginVertical, s.MarginVertical, s.Class))
}
