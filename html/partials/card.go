package htmlpartials

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func BaseCard(title string, text string, footer gomponents.Node) gomponents.Node {
	return html.Div(
		html.Class("card bg-base-300 shadow-md shadow-secondary hover:shadow-lg hover:shadow-primary transition-shadow h-full"),
		html.Div(
			html.Class("card-body"),
			html.H5(
				html.Class("card-title"),
				gomponents.Text(title),
			),
			html.P(
				html.Class("card-text"),
				gomponents.Text(text),
			),
			footer,
		),
	)
}
