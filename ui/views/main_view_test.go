package views

import (
	"github.com/maxence-charriere/go-app/v10/pkg/app"
	"testing"
)

func TestComponentPreRendering(t *testing.T) {
	compo := &MainView{}

	engine := app.NewTestEngine()
	err := engine.Load(compo)
	if err != nil {
		return
	}


}