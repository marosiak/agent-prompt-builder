package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

type Breadcrumb struct {
	Title     string
	OnClick   func(ctx app.Context, e app.Event)
	Active    bool
	Completed bool
}
type BreadcrumbsComponent struct {
	app.Compo
	Breadcrumbs []Breadcrumb
	Class       string
}

func (b *BreadcrumbsComponent) Render() app.UI {
	ul := app.Ul().Class("breadcrumbs text-md lg:text-xl")
	var liList []app.UI
	for _, breadcrumb := range b.Breadcrumbs {
		li := app.Li()
		if breadcrumb.Active {
			li.Class("font-bold")
		}
		if breadcrumb.OnClick != nil {
			li.Body(
				app.IfSlice(breadcrumb.Completed, func() []app.UI {
					return []app.UI{
						&SVGIcon{IconData: SquareCheckIcon, Color: "darkgreen", OpacityPercent: 60, IconSize: IconSizeBig},
						app.Div().Class("h-1 w-1"),
					}
				}),
				app.IfSlice(breadcrumb.Active, func() []app.UI {
					return []app.UI{
						app.A().OnClick(breadcrumb.OnClick).Body(app.Text(breadcrumb.Title)).Class("text-success-200 p-2 m-0 select-none"),
					}
				}).Else(
					func() app.UI {
						return app.A().OnClick(breadcrumb.OnClick).Body(app.Text(breadcrumb.Title)).Class("p-2 m-0 select-none")
					}),
			)
		} else {
			li.Body(app.Text(breadcrumb.Title))
		}
		liList = append(liList, li)
	}
	ul.Body(liList...)
	return app.Div().Class(fmt.Sprintf("breadcrumbs text-sm %s", b.Class)).Body(ul)
}
