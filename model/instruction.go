package model

import "github.com/google/uuid"

type Instruction struct {
	ID         uuid.UUID
	StepNumber int
	Text       string
}
