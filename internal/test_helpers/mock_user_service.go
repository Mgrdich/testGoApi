package test_helpers

import (
	"context"

	"github.com/google/uuid"
	"testGoApi/internal/models"
)

const JWTMockString = "JWTMockString"

type MockUserService struct{}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (m *MockUserService) Login(_ context.Context, _ models.LoginUser) (string, error) {
	return JWTMockString, nil
}

func (m *MockUserService) Create(_ context.Context, param models.CreateUser) (*models.User, error) {
	return &models.User{
		ID:       uuid.New(),
		Username: param.Username,
		Password: param.Password,
		Role:     param.Role,
	}, nil
}
