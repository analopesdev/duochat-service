package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service struct {
	secretKey []byte
}

func NewService(secretKey []byte) *Service {
	return &Service{secretKey: secretKey}
}

func (s *Service) GenerateToken(userID string, nickname string) (string, error) {
	expiresAt := time.Now().Add(30 * 24 * time.Hour) // 30 dias

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		Audience:  jwt.ClaimStrings{nickname}, // opcional; poderia ser custom
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expiresAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
			}
			return s.secretKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	rc, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}
	if rc.ExpiresAt != nil && time.Now().After(rc.ExpiresAt.Time) {
		return nil, errors.New("token expirado")
	}

	nick := ""
	if len(rc.Audience) > 0 {
		nick = rc.Audience[0]
	}

	return &Claims{
		UserID:   uuid.MustParse(rc.Subject),
		Nickname: nick,
		Exp:      rc.ExpiresAt.Time.Unix(),
		Iat:      rc.IssuedAt.Time.Unix(),
	}, nil
}
