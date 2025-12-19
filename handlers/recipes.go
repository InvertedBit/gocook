package handlers

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	slugger "github.com/gosimple/slug"
	"github.com/invertedbit/gocook/database"
	"github.com/invertedbit/gocook/html"
	"github.com/invertedbit/gocook/html/components"
	htmlviews "github.com/invertedbit/gocook/html/views"
	"github.com/invertedbit/gocook/models"
	"github.com/invertedbit/gocook/viewmodels"
	hx "github.com/stackus/hxgo"
	"github.com/stackus/hxgo/hxfiber"
	"gorm.io/gorm"
	"maragu.dev/gomponents"
)

func HandleRecipeList(c *fiber.Ctx) error {

	ctx := context.Background()

	recipes, err := gorm.G[models.Recipe](database.DBConn).Where("entry_type = ?", models.EntryTypeActive).Find(ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to load recipes")
	}

	recipeListPage := html.Page{
		Title:           "Recipes - GoCook",
		PageContent:     htmlviews.RecipesPage(recipes),
		LayoutViewModel: GetLayoutModel(c, "Recipes", false),
	}

	handler := adaptor.HTTPHandler(recipeListPage)
	return handler(c)
}

func TryGetRecipeSlug(c *fiber.Ctx) string {
	slugString := c.Params("slug", "")
	return slugString
}

func HandleRecipeDetail(c *fiber.Ctx) error {

	slug := TryGetRecipeSlug(c)
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid recipe slug")
	}

	recipe, err := gorm.G[models.Recipe](database.DBConn).Where("slug = ?", slug).First(context.Background())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to load recipe")
	}

	layoutViewModel := GetLayoutModel(c, recipe.Name, false)

	recipeDetailPage := html.Page{
		Title:           recipe.Name + " - GoCook",
		PageContent:     htmlviews.RecipeDetailPage(&recipe),
		LayoutViewModel: layoutViewModel,
	}

	handler := adaptor.HTTPHandler(recipeDetailPage)
	return handler(c)
}

func HandleRecipeEdit(c *fiber.Ctx) error {
	slug := TryGetRecipeSlug(c)
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid recipe slug")
	}

	recipe := models.Recipe{
		Slug: slug,
		Name: "Tomato Soup",
		RecipeIngredients: []*models.RecipeIngredient{
			{
				Ingredient: &models.Ingredient{Name: "Tomato", PluralName: "Tomatos"},
				Quantity:   10,
				Unit:       models.UnitPiece,
			},
		},
		Instructions: []*models.InstructionStep{
			{Title: "Cut tomatos", Content: "Cut the tomatos into small pieces.", Order: 1},
			{Title: "Boil tomatos", Content: "Boil the tomatos in water.", Order: 2},
			{Title: "Finish with salt and pepper", Content: "Add salt and pepper to taste.", Order: 3},
		},
	}
	layoutViewModel := GetLayoutModel(c, recipe.Name, false)

	recipeEditPage := html.Page{
		Title:           recipe.Name + " - GoCook",
		PageContent:     htmlviews.RecipeEditPage(&recipe),
		LayoutViewModel: layoutViewModel,
	}

	handler := adaptor.HTTPHandler(recipeEditPage)
	return handler(c)
}

func HandleRecipeUpdate(c *fiber.Ctx) error {
	slug := TryGetRecipeSlug(c)
	if slug == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid recipe slug")
	}

	redirect := false

	recipeData := models.Recipe{}
	c.BodyParser(&recipeData)

	recipe, err := gorm.G[models.Recipe](database.DBConn).Where("slug = ?", slug).First(context.Background())

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to load recipe")
	}

	editTarget := ""

	if recipeData.Name != "" {
		editTarget = "name"
	} else if recipeData.Description != "" {
		editTarget = "description"
	}

	if recipeData.Name != "" {
	}
	switch editTarget {
	case "name":
		recipeData.Slug = slugger.Make(recipeData.Name)
		gorm.G[models.Recipe](database.DBConn).Where("ID = ?", recipe.ID).Updates(context.Background(), models.Recipe{Name: recipeData.Name, Slug: recipeData.Slug})
		hxfiber.Response(c, hx.Location(fmt.Sprintf("/recipes/%s", recipeData.Slug), hx.Target("body")))
		redirect = true
	case "description":
		gorm.G[models.Recipe](database.DBConn).Where("ID = ?", recipe.ID).Update(context.Background(), "description", recipeData.Description)
	}

	layoutViewModel := GetLayoutModel(c, recipe.Name, false)

	if redirect {
		layoutViewModel.LayoutType = viewmodels.LayoutBodyOnly
	} else {
		layoutViewModel.LayoutType = viewmodels.LayoutPartialOnly
	}

	layoutViewModel.ToastViewModel.AddToast("Recipe updated successfully", viewmodels.ToastSuccess, 10)

	var pageContent gomponents.Node

	switch editTarget {
	case "name":
		pageContent = components.EditableTextField("text-3xl font-bold", fmt.Sprintf("/recipes/%s", recipe.Slug), "name", recipeData.Name, "Name")
	case "description":
		pageContent = components.EditableTextArea("mb-4", fmt.Sprintf("/recipes/%s", recipe.Slug), "description", recipeData.Description, "Description")
	default:
		pageContent = htmlviews.RecipeDetailPage(&recipe)
	}

	recipeDetailPage := html.Page{
		Title:           recipe.Name + " - GoCook",
		PageContent:     pageContent,
		LayoutViewModel: layoutViewModel,
	}

	handler := adaptor.HTTPHandler(recipeDetailPage)
	return handler(c)
}

func HandleRecipeCreateForm(c *fiber.Ctx) error {

	hxfiber.Response(c, hx.Location("/recipes/0"))
	c.SendStatus(200)

	layoutViewModel := GetLayoutModel(c, "New Recipe", false)

	recipe := models.Recipe{}

	recipeCreatePage := html.Page{
		Title:           "New Recipe - GoCook",
		PageContent:     htmlviews.RecipeEditPage(&recipe),
		LayoutViewModel: layoutViewModel,
	}

	handler := adaptor.HTTPHandler(recipeCreatePage)
	return handler(c)
}

type RecipeCreateParams struct {
	Name string `json:"name" form:"name"`
}

func HandleRecipeCreate(c *fiber.Ctx) error {
	params := &RecipeCreateParams{}
	parseErr := c.BodyParser(params)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).SendString(parseErr.Error())
	} else if params.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Recipe name is required")
	}

	recipe := models.NewRecipe(params.Name)

	ctx := context.Background()

	err := gorm.G[models.Recipe](database.DBConn).Create(ctx, recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to create recipe")
	}

	hxfiber.Response(c, hx.Location(recipe.GetPermalink()))
	return c.SendStatus(201)
}
