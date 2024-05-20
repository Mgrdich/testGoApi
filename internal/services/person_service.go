package services

import (
	"context"

	"github.com/google/uuid"
	db2 "testGoApi/internal/db"
	db "testGoApi/internal/db/sqlc"
	"testGoApi/internal/models"
)

type PersonServiceImpl struct {
	q *db.Queries
}

func NewPersonServiceImpl(queries *db.Queries) *PersonServiceImpl {
	return &PersonServiceImpl{
		q: queries,
	}
}

func dbPersonToPerson(person db.Person) *models.Person {
	return &models.Person{
		ID:        person.ID.Bytes,
		FirstName: person.FirstName.String,
		LastName:  person.LastName.String,
		CreatedAt: person.CreatedAt.Time,
	}
}

func (s *PersonServiceImpl) GetAll() ([]*models.Person, error) {
	dbPeople, err := s.q.GetAllPerson(context.Background())

	if err != nil {
		return nil, err
	}

	var people []*models.Person

	for _, person := range dbPeople {
		people = append(people, dbPersonToPerson(person))
	}

	return people, nil
}

func (s *PersonServiceImpl) GetByID(id uuid.UUID) (*models.Person, error) {
	dbPerson, err := s.q.GetPerson(context.Background(), db2.ToUUID(id))
	if err != nil {
		return nil, err
	}

	person := dbPersonToPerson(dbPerson)

	return person, nil
}

func (s *PersonServiceImpl) Create(param models.CreatePerson) (*models.Person, error) {
	dbParam := db.CreatePersonParams{
		FirstName: db2.ToText(param.FirstName),
		LastName:  db2.ToText(param.LastName),
	}
	dbPerson, err := s.q.CreatePerson(context.Background(), dbParam)

	if err != nil {
		return nil, err
	}

	person := dbPersonToPerson(dbPerson)

	return person, nil
}
