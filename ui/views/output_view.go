package views

import (
	"github.com/marosiak/agent-prompt-builder/actions"
	. "github.com/marosiak/agent-prompt-builder/ui/components"
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

			&ButtonComponent{
				DefaultState: ButtonState{
					Text: "Copy",
					Icon: &SVGIcon{
						IconData:       CopyIcon,
						Color:          "#9f2f0f",
						OpacityPercent: 40,
					},
					Color: ButtonColorAccent,
				},
				PostClickState: &ButtonState{
					Text:           "Done!",
					Color:          ButtonColorSuccess,
					AnimationClass: "animate-pulse",
					Icon: &SVGIcon{
						IconData:       CircleCheckIcon,
						Color:          "white",
						OpacityPercent: 50,
					},
				},
				Class:   "mb-4",
				OnClick: o.copyOutputPressed(),
			},
			app.Textarea().Class("textarea textarea-bordered h-80 w-full").Text(renderedMasterPrompt).Placeholder("There should be your prompt"),
		},
	}
}

func (o *OutputViewComponent) renderSharingFeature() *CardComponent {
	return &CardComponent{
		Body: []app.UI{
			app.H2().Text("Sharing workspace").Class("text-xl font-bold mb-4"),
			app.P().Class("text-md opacity-80 mb-4").Text("You can share your workspace with others by sending them a link. Just click copy and send it to your mate, you can also store it somewhere in notes and manage versions this way."),
			o.renderCopyShareLinkButton(),
		},
	}
}

func (o *OutputViewComponent) renderCopyShareLinkButton() app.UI {
	return &ButtonComponent{
		DefaultState: ButtonState{
			Text: "Copy link",
			Icon: &SVGIcon{
				IconData:       CopyIcon,
				Color:          "#9f2f0f",
				OpacityPercent: 40,
			},
			Color: ButtonColorSecondary,
		},
		PostClickState: &ButtonState{
			Text:           "Done!",
			Color:          ButtonColorSuccess,
			AnimationClass: "animate-pulse",
			Icon: &SVGIcon{
				IconData:       CircleCheckIcon,
				Color:          "white",
				OpacityPercent: 50,
			},
		},
		OnClick: o.copyLinkPressed(),
	}
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
		ctx.NewAction(actions.ShareWorkspace)
	}
}
