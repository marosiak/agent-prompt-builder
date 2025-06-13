package config

import (
	"github.com/google/uuid"
	"github.com/marosiak/agent-prompt-builder/ui/pages"
	"github.com/maxence-charriere/go-app/v10/pkg/app"
)

const PORT = ":8000"

func GetAppHandler(isStatic bool) *app.Handler {
	handler := &app.Handler{
		Title:       "Prompt Composer",
		Name:        "Prompt Composer",
		Description: "Will help you with building master prompts",
		Version:     uuid.New().String(),
	}

	styles := []string{"/bundle.css"}

	if isStatic {
		handler.Resources = app.GitHubPages("prompt-composer")
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
		return &pages.MainPage{}
	})
	app.Route("/import", func() app.Composer {
		return &pages.ImportPage{}
	})
}
