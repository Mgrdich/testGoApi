package util

import (
	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type GetByIDFunc[K models.Models] func(id uuid.UUID) (*K, error)
