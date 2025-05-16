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
		Version:     uuid.New().String(),
	}

	styles := []string{"/bundle.css"}

	if isStatic {
		handler.Resources = app.GitHubPages("agent-prompt-builder")
	} else {
		var tmpStyles []string
		for _, style := range styles {
			tmpStyles = append(tmpStyles, "/web"+style)
		}

		styles = tmpStyles
	}
	handler.Styles = styles
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
