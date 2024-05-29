package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"testGoApi/configs"
	"testGoApi/internal/models"
)

type TokenServiceImpl struct {
	appConfig *configs.AppConfig
}

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{
		appConfig: configs.GetAppConfig(),
	}
}

func (s *TokenServiceImpl) GenerateJWT(user models.User) (string, error) {
	expirationToken := time.Duration(s.appConfig.TokenExpirationMinutes)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(expirationToken * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(s.appConfig.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenServiceImpl) VerifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.appConfig.JwtSecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
