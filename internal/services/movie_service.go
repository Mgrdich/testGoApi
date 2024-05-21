package services

import (
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

func (s *MoviesServiceImpl) GetAll() ([]*models.Movie, error) {
	return s.movieRepository.GetAll()
}

func (s *MoviesServiceImpl) Get(id uuid.UUID) (*models.Movie, error) {
	return s.movieRepository.GetByID(id)
}

func (s *MoviesServiceImpl) Create(param models.CreateMovieParam) (*models.Movie, error) {
	return s.movieRepository.Save(param)
}

func (s *MoviesServiceImpl) Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return s.movieRepository.UpdateByID(id, param)
}

func (s *MoviesServiceImpl) Delete(id uuid.UUID) error {
	return s.movieRepository.DeleteByID(id)
}
