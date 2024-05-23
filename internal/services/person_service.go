package services

import (
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

func (s *PersonServiceImpl) GetAll() ([]*models.Person, error) {
	return s.personRepository.GetAll()
}

func (s *PersonServiceImpl) Get(id uuid.UUID) (*models.Person, error) {
	return s.personRepository.GetByID(id)
}

func (s *PersonServiceImpl) Create(param models.CreatePerson) (*models.Person, error) {
	return s.personRepository.Save(param)
}
