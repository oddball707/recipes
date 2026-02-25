package model

import (
	"encoding/json"
)

type Ingredient struct {
	Name     string
	Quantity float64
	Unit     *Unit
}

type Unit struct {
	Name         string
	Abbreviation string
}

func (u *Unit) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		// If the data is not a string, try to unmarshal it as a struct
		type Alias Unit
		var a Alias
		if err2 := json.Unmarshal(data, &a); err2 != nil {
			return err // Return the original error if both fail
		}
		*u = Unit(a)
		return nil
	}
	u.Abbreviation = s
	return nil
}
