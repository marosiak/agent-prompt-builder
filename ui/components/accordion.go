package components

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// AccordionComponent renders a styled accordion with collapsible items
type AccordionComponent struct {
	app.Compo
	Items        []AccordionItem // List of accordion items
	OpenedIndex  int             // Index of initially opened item (-1 for all closed)
	MultipleOpen bool            // Whether multiple items can be open simultaneously
	Class        string          // Additional CSS classes for the accordion container
	radioGroupID string          // Unique ID for the radio group
}

// AccordionItem represents a single collapsible section in the accordion
type AccordionItem struct {
	Title         string   // Title text for the section header
	Content       []app.UI // Content elements to display when expanded
	ContentString string   // Optional simple text content
	TitleClass    string   // Additional CSS classes for the title
	ContentClass  string   // Additional CSS classes for the content
}

// OnMount initializes the component when it's mounted to the DOM
func (a *AccordionComponent) OnMount(ctx app.Context) {
	ctx.Dispatch(func(ctx app.Context) {
		a.radioGroupID = fmt.Sprintf("accordion-%s", uuid.NewString())
	})

}

// Render implements the app.UI interface
func (a *AccordionComponent) Render() app.UI {
	// Ensure we have a radio group ID even if OnMount hasn't run yet
	if a.radioGroupID == "" {
		a.radioGroupID = fmt.Sprintf("accordion-%s", uuid.NewString())
	}

	return app.Div().Class("join join-vertical bg-base-100 " + a.Class).Body(
		app.If(a.Items != nil, func() app.UI {
			return app.Range(a.Items).Slice(func(i int) app.UI {
				item := a.Items[i]

				// Determine input type based on whether multiple items can be open
				inputType := "radio"
				if a.MultipleOpen {
					inputType = "checkbox"
				}

				return app.Div().Class("collapse collapse-arrow join-item border-base-300 border").Body(
					app.Input().
						Type(inputType).
						Name(a.radioGroupID).
						Checked(i == a.OpenedIndex),
					app.Div().
						Class("collapse-title font-semibold "+item.TitleClass).
						Text(item.Title),
					app.Div().
						Class("collapse-content text-sm "+item.ContentClass).
						Body(
							app.If(item.ContentString != "", func() app.UI {
								return app.P().Text(item.ContentString)
							}),
							app.IfSlice(len(item.Content) > 0, func() []app.UI {
								return item.Content
							}),
						),
				)
			})
		}),
	)
}
