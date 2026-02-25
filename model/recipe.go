package model

import "github.com/google/uuid"

type Recipe struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Ingredients  []Ingredient
	Instructions []Instruction
}
