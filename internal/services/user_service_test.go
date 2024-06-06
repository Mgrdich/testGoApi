package services

import (
	"context"
	"errors"
	"testGoApi/configs"
	"testGoApi/internal/util"
	"testing"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"testGoApi/internal/models"
)

const username = "username"
const password = "password"
const basicRole = models.BasicRole

type MockUserRepository struct {
	GetByUsernameFunc func(username string) (*models.User, error)
	SaveFunc          func(param models.CreateUser) (*models.User, error)
	GetByIDFunc       func(id uuid.UUID) (*models.User, error)
}

func (m *MockUserRepository) GetByUsername(_ context.Context, username string) (*models.User, error) {
	return m.GetByUsernameFunc(username)
}

func (m *MockUserRepository) Save(_ context.Context, param models.CreateUser) (*models.User, error) {
	return m.SaveFunc(param)
}

func (m *MockUserRepository) GetByID(_ context.Context, id uuid.UUID) (*models.User, error) {
	return m.GetByIDFunc(id)
}

func TestNewUserService_Login(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment:            util.DevEnvironment,
		TokenExpirationMinutes: 10,
		JwtSecretKey:           []byte("DummyJWTSecretKey"),
	})

	tokenService := NewTokenServiceImpl()

	hash, err := GenerateHashPassword(password)

	if err != nil {
		t.Fatalf("Generte hash expected no error, got %v", err)
	}

	userService := NewUserServiceImpl(&MockUserRepository{
		GetByUsernameFunc: func(username string) (*models.User, error) {
			return &models.User{
				ID:       uuid.New(),
				Username: username,
				Password: hash,
				Role:     models.BasicRole,
			}, nil
		},
	}, tokenService)

	token, err := userService.Login(context.Background(), models.LoginUser{Username: username, Password: password})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(token) == 0 {
		t.Fatalf("token should not be empty %v", err)
	}

	jwtToken, err := tokenService.VerifyJWT(token)
	if err != nil {
		t.Fatalf("VerifyJWt expected no error, got %v", err)
	}

	user, err := tokenService.ParseJWT(jwtToken)
	if err != nil {
		t.Fatalf("ParseJWT expected no error, got %v", err)
	}

	if role, ok := models.LookUpRole(user.Role); user.Username != username && ok && role == basicRole {
		t.Fatalf("Usernanme values should match with the content in the token")
	}
}

func TestNewUserService_Create(t *testing.T) {
	saveFunc := func(param models.CreateUser) (*models.User, error) {
		return &models.User{
			Username: param.Username,
			Password: param.Password,
			Role:     param.Role,
		}, nil
	}

	userService := NewUserServiceImpl(&MockUserRepository{
		SaveFunc: saveFunc,
	}, NewTokenServiceImpl())

	created, err := userService.Create(context.Background(), models.CreateUser{
		Username: username,
		Password: password,
		Role:     basicRole,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	hash, err := GenerateHashPassword(password)

	if err != nil {
		t.Fatalf("Generte hash expected no error, got %v", err)
	}

	if created.Username != username && created.Role == basicRole && created.Password == hash {
		t.Fatalf("User Fields values should match with the created value")
	}
}

func TestNewUserService_Get(t *testing.T) {
	id := uuid.New()

	expectedUser := &models.User{
		ID:       id,
		Username: "",
		Password: "",
		Role:     0,
	}

	getFuncId := func(userId uuid.UUID) (*models.User, error) {
		if userId != id {
			return nil, errors.New("user not found")
		}

		return expectedUser, nil
	}

	userService := NewUserServiceImpl(&MockUserRepository{
		GetByIDFunc: getFuncId,
	}, NewTokenServiceImpl())

	user, err := userService.Get(context.Background(), id)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user != expectedUser {
		t.Errorf("expected user to be %v, got %v", expectedUser, user)
	}
}

func TestGenerateHashPassword(t *testing.T) {

	expectedHash, err := GenerateHashPassword(password)

	if err != nil {
		t.Fatalf("Generte hash expected no error, got %v", err)
	}

	if err != nil {
		t.Fatalf("Generte hash function expected no error, got %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(expectedHash), []byte(password))

	if err != nil {
		t.Fatalf("Compare hash returned an error")
	}
}
