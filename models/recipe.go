package models

import (
	"fmt"

	"github.com/gosimple/slug"
)

type InstructionStep struct {
	Model
	RecipeID string
	Title    string
	Content  string
	Order    int
}

type Recipe struct {
	CoWModel
	Name              string
	Slug              string `gorm:"uniqueIndex"`
	Description       string
	RecipeIngredients []*RecipeIngredient
	Instructions      []*InstructionStep `gorm:"foreignKey:RecipeID"`
	Previous          *Recipe            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (r *Recipe) GetPermalink() string {
	return fmt.Sprintf("/recipes/%s", r.Slug)
}

func NewRecipe(name string) *Recipe {
	slug := slug.Make(name)

	recipe := &Recipe{
		Name: name,
		Slug: slug,
	}

	return recipe
}
