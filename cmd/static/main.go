package main

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"log"
	"makerworld-analytics/ui/views"
)

func main() {
	// TODO fix DRY
	app.Route("/", func() app.Composer {
		return &views.MainView{}
	})

	app.RunWhenOnBrowser()

	err := app.GenerateStaticWebsite(".", &app.Handler{
		Name:        "Master prompt builder",
		Description: "Will help you with building agents",
		Resources:   app.GitHubPages("makerworld-analytics"),
		Scripts:     []string{"https://cdn.jsdelivr.net/npm/@tailwindcss/browser@4", "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"},
		Styles:      []string{"https://cdn.jsdelivr.net/npm/daisyui@5"},
		Icon: app.Icon{
			Default:  "/web/icon_x192.png",
			Large:    "/web/icon_x512.png",
			Maskable: "/web/icon_x512.png",
			SVG:      "/web/icon_x512.svg",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
