package components

import (
	"log/slog"

	"github.com/marosiak/agent-prompt-builder/domain"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// OutputViewComponent displays the generated prompt and sharing options
type OutputViewComponent struct {
	app.Compo
	MasterPrompt domain.MasterPrompt
}

func (o *OutputViewComponent) Render() app.UI {
	renderedMasterPrompt, err := o.MasterPrompt.String()
	if err != nil {
		slog.Error("rendering master prompt", slog.Any("err", err))
		renderedMasterPrompt = "Error generating prompt"
	}

	return app.Div().Body(
		o.renderOutputSection(renderedMasterPrompt),
		o.renderSharingFeature(),
	)
}

func (o *OutputViewComponent) renderOutputSection(renderedMasterPrompt string) *CardComponent {
	return &CardComponent{
		Class: "mb-4",
		Body: []app.UI{
			app.Div().Class("mb-4 flex flex-row").ID("output-section").Body(
				&SVGIcon{IconData: MagicWandIcon, Color: "green", IconSize: IconSizeLarge, OpacityPercent: 30},
				app.H2().Text("Your prompt").Class("ml-2 text-xl font-bold mb-1"),
			),
			app.P().Class("text-md opacity-80 mb-6").Text("Paste it into your chat gpt space, copilot or any other LLM"),

			app.Button().Class("btn btn-primary flex flex-row align-center mb-4").Body(
				&SVGIcon{IconData: CopyIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 25},
				app.P().Class("text-md").Text("Copy"),
			).OnClick(o.copyOutputPressed()),
			app.Textarea().Class("textarea textarea-bordered h-80 w-full").Text(renderedMasterPrompt).Placeholder("There should be your prompt"),
		},
	}
}

func (o *OutputViewComponent) renderSharingFeature() *CardComponent {
	return &CardComponent{
		Body: []app.UI{
			app.H2().Text("Sharing workspace").Class("text-xl font-bold mb-4"),
			app.P().Class("text-md opacity-80 mb-4").Text("You can share your workspace with others by sending them a link. Just click copy and send it to your mate, you can also store it somewhere in notes and manage versions this way."),
			o.renderShareButton("Copy link"),
		},
	}
}

func (o *OutputViewComponent) renderShareButton(text string) app.HTMLButton {
	return app.Button().Class("btn btn-secondary flex flex-row align-center").Body(
		&SVGIcon{IconData: LinkIcon, Color: "black", IconSize: IconSizeBig, OpacityPercent: 30},
		app.P().Class("text-md").Text(text),
	).OnClick(o.copyLinkPressed())
}

func (o *OutputViewComponent) copyOutputPressed() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		out, err := o.MasterPrompt.String()
		if err != nil {
			return
		}

		app.Window().Get("navigator").Get("clipboard").Call("writeText", out)
	}
}

func (o *OutputViewComponent) copyLinkPressed() func(ctx app.Context, e app.Event) {
	return func(ctx app.Context, e app.Event) {
		recipeBase64, err := o.MasterPrompt.ToBase64()
		if err != nil {
			slog.Error("encoding master prompt to base64", slog.Any("err", err))
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
	}
}
