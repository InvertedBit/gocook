package html

import (
	"errors"
	"fmt"
	"io"

	"github.com/invertedbit/gocook/viewmodels"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

type Layout struct {
	LayoutViewModel *viewmodels.LayoutViewModel
}

func (l *Layout) IsOOB() bool {
	layoutType := l.LayoutViewModel.LayoutType
	return layoutType == viewmodels.LayoutPartialOnly || layoutType == viewmodels.LayoutPartialWithFragments
}

func (l *Layout) GetHeader() gomponents.Node {

	return html.Header(
		html.ID("header"),
		gomponents.If(l.IsOOB(), htmx.SwapOOB("true")),
		html.Div(
			htmx.Boost("true"),
			html.Class("navbar bg-base-300 shadow-sm"),
			html.Div(
				html.Class("navbar-start"),
				html.Div(
					html.Class("dropdown"),
					html.Div(
						html.TabIndex("0"),
						html.Role("button"),
						html.Class("btn btn-ghost lg:hidden"),
						html.H1(gomponents.Text("GoCook")),
					),
					html.Ul(
						html.TabIndex("0"),
						html.Class("menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"),
						gomponents.Map(l.LayoutViewModel.Navbar.Items, func(item *viewmodels.NavbarMenuItem) gomponents.Node {
							return html.Li(
								html.A(
									html.Href(item.Link),
									gomponents.Text(item.Label),
								),
							)
						}),
					),
				),
				html.A(
					html.Class("btn btn-ghost text-xl"),
					html.Href("/"),
					gomponents.Text("GoCook"),
				),
			),
			html.Div(
				html.Class("navbar-center hidden lg:flex"),
				html.Ul(
					html.Class("menu menu-horizontal px-1"),
					gomponents.Map(l.LayoutViewModel.Navbar.Items, func(item *viewmodels.NavbarMenuItem) gomponents.Node {
						return html.Li(
							gomponents.If(len(item.Children) > 0,
								html.Details(
									html.Summary(
										html.Class("text-nowrap"),
										gomponents.Text(item.Label),
									),
									html.Ul(
										html.Class("z-10 p-2"),
										gomponents.Map(item.Children, func(child *viewmodels.NavbarMenuItem) gomponents.Node {
											return html.Li(
												html.A(
													html.Href(child.Link),
													html.Class("text-nowrap"),
													gomponents.Text(child.Label),
												),
											)
										}),
									),
								),
							),
							gomponents.If(len(item.Children) == 0,
								html.A(
									html.Class("text-nowrap"),
									html.Href(item.Link),
									gomponents.Text(item.Label),
								),
							),
						)
					}),
				),
			),
			html.Div(
				html.Class("navbar-end"),
				html.A(
					html.Class("btn"),
					gomponents.Text("Button"),
				),
			),
		),
	)
}

func (l *Layout) GetBody(pageContent gomponents.Node) []gomponents.Node {
	return []gomponents.Node{
		html.Div(
			l.GetToastContainer(),
			html.Class("min-h-screen flex flex-col"),
			l.GetHeader(),
			html.Main(
				html.Class("container mx-auto p-4 flex-grow"),
				pageContent,
			),
			html.Footer(
				html.Class("footer p-4 bg-neutral text-neutral-content footer-center"),
				html.Div(
					html.P(gomponents.Raw("Copyright &copy; 2025 - All rights reserved by GoCook")),
				),
			),
		),
	}
}

func (l *Layout) Render(w io.Writer, pageContent gomponents.Node) error {
	switch l.LayoutViewModel.LayoutType {
	case viewmodels.LayoutFull:
		return l.RenderFullLayout(w, pageContent)
	case viewmodels.LayoutBodyOnly:
		return l.RenderBodyOnly(w, pageContent)
	case viewmodels.LayoutPartialOnly:
		return l.RenderPartialOnly(w, pageContent)
	case viewmodels.LayoutPartialWithFragments:
		return l.RenderPartialWithFragments(w, pageContent)
	case viewmodels.LayoutError:
		return l.RenderErrorLayout(w, pageContent)
	default:
		return l.RenderFullLayout(w, pageContent)
	}
}

func (l *Layout) RenderBodyOnly(w io.Writer, pageContent gomponents.Node) error {
	err :=
		html.TitleEl(
			htmx.SwapOOB("true"),
			gomponents.Text(l.LayoutViewModel.Page),
		).Render(w)
	if err != nil {
		return err
	}
	return html.Body(l.GetBody(pageContent)...).Render(w)
}

func (l *Layout) RenderPartialOnly(w io.Writer, pageContent gomponents.Node) error {
	err := html.TitleEl(
		gomponents.Text(l.LayoutViewModel.Page),
	).Render(w)
	if err != nil {
		return err
	}
	err = l.GetHeader().Render(w)
	if err != nil {
		return err
	}
	err = l.GetToastContainer().Render(w)
	if err != nil {
		return err
	}
	return pageContent.Render(w)
}

func (l *Layout) RenderPartialWithFragments(w io.Writer, pageContent gomponents.Node) error {
	return errors.New("not implemented")
}

func (l *Layout) RenderErrorLayout(w io.Writer, pageContent gomponents.Node) error {
	return errors.New("not implemented")
}

func (l *Layout) GetToastContainer() gomponents.Node {
	return html.Div(
		html.ID("toast-container"),
		gomponents.If(l.IsOOB(), htmx.SwapOOB("true")),
		html.Class("toast toast-top toast-end"),
		gomponents.Map(l.LayoutViewModel.ToastViewModel.Messages, func(toast viewmodels.ToastMessage) gomponents.Node {
			return html.Div(
				html.Class(fmt.Sprintf("alert %s", toast.GetClassName())),
				html.Span(
					gomponents.Text(toast.Message),
				),
			)
		}),
	)
}

func (l *Layout) RenderFullLayout(w io.Writer, pageContent gomponents.Node) error {
	return components.HTML5(components.HTML5Props{
		Title: l.LayoutViewModel.Page,
		Head: []gomponents.Node{
			html.Meta(html.Charset("UTF-8")),
			html.Meta(
				html.Name("viewport"),
				html.Content("width=device-width, initial-scale=1"),
			),
			html.Link(
				html.Rel("stylesheet"),
				html.Href("/css/theme.css"),
			),
			html.Link(
				html.Rel("stylesheet"),
				html.Href("https://cdn.jsdelivr.net/npm/remixicon@4.5.0/fonts/remixicon.css"),
			),
			html.Script(
				html.Src("https://cdn.jsdelivr.net/npm/htmx.org@2.0.6/dist/htmx.min.js"),
			),
			html.Script(
				html.Src("https://unpkg.com/hyperscript.org@0.9.14"),
			),
		},
		Body: l.GetBody(pageContent),
	}).Render(w)
}
