package views

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log/slog"
	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/marosiak/agent-prompt-builder/state"
)

type ImportView struct {
	app.Compo
	MasterPrompt *domain.MasterPrompt
}

func (i *ImportView) Render() app.UI {
	// dialog if you want to import data and about consequences
	return app.Div().Class("flex flex-col items-center justify-center h-screen").Body(
		app.Div().Class("card w-96 bg-base-100 shadow-xl").Body(
			app.Div().Class("card-body").Body(
				app.H2().Class("card-title").Text("Import Data"),
				app.P().Text("Importing data will overwrite any existing data. Please ensure you have a backup before proceeding."),
				app.B().Text("Do you want to proceed?"),
				app.Div().Class("card-actions justify-end").Body(
					app.Button().Class("btn btn-primary").Text("Yes").OnClick(i.importData()),
					app.Button().Class("btn btn-secondary").Text("No").OnClick(func(ctx app.Context, e app.Event) {
						i.redirectHome(ctx)
					},
					),
				),
			),
		),
	)

}

func (i *ImportView) redirectHome(ctx app.Context) {
	link := ctx.Page().URL()
	link.Path = "/"
	query := link.Query() // Get a copy of the query parameters
	query.Del("data")     // Modify the query parameters
	link.RawQuery = query.Encode()
	ctx.NavigateTo(link) // Navigate to the new URL
}

func (i *ImportView) importData() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		defer i.redirectHome(ctx)
		data := ctx.Page().URL().Query().Get("data")
		if data == "" {
			// TODO: Handle error
			slog.Error("No data provided for import")
			return
		}

		i.MasterPrompt = new(domain.MasterPrompt)
		err := i.MasterPrompt.FromBase64(data)
		if err != nil {
			// TODO: Handle error
			slog.Error("decode data", "error", err)
			return
		}

		state.DelMasterPrompt(ctx)
		state.SetMasterPrompt(ctx, i.MasterPrompt)
	}
}
