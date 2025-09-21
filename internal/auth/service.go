package auth

import (
	"errors"
	"time"

	"github.com/analopesdev/duochat-service/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ValidateToken(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Values.AuthSecret), nil
	})

	return token.Valid
}

func (s *Service) GenerateJWT(sub string, username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      sub,
		"username": username,
		"exp":      time.Now().Add(30 * 24 * time.Hour).Unix(), // 30 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(config.Values.AuthSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Values.AuthSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}
