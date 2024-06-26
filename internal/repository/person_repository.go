package repository

import (
	"context"

	"github.com/google/uuid"
	db2 "testGoApi/internal/db"
	db "testGoApi/internal/db/sqlc"
	"testGoApi/internal/models"
)

type PersonRepositoryImpl struct {
	q *db.Queries
}

func NewPersonRepositoryImpl(queries *db.Queries) *PersonRepositoryImpl {
	return &PersonRepositoryImpl{
		q: queries,
	}
}

func mapDBPersonToModelPerson(person db.Person) *models.Person {
	return &models.Person{
		ID:        person.ID.Bytes,
		FirstName: person.FirstName.String,
		LastName:  person.LastName.String,
		CreatedAt: person.CreatedAt.Time,
	}
}

func (s *PersonRepositoryImpl) GetAll(ctx context.Context) ([]*models.Person, error) {
	dbPeople, err := s.q.GetAllPerson(ctx)

	if err != nil {
		return nil, err
	}

	var people []*models.Person

	for _, person := range dbPeople {
		people = append(people, mapDBPersonToModelPerson(person))
	}

	return people, nil
}

func (s *PersonRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*models.Person, error) {
	dbPerson, err := s.q.GetPerson(ctx, db2.ToUUID(id))
	if err != nil {
		return nil, err
	}

	person := mapDBPersonToModelPerson(dbPerson)

	return person, nil
}

func (s *PersonRepositoryImpl) Save(ctx context.Context, param models.CreatePerson) (*models.Person, error) {
	dbParam := db.CreatePersonParams{
		FirstName: db2.ToText(param.FirstName),
		LastName:  db2.ToText(param.LastName),
	}
	dbPerson, err := s.q.CreatePerson(ctx, dbParam)

	if err != nil {
		return nil, err
	}

	person := mapDBPersonToModelPerson(dbPerson)

	return person, nil
}
