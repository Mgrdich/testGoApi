package services

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
	"testGoApi/internal/repository"
)

type MoviesServiceImpl struct {
	movieRepository repository.MovieRepository
}

func NewMoviesServiceImpl(movieRepository repository.MovieRepository) *MoviesServiceImpl {
	return &MoviesServiceImpl{
		movieRepository: movieRepository,
	}
}

func (s *MoviesServiceImpl) GetAll(ctx context.Context) ([]*models.Movie, error) {
	return s.movieRepository.GetAll(ctx)
}

func (s *MoviesServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.Movie, error) {
	return s.movieRepository.GetByID(ctx, id)
}

func (s *MoviesServiceImpl) Create(ctx context.Context, param models.CreateMovie) (*models.Movie, error) {
	return s.movieRepository.Save(ctx, param)
}

func (s *MoviesServiceImpl) Update(
	ctx context.Context,
	id uuid.UUID,
	param models.UpdateMovie) (*models.Movie, error) {
	return s.movieRepository.UpdateByID(ctx, id, param)
}

func (s *MoviesServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.movieRepository.DeleteByID(ctx, id)
}
