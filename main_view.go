func (m *MainView) renderScrollToBottomButton() app.HTMLButton {
	return app.Button().
		Class("btn btn-circle scroll-to-bottom-btn").
		OnClick(func(ctx app.Context, e app.Event) {
			// Try to scroll to Output section first
			if outputElem := app.Window().GetElementByID("output-section"); outputElem != nil {
				outputElem.Call("scrollIntoView", map[string]interface{}{
					"behavior": "smooth",
					"block":    "start",
				})
			} else {
				// Fallback: scroll to bottom of page
				app.Window().GetDocument().Get("documentElement").Set("scrollTop",
					app.Window().GetDocument().Get("documentElement").Get("scrollHeight"))
			}
		}).
		Body(
			&SVGIcon{
				IconData:       TrashIcon,
				Color:          "white",
				OpacityPercent: 90,
			},
		)
}
