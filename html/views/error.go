package htmlviews

import (
	"github.com/invertedbit/gocook/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func NotFoundPage(vm *viewmodels.NotFoundViewModel) gomponents.Node {
	return html.Div(
		html.Class("text-center"),
		html.H1(
			html.Class("text-5xl font-bold"),
			gomponents.Text("404"),
		),
		html.P(
			html.Class("text-xl mt-4"),
			gomponents.Text("Page Not Found"),
		),
		gomponents.If(vm.HasCurrentURL(),
			html.A(
				html.Class("btn btn-primary mt-6"),
				html.Href(vm.CurrentURL),
				htmx.Boost("true"),
				gomponents.Text("Go back"),
			),
		),
		gomponents.If(!vm.HasCurrentURL(),
			html.A(
				html.Class("btn btn-primary mt-6"),
				html.Href("/"),
				htmx.Boost("true"),
				gomponents.Text("Go to Home"),
			),
		),
	)
}
