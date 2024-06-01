package test_helpers

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MockMovieService struct{}

func (*MockMovieService) GetAll(_ context.Context) ([]*models.Movie, error) {
	return nil, nil
}
func (*MockMovieService) Get(_ context.Context, id uuid.UUID) (*models.Movie, error) {
	return &models.Movie{ID: id}, nil
}
func (*MockMovieService) Create(_ context.Context, param models.CreateMovieParam) (*models.Movie, error) {
	return &models.Movie{
		Title:       param.Title,
		Director:    param.Director,
		TicketPrice: param.TicketPrice,
	}, nil
}
func (*MockMovieService) Delete(_ context.Context, _ uuid.UUID) error {
	return nil
}
func (*MockMovieService) Update(_ context.Context, id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return &models.Movie{
		ID:          id,
		Title:       param.Title,
		Director:    param.Director,
		TicketPrice: param.TicketPrice,
	}, nil
}
