package components

import (
	"github.com/marosiak/agent-prompt-builder/state"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// NavigationComponent handles the breadcrumb navigation
type NavigationComponent struct {
	app.Compo
	CurrentPageIndex int
}

func (n *NavigationComponent) Render() app.UI {
	onClickBreadCrumb := func(ctx app.Context, index int) {
		state.SetCurrentPageIndex(ctx, index)
	}

	return &NavbarComponent{
		Class:          "py-2 mb-12",
		StartComponent: app.P().Text("Prompt Composer").Class("font-bold text-accent-content text-md hidden xl:inline"),
		CenterComponent: &BreadcrumbsComponent{
			Breadcrumbs: []Breadcrumb{
				{Title: "Introduction", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 0) }, Active: n.CurrentPageIndex == 0, Completed: true},
				{Title: "Template", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 1) }, Active: n.CurrentPageIndex == 1, Completed: n.CurrentPageIndex >= 1},
				{Title: "Style", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 2) }, Active: n.CurrentPageIndex == 2, Completed: n.CurrentPageIndex >= 2},
				{Title: "Rules", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 3) }, Active: n.CurrentPageIndex == 3, Completed: n.CurrentPageIndex >= 3},
				{Title: "Team", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 4) }, Active: n.CurrentPageIndex == 4, Completed: n.CurrentPageIndex >= 4},
				{Title: "Prompt ready", OnClick: func(ctx app.Context, e app.Event) { onClickBreadCrumb(ctx, 5) }, Active: n.CurrentPageIndex == 5, Completed: n.CurrentPageIndex >= 5},
			},
		},
		EndComponent: app.Div().Class("flex flex-row hidden xl:inline").Body(
			n.renderShareButton(),
		),
	}
}

func (n *NavigationComponent) renderShareButton() app.HTMLButton {
	return app.Button().Class("btn btn-secondary flex flex-row align-center").Body(
		&SVGIcon{IconData: LinkIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 30},
		app.P().Class("text-md").Text("Share workspace"),
	).OnClick(func(ctx app.Context, e app.Event) {
		// Get masterPrompt from state
		masterPrompt := state.GetMasterPrompt(ctx)
		recipeBase64, err := masterPrompt.ToBase64()
		if err != nil {
			return
		}

		link := ctx.Page().URL()
		if link.Path[len(link.Path)-1] != '/' {
			link.Path += "/"
		}

		link.Path += "import"
		query := link.Query()
		query.Set("data", recipeBase64)
		link.RawQuery = query.Encode()

		app.Window().Get("navigator").Get("clipboard").Call("writeText", link.String())
		app.Window().Get("alert").Invoke("Link copied to clipboard! You can now share it with your friends. or save for later")
	})
}
