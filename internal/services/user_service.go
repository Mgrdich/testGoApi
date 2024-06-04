package services

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"testGoApi/internal/models"
	"testGoApi/internal/repository"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	tokenService   TokenService
}

func NewUserServiceImpl(userRepository repository.UserRepository, tokenService TokenService) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (uS *UserServiceImpl) Login(ctx context.Context, param models.LoginUser) (string, error) {
	user, err := uS.userRepository.GetByUsername(ctx, param.Username)

	if err != nil {
		return "", err
	}

	bytePasswordHashed := []byte(user.Password)
	err = bcrypt.CompareHashAndPassword(bytePasswordHashed, []byte(param.Password))

	if err != nil {
		return "", err
	}

	return uS.tokenService.GenerateJWT(&models.TokenizedUser{
		ID:       user.ID,
		Username: user.Username,
		Role:     models.LookUpRoleString(user.Role),
	})
}

func (uS *UserServiceImpl) Create(ctx context.Context, param models.CreateUser) (*models.User, error) {
	password := param.Password

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	param.Password = string(hash)

	return uS.userRepository.Save(ctx, param)
}

func (uS *UserServiceImpl) Get(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return uS.userRepository.GetByID(ctx, id)
}
