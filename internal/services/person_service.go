package services

import (
	"context"

	db2 ".com/internal/db"
	db ".com/internal/db/sqlc"
	".com/internal/models"
	"github.com/google/uuid"
)

type PersonService struct {
	q *db.Queries
}

func NewPersonService() *PersonService {
	return &PersonService{
		q: db2.GetPQueries(),
	}
}

func dbPersonToPerson(person db.Person) models.Person {
	return models.Person{
		ID:        person.ID.Bytes,
		FirstName: person.FirstName.String,
		LastName:  person.LastName.String,
		CreatedAt: person.CreatedAt.Time,
	}
}

func (p *PersonService) GetAll() ([]models.Person, error) {
	dbPeople, err := p.q.GetAllPerson(context.Background())

	if err != nil {
		return nil, err
	}

	var people []models.Person

	for _, person := range dbPeople {
		people = append(people, dbPersonToPerson(person))
	}

	return people, nil
}

func (s *PersonService) GetByID(id uuid.UUID) (*models.Person, error) {
	dbPerson, err := s.q.GetPerson(context.Background(), db2.ToUUID(id))
	if err != nil {
		return nil, err
	}

	person := dbPersonToPerson(dbPerson)

	return &person, nil
}

func (s *PersonService) Create(param models.CreatePerson) (*models.Person, error) {
	dbParam := db.CreatePersonParams{
		FirstName: db2.ToText(param.FirstName),
		LastName:  db2.ToText(param.LastName),
	}
	dbPerson, err := s.q.CreatePerson(context.Background(), dbParam)

	if err != nil {
		return nil, err
	}

	person := dbPersonToPerson(dbPerson)

	return &person, nil
}
