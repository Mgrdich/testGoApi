package services

import (
	"context"
	"math/big"
	"time"

	db2 ".com/internal/db"
	db ".com/internal/db/sqlc"
	".com/internal/models"
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
	dbMovie, err := s.q.GetMovie(context.Background(), db2.ToUUID(id))
	if err != nil {
		return nil, err
	}

	movie := dbMovieToMovie(dbMovie)

	return &movie, nil
}

func (s *MoviesService) Create(param models.CreateMovieParam) (*models.Movie, error) {
	dbParam := db.CreateMovieParams{
		Title:       db2.ToText(param.Title),
		Director:    db2.ToText(param.Director),
		ReleaseAt:   db2.ToDate(param.ReleaseDate),
		TicketPrice: db2.ToNumeric(big.NewInt(int64(param.TicketPrice))), // TODO research and fix this type
	}
	dbMovie, err := s.q.CreateMovie(context.Background(), dbParam)

	if err != nil {
		return nil, err
	}

	movie := dbMovieToMovie(dbMovie)

	return &movie, nil
}

func (s *MoviesService) Update(id uuid.UUID, param models.UpdateMovieParam) (*models.Movie, error) {
	dbParam := db.UpdateMovieParams{
		ID:        db2.ToUUID(id),
		Title:     db2.ToText(param.Title),
		Director:  db2.ToText(param.Director),
		ReleaseAt: db2.ToDate(param.ReleaseDate),
		// TODO research and fix this type , i think it should be string
		TicketPrice: db2.ToNumeric(big.NewInt(int64(param.TicketPrice))),
		UpdatedAt:   db2.ToTimeStamp(time.Now().UTC()),
	}
	dbMovie, err := s.q.UpdateMovie(context.Background(), dbParam)

	if err != nil {
		return nil, err
	}

	movie := dbMovieToMovie(dbMovie)

	return &movie, nil
}

func (s *MoviesService) Delete(id uuid.UUID) error {
	return s.q.DeleteMovie(context.Background(), pgtype.UUID{
		Bytes: id,
		Valid: true,
	})
}
