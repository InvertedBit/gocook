package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/invertedbit/gocook/html"
	htmlviews "github.com/invertedbit/gocook/html/views"
	"github.com/invertedbit/gocook/viewmodels"
)

func GetNavbarModel(c *fiber.Ctx) *viewmodels.NavbarViewModel {
	navbarModel := viewmodels.NewNavbarViewModel()

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label: "Dashboard",
		Link:  "/",
	})

	myRecipes := viewmodels.NavbarMenuItem{
		Label: "My Recipes",
		Link:  "/recipes",
	}

	categories := viewmodels.NavbarMenuItem{
		Label: "Categories",
		Link:  "/recipe-categories",
	}

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label:    "Recipes",
		Link:     "#",
		Children: []*viewmodels.NavbarMenuItem{&myRecipes, &categories},
	})

	navbarModel.AddItem(&viewmodels.NavbarMenuItem{
		Label: "Ingredients",
		Link:  "/ingredients",
	})

	return navbarModel
}

func GetLayoutModel(c *fiber.Ctx, title string, isAuthenticated bool) *viewmodels.LayoutViewModel {
	return viewmodels.NewLayoutViewModel(title, GetNavbarModel(c), false, 2025, c)
}

func New() *fiber.App {
	app := fiber.New()

	app.Static("/", "./assets")

	app.Get("/", HandleViewHome)

	app.Get("/recipes", HandleRecipeList)

	app.Get("/recipes/new", HandleRecipeCreateForm)

	app.Get("/recipes/:slug", HandleRecipeDetail)

	app.Put("/recipes/:slug", HandleRecipeUpdate)

	app.Get("/recipes/:slug/edit", HandleRecipeEdit)

	app.Post("/recipes", HandleRecipeCreate)

	app.Get("*", HandleNotFound)

	return app
}

func HandleNotFound(c *fiber.Ctx) error {
	// notFound := views.NotFoundPage(GetLayoutModel(c, "Not Found", false))
	notFoundPage := html.Page{
		Title:           "404 Not Found - GoCook",
		PageContent:     htmlviews.NotFoundPage(viewmodels.NewNotFoundViewModel(c)),
		LayoutViewModel: GetLayoutModel(c, "Not Found", false),
	}

	handler := adaptor.HTTPHandler(notFoundPage)
	return handler(c)
}

func HandleViewHome(c *fiber.Ctx) error {
	// home := views.HomePage(GetLayoutModel(c, "Dashboard", false), "Hello World!")

	homePage := html.Page{
		Title:           "Home - GoCook",
		PageContent:     htmlviews.HomePage(),
		LayoutViewModel: GetLayoutModel(c, "Dashboard", false),
	}

	handler := adaptor.HTTPHandler(homePage)

	return handler(c)
}
