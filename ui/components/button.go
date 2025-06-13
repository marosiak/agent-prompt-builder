package components

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"time"
)

// ButtonColor represents the available colors for buttons
type ButtonColor string

const (
	ButtonColorPrimary   ButtonColor = "btn-primary"
	ButtonColorSecondary ButtonColor = "btn-secondary"
	ButtonColorSuccess   ButtonColor = "btn-success"
	ButtonColorError     ButtonColor = "btn-error"
	ButtonColorAccent    ButtonColor = "btn-accent"
	ButtonColorWarning   ButtonColor = "btn-warning"
	ButtonColorInfo      ButtonColor = "btn-info"
)

// ButtonState represents the state of a button
type ButtonState struct {
	Text           string
	Color          ButtonColor
	AnimationClass string
	Icon           *SVGIcon
}

// ButtonComponent represents a button with toggle capability
type ButtonComponent struct {
	app.Compo
	currentState *ButtonState

	DefaultState        ButtonState
	PostClickState      *ButtonState
	DurationOfPostClick time.Duration
	Class               string
	OnClick             app.EventHandler
}

// HandleClick processes the button click with optional temporary toggle
func (b *ButtonComponent) HandleClick(ctx app.Context, e app.Event) {
	if b.PostClickState != nil && b.currentState != b.PostClickState {
		b.currentState = b.PostClickState

		if b.DurationOfPostClick.Milliseconds() < 1 {
			b.DurationOfPostClick = time.Second
		}

		ctx.After(b.DurationOfPostClick, func(ctx app.Context) {
			b.currentState = &b.DefaultState
		})
	}

	if b.OnClick != nil {
		b.OnClick(ctx, e)
	}
}

// OnDismount cleans up any active timers
func (b *ButtonComponent) OnDismount() {
}

func (b *ButtonComponent) OnMount(ctx app.Context) {
	b.currentState = &b.DefaultState
}

func (b *ButtonComponent) Render() app.UI {
	if b.currentState == nil {
		b.currentState = &b.DefaultState
	}

	// Build the class string
	class := "btn "
	if b.currentState.Color != "" {
		class += string(b.currentState.Color)
	} else {
		class += " btn-primary" // Default color
	}

	class += " flex flex-row align-center"

	if b.Class != "" {
		class += " " + b.Class
	}

	if b.currentState.AnimationClass != "" {
		class += " " + b.currentState.AnimationClass
	}

	return app.Button().Class(class).Body(
		app.If(b.currentState.Icon != nil, func() app.UI {
			return b.currentState.Icon
		}),
		app.P().Class("text-md").Text(b.currentState.Text),
	).OnClick(b.HandleClick)
}
