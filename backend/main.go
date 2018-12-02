package main

import (
	"fmt"
)

type Ingredient struct {
	Name string
}

type Recipe struct {
	Ingredients map[string]uint8
}

func NewRecipe() Recipe {
	return Recipe{make(map[string]uint8, 0)}
}

func (r *Recipe) AddIngredient(i Ingredient, count uint8) {
	r.Ingredients[i.Name] = count
}

func main() {
	coke := Ingredient{"Coca-cola"}
	rum := Ingredient{"Rum"}
	//whisky := Ingredient{"Whisky"}

	recipe := NewRecipe()
	recipe.AddIngredient(coke, 100)
	recipe.AddIngredient(rum, 30)

	fmt.Println(recipe)
}
