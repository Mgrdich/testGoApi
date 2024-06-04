package util

import (
	"fmt"

	"github.com/google/uuid"
)

type DuplicateKeyError struct {
	ID uuid.UUID
}

func (e *DuplicateKeyError) Error() string {
	return fmt.Sprintf("duplicate movie id: %v", e.ID)
}

type RecordNotFoundError struct{}

func (e *RecordNotFoundError) Error() string {
	return "record not found"
}

type ContextCouldNotBeFetchedError struct{}

func (e *ContextCouldNotBeFetchedError) Error() string {
	return "Context could not be fetched"
}

func NewContextCouldNotBeFetchedError() *ContextCouldNotBeFetchedError {
	return &ContextCouldNotBeFetchedError{}
}

type SlugIDIsNotDefined struct{}

func (e *SlugIDIsNotDefined) Error() string {
	return "Slug id is not found"
}

func NewSlugIDIsNotDefined() *SlugIDIsNotDefined {
	return &SlugIDIsNotDefined{}
}
