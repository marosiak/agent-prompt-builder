package components

import (
	"fmt"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"reflect"
)

type DropdownComponent[T any] struct {
	app.Compo
	OptionDataList   []OptionData
	Icon             *SVGIcon
	OnClick          func(ctx app.Context, value T)
	Class            string
	Text     string
	Position DropdownPosition
}

type DropdownPosition string

const (
	DropdownPositionNone   DropdownPosition = ""
	DropdownPositionTop    DropdownPosition = "dropdown-top"
	DropdownPositionBottom DropdownPosition = "dropdown-bottom"
	DropdownPositionLeft   DropdownPosition = "dropdown-left"
	DropdownPositionRight  DropdownPosition = "dropdown-right"
	DropdownPositionCenter DropdownPosition = "dropdown-center"
)

func (d *DropdownComponent[T]) onClick(value any) func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		if d.OnClick == nil {
			return
		}
		v := value

		var output any
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			output = v.(string)
		case reflect.Int:
			output = v.(int)
		case reflect.Bool:
			output = v.(bool)
		default:
			output = v.(T)
		}

		d.OnClick(ctx, output.(T))
	}
}

func (d *DropdownComponent[T]) Render() app.UI {
	return app.Div().Class(fmt.Sprintf("dropdown dropdown-hover %s %s", d.Position, d.Class)).Body(

		app.Div().TabIndex(0).Role("button").Class("btn align-center").Body(

			app.If(d.Icon != nil, func() app.UI {
				return d.Icon
			}),
			app.Text(d.Text),
		),

		app.Ul().TabIndex(0).Class("dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm").Body(
			app.Range(d.OptionDataList).Slice(func(i int) app.UI {
				return app.Li().OnClick(d.onClick(d.OptionDataList[i].Value)).Body(
					app.A().Text(d.OptionDataList[i].Label),
				)
			}),
		),
	)
}
