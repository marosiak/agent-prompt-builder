package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// ListComponent renders a styled list with customizable items
type ListComponent struct {
	app.Compo
	Title          string     // Optional title for the list
	Items          []ListItem // List of items to display
	Class          string     // Additional CSS classes for the list container
	DisableShadows bool
}

// ListItem represents a single item in the list
type ListItem struct {
	Leading       app.UI // Optional leading element (typically an image)
	Title         string // Main title text
	Subtitle      string // Secondary text shown below the title
	ContentString string // Optional descriptive content text
	Content       []app.UI
	Trailing      []app.UI // Optional trailing elements (typically action buttons)
	Class         string   // Additional CSS classes for this item
	ContentClass  string
}

// Render implements the app.UI interface
func (l *ListComponent) Render() app.UI {
	class := "list bg-base-100 rounded-box shadow-md " + l.Class
	if l.DisableShadows {
		class = "list bg-base-100 rounded-box " + l.Class
	}
	return app.Ul().Class(class).Body(
		app.If(l.Title != "", func() app.UI {
			return app.Li().Class("p-4 pb-2 text-xs opacity-60 tracking-wide").Text(l.Title)
		}),
		app.If(l.Items != nil, func() app.UI {
			return app.Range(l.Items).Slice(func(i int) app.UI {
				item := l.Items[i]
				return app.Li().Class("list-row "+item.Class).Body(
					app.If(item.Leading != nil, func() app.UI {
						return app.Div().Body(item.Leading)
					}),
					app.If(item.Title != "" || item.Subtitle != "", func() app.UI {
						return app.Div().Body(
							app.If(item.Title != "", func() app.UI {
								return app.Div().Text(item.Title)
							}),
							app.If(item.Subtitle != "", func() app.UI {
								return app.Div().Class("text-xs uppercase font-semibold opacity-60").Text(item.Subtitle)
							}),
						)
					}),

					app.Div().Class("list-col-wrap").Body(

						app.If(item.ContentString != "", func() app.UI {
							return app.P().Class(fmt.Sprintf("text-xs %s", item.ContentClass)).Text(item.ContentString)
						}),

						app.IfSlice(len(item.Content) > 0, func() []app.UI {
							return item.Content
						}),
					),

					app.If(item.Trailing != nil && len(item.Trailing) > 0, func() app.UI {
						return app.Range(item.Trailing).Slice(func(j int) app.UI {
							return item.Trailing[j]
						})
					}),
				)
			})
		}),
	)
}
