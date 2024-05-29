package services

import (
	"context"
	"github.com/google/uuid"
	"testGoApi/internal/models"
	"testGoApi/internal/repository"
)

type PersonServiceImpl struct {
	personRepository repository.PersonRepository
}

func NewPersonServiceImpl(personRepository repository.PersonRepository) *PersonServiceImpl {
	return &PersonServiceImpl{
		personRepository: personRepository,
	}
}

func (s *PersonServiceImpl) GetAll(ctx context.Context) ([]*models.Person, error) {
	return s.personRepository.GetAll(ctx)
}

func (s *PersonServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	return s.personRepository.GetByID(ctx, id)
}

func (s *PersonServiceImpl) Create(ctx context.Context, param models.CreatePerson) (*models.Person, error) {
	return s.personRepository.Save(ctx, param)
}
