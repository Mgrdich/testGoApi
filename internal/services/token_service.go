package services

import (
	"encoding/json"
	"errors"
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

func (s *TokenServiceImpl) GenerateJWT(user *models.TokenizedUser) (string, error) {
	expirationToken := time.Duration(s.appConfig.TokenExpirationMinutes)

	// Marshal the user data to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": string(userJSON),
		"exp":  time.Now().Add(expirationToken * time.Minute).Unix(),
	})

	tokenString, err := token.SignedString(s.appConfig.JwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *TokenServiceImpl) VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return s.appConfig.JwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func (s *TokenServiceImpl) ParseJWT(token *jwt.Token) (*models.TokenizedUser, error) {
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userJSON, ok := claims["user"].(string)

		if !ok {
			return nil, errors.New("invalid user claim")
		}

		var user *models.TokenizedUser
		err := json.Unmarshal([]byte(userJSON), &user)

		if err != nil {
			return nil, err
		}

		return user, nil
	}

	return nil, errors.New("claiming token failed")
}
