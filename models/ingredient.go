package models

type IngredientUnit int

const (
	UnitPiece IngredientUnit = iota
	UnitMass
	UnitVolume
	UnitLength
)

type Ingredient struct {
	Model
	Name       string
	PluralName string
	Units      IngredientUnit
}
