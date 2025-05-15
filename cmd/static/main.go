package main

import (
	"github.com/google/uuid"
	"github.com/marosiak/agent-prompt-builder/ui/views"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
)

func main() {
	// TODO fix DRY
	app.Route("/", func() app.Composer {
		return &views.MainView{}
	})

	app.Route("/import", func() app.Composer {
		return &views.ImportView{}
	})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Master prompt builder",
		Description: "Will help you with building agents",
		Resources:   app.GitHubPages("agent-prompt-builder"),
		Scripts:     []string{"https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"},
		Styles:  []string{"/web/bundle.css"},
		Version: uuid.New().String(),
	})

	if err != nil {
		log.Fatal(err)
	}
}
