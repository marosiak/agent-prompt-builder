package main

import (
	"github.com/google/uuid"
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

	app.RunWhenOnBrowser()
	http.Handle("/", &app.Handler{
		Name:        "Master prompt builder",
		Description: "Will help you with building agents",
		Scripts:     []string{"https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4"},
		Styles:  []string{"/web/bundle.css"},
		Version: uuid.New().String(),
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
