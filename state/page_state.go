package state

import "github.com/maxence-charriere/go-app/v10/pkg/app"

func PageStateKey() string {
	return "page-state"
}

func SetCurrentPageIndex(ctx app.Context, currentIndex int) {
	ctx.SetState(PageStateKey(), currentIndex).Persist()
}

func GetCurrentPageIndex(ctx app.Context) int {
	var currentIndex int
	ctx.GetState(PageStateKey(), &currentIndex)
	return currentIndex
}
