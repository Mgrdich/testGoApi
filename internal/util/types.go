package util

import "github.com/google/uuid"

type GetByIDFunc[K any] func(id uuid.UUID) (*K, error)
