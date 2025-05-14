package main

import (
	"log"
	"makerworld-analytics/ui/views"
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
		Scripts:     []string{"https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4", "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"},
		Styles:      []string{"https://cdn.jsdelivr.net/npm/daisyui@5"},

		Icon: app.Icon{
			Default:  "/web/icon_x192.png",
			Large:    "/web/icon_x512.png",
			Maskable: "/web/icon_x512.png",
			SVG:      "/web/icon_x512.svg",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
