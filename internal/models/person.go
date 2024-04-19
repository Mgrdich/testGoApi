package models

import (
	"time"

	"github.com/google/uuid"
)

type Person struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	CreatedAt time.Time
}

type CreatePerson struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
}
