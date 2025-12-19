package models

import (
	"fmt"
)

type RecipeIngredient struct {
	Model
	RecipeID     string
	IngredientID string
	Recipe       Recipe
	Ingredient   *Ingredient
	Quantity     float64
	Unit         IngredientUnit
}

func (ri *RecipeIngredient) ToString() string {
	intQuantity := int(ri.Quantity)
	if ri.Quantity == 1 {
		return fmt.Sprintf("%d %s", intQuantity, ri.Ingredient.Name)
	} else {
		if ri.Quantity == float64(intQuantity) {
			return fmt.Sprintf("%d %s", intQuantity, ri.Ingredient.PluralName)
		} else {
			return fmt.Sprintf("%f %s", ri.Quantity, ri.Ingredient.PluralName)
		}
	}
}
