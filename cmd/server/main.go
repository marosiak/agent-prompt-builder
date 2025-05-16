package main

import (
	"github.com/marosiak/agent-prompt-builder/config"
	"github.com/marosiak/agent-prompt-builder/ui/views"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

// The main function is the entry point where the app is configured and started.
// It is executed in 2 different environments: A client (the web browser) and a
// server.
func main() {

	app.Route("/", func() app.Composer {
		return &views.MainView{}
	})
	app.Route("/import", func() app.Composer {
		return &views.ImportView{}
	})

	config.RegisterRoutes()

	app.RunWhenOnBrowser()
	http.Handle("/", config.GetAppHandler(true))

	if err := http.ListenAndServe(config.PORT, nil); err != nil {
		log.Fatal(err)
	}
}
