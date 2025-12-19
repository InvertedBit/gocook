package htmlviews

import (
	"fmt"

	"github.com/invertedbit/gocook/html/components"
	htmlpartials "github.com/invertedbit/gocook/html/partials"
	"github.com/invertedbit/gocook/models"
	"maragu.dev/gomponents"
	htmx "maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func RecipesPage(recipes []models.Recipe) gomponents.Node {
	return html.Div(
		html.Div(
			html.Class("mb-4 text-center"),
			html.Button(
				html.Class("btn btn-primary"),
				gomponents.Text("Add New Recipe"),
				// htmx.Get("/recipes/new"),
				// htmx.Target("body"),
				gomponents.Attr("onclick", "new_recipe_modal.showModal();"),
			),
		),
		html.Dialog(
			html.ID("new_recipe_modal"),
			html.Class("modal mx-auto"),
			html.Div(
				html.Class("modal-box"),
				html.H3(
					html.Class("font-bold text-lg"),
					gomponents.Text("Create New Recipe"),
				),
				html.Form(
					htmx.Post("/recipes"),
					htmx.Target("body"),
					html.Div(
						html.Input(
							html.Type("text"),
							html.Name("name"),
							html.Placeholder("Recipe Name"),
							html.Class("input input-bordered w-full mb-4"),
						),
						html.Button(
							html.Class("btn btn-primary"),
							gomponents.Text("Create Recipe"),
						),
					),
				),
			),
		),
		html.Div(
			html.Class("grid grid-cols-4 gap-5 justify-center"),
			gomponents.Map(recipes, func(recipe models.Recipe) gomponents.Node {
				return html.Div(
					// html.Class("w-1/4"),
					htmlpartials.RecipeCard(recipe),
				)
			}),
		),
	)
}

func RecipeDetailNavigation(recipe *models.Recipe) gomponents.Node {
	leftButton := html.A(
		html.Class("btn btn-secondary btn-sm"),
		gomponents.Text("Back"),
		html.Href("/recipes"),
		htmx.Boost("true"),
	)
	rightButton := html.A(
		html.Class("btn btn-primary btn-sm"),
		gomponents.Text("History"),
		htmx.Get(fmt.Sprintf("/recipes/%s/history", recipe.Slug)),
		htmx.PushURL(fmt.Sprintf("/recipes/%s/history", recipe.Slug)),
		htmx.Target("#recipe-detail-container"),
	)
	if recipe.EntryType == models.EntryTypeEdit {
		leftButton = html.A(
			html.Class("btn btn-secondary btn-sm"),
			gomponents.Text("Cancel"),
			htmx.Get(fmt.Sprintf("/recipes/%s", recipe.Slug)),
			htmx.PushURL(fmt.Sprintf("/recipes/%s", recipe.Slug)),
			htmx.Target("#recipe-detail-container"),
		)
		rightButton = html.A(
			html.Class("btn btn-primary btn-sm"),
			gomponents.Text("Save"),
			htmx.Put(fmt.Sprintf("/recipes/%s", recipe.Slug)),
			htmx.PushURL(fmt.Sprintf("/recipes/%s", recipe.Slug)),
			htmx.Target("#recipe-detail-container"),
		)
	}
	return html.Div(
		html.Class("flex flex-row w-3xl mx-auto justify-between mb-4"),
		leftButton,
		rightButton,
	)
}

func RecipeDetailPage(recipe *models.Recipe) gomponents.Node {
	return html.Div(
		html.ID("recipe-detail-container"),
		RecipeDetailNavigation(recipe),
		html.Div(
			html.Class("max-w-3xl mx-auto p-6 rounded-lg shadow-md bg-base-300 text-base-content shadow-primary"),
			components.EditableImage("w-full h-64 mb-4 bg-gray-200", fmt.Sprintf("/recipes/%s", recipe.Slug), recipe.Media, "Recipe Image"),
			html.H1(
				html.Class("text-3xl font-bold mb-4"),
				gomponents.Text(recipe.Name),
			),
			components.EditableTextField("text-3xl pb-2 font-bold", fmt.Sprintf("/recipes/%s", recipe.Slug), "name", recipe.Name, "Name"),
			components.EditableTextArea("mb-4", fmt.Sprintf("/recipes/%s", recipe.Slug), "description", recipe.Description, "Description"),
			html.H2(
				html.Class("text-2xl font-semibold mb-2"),
				gomponents.Text("Ingredients"),
			),
			html.Ul(
				html.Class("list-disc list-inside mb-4"),
				gomponents.Map(recipe.RecipeIngredients, func(ingredient *models.RecipeIngredient) gomponents.Node {
					return html.Li(
						gomponents.Text(ingredient.ToString()),
					)
				}),
			),
			html.H2(
				html.Class("text-2xl font-semibold mb-2"),
				gomponents.Text("Instructions"),
			),
			html.Ul(
				html.Class("list-decimal list-inside"),
				gomponents.Map(recipe.Instructions, func(instruction *models.InstructionStep) gomponents.Node {
					return html.Li(
						html.Class("mb-2"),
						gomponents.Text(instruction.Title+": "+instruction.Content),
					)
				}),
			),
		),
	)
}

func RecipeEditPage(recipe *models.Recipe) gomponents.Node {
	return html.Div(
		html.ID(fmt.Sprintf("recipe-container-%d", recipe.ID)),
		html.Div(
			html.Class("flex flex-row w-3xl mx-auto justify-between mb-4"),
			html.A(
				html.Class("btn btn-secondary btn-sm"),
				gomponents.Text("Discard"),
				htmx.Get(fmt.Sprintf("/recipes/%d", recipe.ID)),
				htmx.PushURL(fmt.Sprintf("/recipes/%d", recipe.ID)),
				htmx.Target(fmt.Sprintf("#recipe-container-%d", recipe.ID)),
			),
			html.A(
				html.Class("btn btn-primary btn-sm"),
				gomponents.Text("Save"),
				htmx.Put(fmt.Sprintf("/recipes/%d", recipe.ID)),
				htmx.PushURL(fmt.Sprintf("/recipes/%d", recipe.ID)),
				htmx.Target(fmt.Sprintf("#recipe-container-%d", recipe.ID)),
			),
		),
		html.Div(
			html.Class("max-w-3xl mx-auto p-6 rounded-lg shadow-md bg-base-300 text-base-content shadow-primary"),
			html.Input(
				html.Type("text"),
				html.Value(recipe.Name),
				html.Class("text-3xl font-bold mb-4"),
			),
			html.Input(
				html.Type("text"),
				html.Value(recipe.Description),
				html.Class("mb-4"),
			),
			html.H2(
				html.Class("text-2xl font-semibold mb-2"),
				gomponents.Text("Ingredients"),
			),
			html.Ul(
				html.Class("list-disc list-inside mb-4"),
				gomponents.Map(recipe.RecipeIngredients, func(ingredient *models.RecipeIngredient) gomponents.Node {
					return html.Li(
						gomponents.Text(ingredient.ToString()),
					)
				}),
			),
			html.H2(
				html.Class("text-2xl font-semibold mb-2"),
				gomponents.Text("Instructions"),
			),
			html.Ul(
				html.Class("list-decimal list-inside"),
				gomponents.Map(recipe.Instructions, func(instruction *models.InstructionStep) gomponents.Node {
					return html.Li(
						html.Class("mb-2"),
						gomponents.Text(instruction.Title+": "+instruction.Content),
					)
				}),
			),
		),
	)
}
