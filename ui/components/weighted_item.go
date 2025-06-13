package components

import (
	"log/slog"
	"strconv"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// WeightedItemComponent displays an item with an editable name and weight slider
type WeightedItemComponent struct {
	app.Compo
	ID        string
	Name      string
	Weight    int
	OnDelete  func(ctx app.Context, id string)
	OnUpdate  func(ctx app.Context, id string, name *string, weight *int)
	ShowTrash bool
}

func (w *WeightedItemComponent) Render() app.UI {
	onTextChange := func(ctx app.Context, e app.Event) {
		newName := ctx.JSSrc().Get("value").String()
		if w.OnUpdate != nil {
			w.OnUpdate(ctx, w.ID, &newName, nil)
		}
	}
	return app.Div().Class("flex flex-row w-full justify-between items-center mb-2").Body(
		app.Div().Class("flex flex-row w-full").Body(
			app.If(w.ShowTrash, func() app.UI {
				return app.Button().ID("btn-feature-" + w.ID).Class("btn btn-circle btn-error mr-2").OnClick(func(ctx app.Context, e app.Event) {
					if w.OnDelete != nil {
						w.OnDelete(ctx, w.ID)
					}
				}).Body(
					&SVGIcon{
						IconData:       TrashIcon,
						Color:          "white",
						OpacityPercent: 90,
					})
			}),
			app.Div().Class("flex flex-col w-full").Body(
				app.If(w.Name == "", func() app.UI {
					return &Spacer{Size: SpacerSizeMedium}
				}),
				app.Input().Type("text").Placeholder("Create new feature / rule").Class("input input-md w-full").Value(w.Name).
					OnChange(onTextChange),
			),
		),

		app.If(w.Name != "", func() app.UI {
			handleSliderInput := func(ctx app.Context, e app.Event) {
				newWeight, err := strconv.Atoi(ctx.JSSrc().Get("value").String())
				if err != nil {
					slog.Error("converting weight to int", slog.Any("err", err))
					return
				}
				if w.OnUpdate != nil {
					w.OnUpdate(ctx, w.ID, nil, &newWeight)
				}
			}

			return app.Input().Class("w-96 ml-4").Type("range").Min(0).Max(100).Class("range").Value(w.Weight).
				OnChange(handleSliderInput).
				OnPaste(handleSliderInput).
				OnDrag(handleSliderInput).OnMouseMove(handleSliderInput)
		}),
	)
}
