package services

import (
	db2 ".com/internal/db"
	".com/internal/db/sqlc"
	".com/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type MoviesService struct {
	q *db.Queries
}

func NewMoviesService() *MoviesService {
	return &MoviesService{
		q: db2.GetPQueries(),
	}
}

func dbMovieToMovie(movie db.Movie) models.Movie {
	ticketPrice, _ := movie.TicketPrice.Float64Value()

	return models.Movie{
		ID:          movie.ID.Bytes,
		Title:       movie.Title.String,
		Director:    movie.Director.String,
		ReleaseDate: movie.ReleaseAt.Time,
		TicketPrice: ticketPrice.Float64,
		CreatedAt:   movie.CreatedAt.Time,
		UpdatedAt:   movie.UpdatedAt.Time,
	}
}

func (s *MoviesService) GetAll() ([]models.Movie, error) {
	dbMovies, err := s.q.GetAllMovies(context.Background())
	if err != nil {
		return nil, err
	}

	var movies []models.Movie
	for _, mm := range dbMovies {
		movies = append(movies, dbMovieToMovie(mm))
	}

	return movies, nil
}

func (s *MoviesService) GetByID(id uuid.UUID) (*models.Movie, error) {
	dbMovie, err := s.q.GetMovie(context.Background(), pgtype.UUID{Bytes: id})
	if err != nil {
		return nil, err
	}

	movie := dbMovieToMovie(dbMovie)

	return &movie, nil
}

func (s *MoviesService) Create(param models.CreateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func (s *MoviesService) Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	return nil, nil
}

func (s *MoviesService) Delete(id uuid.UUID) error {
	return nil
}
