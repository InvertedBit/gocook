package htmlviews

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func HomePage() gomponents.Node {
	return html.Div(
		html.H1(
			gomponents.Text("Welcome to GoCook!"),
		),
	)
}
