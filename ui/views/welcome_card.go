package views

import (
	. "github.com/marosiak/agent-prompt-builder/ui/components"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// WelcomeCardComponent displays the welcome information
type WelcomeCardComponent struct {
	app.Compo
}

func (w *WelcomeCardComponent) Render() app.UI {
	return &CardComponent{
		Body: []app.UI{
			app.H2().Text("Introduction").Class("text-xl font-bold mb-4"),
			app.P().Text("This is a tool to help you generate a master prompt for your LLM agent.").Class("text-md opacity-80 mb-1"),
			app.P().Text("Data is stored in your browser, so you won't lose anything after refresh").Class("text-md opacity-80 mb-12"),

			&CardComponent{
				Class: "w-112",
				Body: []app.UI{
					app.H2().Text("Why should I use 'Role-Prompting' technique?").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-75").Text("It's easy to understand such prompt for both human and AI - because we're familiar with roles in our life."),
					app.P().Class("text-md opacity-100 mt-2").Text("The results of your prompts will be more consistent and predictable."),
					app.A().Href("https://www.prompthub.us/blog/role-prompting-does-adding-personas-to-your-prompts-really-make-a-difference").Text("Read more").Class("link link-info").Target("_blank"),
				},
			},
			&Spacer{},
			&CardComponent{
				Class: "w-112",
				Body: []app.UI{
					app.H2().Text("How to start?").Class("text-xl font-bold mb-4"),
					app.P().Class("text-md opacity-75").Text("Navigate by clicking links in navigation or keyboard arrows."),
					app.P().Class("text-md opacity-100 mt-2").Text("For building persona I highly recommend to use 'Compose team' button - it will let you import pre-defined roles"),
				},
			},
		},
	}
}
