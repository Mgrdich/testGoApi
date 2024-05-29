package services

import (
	"context"

	"testGoApi/internal/models"
	"testGoApi/internal/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (uS *UserServiceImpl) Get(ctx context.Context, username string) (*models.User, error) {
	return uS.userRepository.GetByUsername(ctx, username)
}

func (uS *UserServiceImpl) Create(ctx context.Context, param models.CreateUser) (*models.User, error) {
	return uS.userRepository.Save(ctx, param)
}
