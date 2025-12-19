package viewmodels

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/stackus/hxgo/hxfiber"
)

type LayoutType int

const (
	LayoutFull LayoutType = iota
	LayoutBodyOnly
	LayoutPartialOnly
	LayoutPartialWithFragments
	LayoutError
)

type LayoutViewModel struct {
	Page           string
	Navbar         *NavbarViewModel
	IsError        bool
	CopyrightYear  int
	LayoutType     LayoutType
	ToastViewModel *ToastViewModel
}

func NewLayoutViewModel(page string, navbar *NavbarViewModel, isError bool, copyrightYear int, c *fiber.Ctx) *LayoutViewModel {
	layoutType := LayoutFull
	hxTarget := hxfiber.GetTarget(c)
	if hxfiber.IsBoosted(c) || hxfiber.IsHtmx(c) {
		if hxTarget == "body" || hxTarget == "" {
			layoutType = LayoutBodyOnly
			fmt.Println("Boosted request detected")
		} else {
			layoutType = LayoutPartialOnly
			fmt.Println("HTMX partial request detected")
		}
	}
	return &LayoutViewModel{
		Page:           page,
		Navbar:         navbar,
		IsError:        isError,
		CopyrightYear:  copyrightYear,
		LayoutType:     layoutType,
		ToastViewModel: NewToastViewModel(),
	}
}
