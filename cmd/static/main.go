package main

import (
	"github.com/marosiak/agent-prompt-builder/ui/pages"
	"log"

	"github.com/marosiak/agent-prompt-builder/config"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

func main() {
	// TODO fix DRY
	app.Route("/", func() app.Composer {
		return &pages.MainPage{}
	})
	app.Route("/import", func() app.Composer {
		return &pages.ImportPage{}
	})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", config.GetAppHandler(true))

	if err != nil {
		log.Fatal(err)
	}
}
