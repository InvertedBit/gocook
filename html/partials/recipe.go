package htmlpartials

import (
	"github.com/invertedbit/gocook/models"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func RecipeCard(recipe models.Recipe) gomponents.Node {
	return BaseCard(recipe.Name, recipe.Description, html.Div(
		html.Class("mt-6"),
		html.A(
			html.Class("btn btn-primary"),
			html.Href(recipe.GetPermalink()),
			htmx.Boost("true"),
			gomponents.Text("View Recipe"),
		),
	))
	return html.Div(
		html.Class("card bg-base-300 shadow-md shadow-secondary hover:shadow-lg hover:shadow-primary transition-shadow h-full"),
		html.Div(
			html.Class("card-body"),
			html.H5(
				html.Class("card-title"),
				gomponents.Text(recipe.Name),
			),
			html.P(
				html.Class("card-text"),
				gomponents.Text(recipe.Description),
			),
			html.Div(
				html.Class("mt-6"),
				html.A(
					html.Class("btn btn-primary"),
					html.Href(recipe.GetPermalink()),
					htmx.Boost("true"),
					gomponents.Text("View Recipe"),
				),
			),
		),
	)
}

func NewRecipeCard() gomponents.Node {
	return BaseCard("New Recipe", "Create a new recipe to get started.", html.Div(
		html.Class("mt-6"),
		html.A(
			html.Class("btn btn-primary"),
			html.Href("/recipes/new"),
			htmx.Boost("true"),
			gomponents.Text("Create Recipe"),
		),
	))
}
