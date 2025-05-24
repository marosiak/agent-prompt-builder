package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type StatusColor string

const (
	StatusColorPrimary   StatusColor = "status-primary"
	StatusColorSecondary StatusColor = "status-secondary"
	StatusColorAccent    StatusColor = "status-accent"
	StatusColorNeutral   StatusColor = "status-neutral"
	StatusColorInfo      StatusColor = "status-info"
	StatusColorSuccess   StatusColor = "status-success"
	StatusColorWarning   StatusColor = "status-warning"
	StatusColorError     StatusColor = "status-error"
)

type StatusSize string

const (
	StatusSizeXSmall StatusSize = "status-xs"
	StatusSizeSmall  StatusSize = "status-sm"
	StatusSizeMedium StatusSize = "status-md"
	StatusSizeLarge  StatusSize = "status-lg"
	StatusSizeXLarge StatusSize = "status-xl"
)

type StatusComponent struct {
	app.Compo
	Color StatusColor
	Size  StatusSize
	Label app.UI
}

func (s *StatusComponent) Render() app.UI {
	if s.Label != nil {
		return app.Div().Class("flex flex-row items-center gap-2").Body(
			s.renderIndicator(),
			s.Label,
		)
	}
	return s.renderIndicator()
}

func (s *StatusComponent) renderIndicator() app.HTMLDiv {
	return app.Div().
		Class("status "+string(s.Size)).
		Class(string(s.Color)).
		Aria("label", "status")
}
