package model

import "github.com/google/uuid"

type Ingredient struct {
	ID       uuid.UUID
	Name     string
	Quantity float64
	Unit     Unit
}

type Unit string

const (
	UnitGram       Unit = "ounce"
	UnitMilliliter Unit = "milliliter"
)
