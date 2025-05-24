package main

import (
	"github.com/marosiak/agent-prompt-builder/ui/pages"
	"log"
	"net/http"

	"github.com/marosiak/agent-prompt-builder/config"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {

	app.Route("/", func() app.Composer {
		return &pages.MainPage{}
	})
	app.Route("/import", func() app.Composer {
		return &pages.ImportPage{}
	})

	config.RegisterRoutes()

	app.RunWhenOnBrowser()
	http.Handle("/", config.GetAppHandler(false))

	if err := http.ListenAndServe(config.PORT, nil); err != nil {
		log.Fatal(err)
	}
}
