package util

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type GetByIDFunc[K models.Models] func(ctx context.Context, id uuid.UUID) (*K, error)
