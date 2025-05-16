package config

import (
	"github.com/google/uuid"
	"github.com/marosiak/agent-prompt-builder/ui/views"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const PORT = ":8000"

func GetAppHandler(isStatic bool) *app.Handler {
	handler := &app.Handler{
		Title:       "Master Prompt Builder",
		Name:        "Master prompt builder",
		Description: "Will help you with building agents",
		Styles:      []string{"/web/bundle.css"},
		Version:     uuid.New().String(),
	}
	if isStatic {
		handler.Resources = app.GitHubPages("agent-prompt-builder")
	}
	return handler
}

func RegisterRoutes() {
	app.Route("/", func() app.Composer {
		return &views.MainView{}
	})
	app.Route("/import", func() app.Composer {
		return &views.ImportView{}
	})
}
