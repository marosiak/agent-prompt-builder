package components

import "github.com/maxence-charriere/go-app/v10/pkg/app"

// Step represents a single step in the StepsComponent component.
type Step struct {
	Title    string
	Active   bool
	Complete bool
}

// StepsComponent is a versatile and type-safe component for rendering steps.
type StepsComponent struct {
	app.Compo
	Steps      []Step // A slice of Step structs to define the steps dynamically.
	IsVertical bool
}

func (s *StepsComponent) Render() app.UI {
	ul := app.Ul().Class("steps")
	if s.IsVertical {
		ul.Class("steps-vertical")
	} else {
		ul.Class("steps-horizontal")
	}

	// Ensure all steps are rendered with the correct classes
	var liList []app.UI
	for _, step := range s.Steps {
		li := app.Li().Class("step")
		if step.Complete {
			li.Class("step-primary")
		} else if step.Active {
			li.Class("step-active")
		}
		li.Body(app.Text(step.Title))
		liList = append(liList, li)

	}

	ul.Body(liList...) // Append each step to the unordered list
	return ul
}
