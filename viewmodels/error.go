package viewmodels

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stackus/hxgo/hxfiber"
)

type NotFoundViewModel struct {
	CurrentURL string
}

func NewNotFoundViewModel(c *fiber.Ctx) *NotFoundViewModel {
	currentUrl := hxfiber.GetCurrentUrl(c)
	return &NotFoundViewModel{
		CurrentURL: currentUrl,
	}
}

func (vm *NotFoundViewModel) HasCurrentURL() bool {
	return vm.CurrentURL != ""
}
