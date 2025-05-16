package main

import (
	"github.com/marosiak/agent-prompt-builder/config"
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

	err := app.GenerateStaticWebsite(".", config.GetAppHandler(true))

	if err != nil {
		log.Fatal(err)
	}
}
