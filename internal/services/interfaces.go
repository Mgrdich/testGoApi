package services

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

type MovieService interface {
	GetAll(ctx context.Context) ([]*models.Movie, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Movie, error)
	Create(ctx context.Context, param models.CreateMovie) (*models.Movie, error)
	Update(ctx context.Context, id uuid.UUID, param models.UpdateMovie) (*models.Movie, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type PersonService interface {
	GetAll(ctx context.Context) ([]*models.Person, error)
	Get(ctx context.Context, id uuid.UUID) (*models.Person, error)
	Create(ctx context.Context, param models.CreatePerson) (*models.Person, error)
}

type TokenService interface {
	GenerateJWT(user *models.TokenizedUser) (string, error)
	VerifyJWT(tokenString string) error
	ParseJWT(tokenString string) (*models.TokenizedUser, error)
}

type UserService interface {
	Login(ctx context.Context, parser models.LoginUser) (string, error)
	Create(ctx context.Context, param models.CreateUser) (*models.User, error)
}
