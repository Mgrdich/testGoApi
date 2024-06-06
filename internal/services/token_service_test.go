package services

import (
	"testing"

	"github.com/google/uuid"
	"testGoApi/configs"
	"testGoApi/internal/models"
	"testGoApi/internal/util"
)

func TestNewTokenService(t *testing.T) {
	configs.SetAppConfig(&configs.AppConfig{
		Environment:            util.DevEnvironment,
		TokenExpirationMinutes: 10,
		JwtSecretKey:           []byte("DummyJWTSecretKey"),
	})

	tokenService := NewTokenServiceImpl()

	id := uuid.New()
	role := models.LookUpRoleString(models.AdminRole)
	expectedUser := &models.TokenizedUser{ID: id, Username: username, Role: role}

	token, err := tokenService.GenerateJWT(expectedUser)

	if err != nil {
		t.Fatalf("Generte JWT expected no error, got %v", err)
	}

	jwtToken, err := tokenService.VerifyJWT(token)

	if err != nil {
		t.Fatalf("Verify JWT expected no error, got %v", err)
	}

	user, err := tokenService.ParseJWT(jwtToken)

	if err != nil {
		t.Fatalf("Parse JWT expected no error, got %v", err)
	}

	if user.ID != expectedUser.ID {
		t.Fatalf("ID from the token should match the generated one %v  %v", user.ID, expectedUser.ID)
	}

	if user.Username != expectedUser.Username || user.ID != expectedUser.ID || user.Role != expectedUser.Role {
		t.Fatalf("Username from the token should match the generated one %v  %v", user.Username, expectedUser.Username)
	}

	if user.Role != expectedUser.Role {
		t.Fatalf("Role from the token should match the generated one %v  %v", user.Role, expectedUser.Role)
	}
}
