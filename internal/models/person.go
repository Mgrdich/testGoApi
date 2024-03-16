package models

import (
	"github.com/google/uuid"
)

type Person struct {
	ID       uuid.UUID
	Name     string
	LastName string
}

type CreatePerson struct {
	ID       uuid.UUID
	Name     string
	LastName string
}
