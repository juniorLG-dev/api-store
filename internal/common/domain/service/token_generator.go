package service

import (
	"loja/internal/common/domain/entities"
	"loja/internal/configuration/handler_err"
	
	"github.com/golang-jwt/jwt"

	"os"
	"time"
	"strings"
	"errors"
)

type TokenInfoDTO struct {
	ID string
	Name string
	Username string
	Type string
}

type TokenGenerator struct {
	Value string
}

func NewTokenGenerator(typeUser string) *TokenGenerator {
	return &TokenGenerator{
		Value: typeUser,
	}
}

func (t *TokenGenerator) GenerateToken(user *entities.User) (string, *handler_err.InfoErr) {
	secretKey := os.Getenv("SECRET_KEY")

	claims := jwt.MapClaims{
		"id": user.GetID(),
		"name": user.GetName(),
		"username": user.GetUsername(),
		"type": t.Value,
		"exp": time.Now().Add(time.Hour * 5).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", &handler_err.InfoErr{
			Message: "error creating jwt",
			Err: handler_err.ErrInternal,
		}
	}

	return tokenString, &handler_err.InfoErr{}
}

func (t *TokenGenerator) VerifyToken(tokenValue string) (TokenInfoDTO, *handler_err.InfoErr) {
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(t.removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}

		return nil, errors.New("invalid token")
	})
	if err != nil {
		return TokenInfoDTO{}, &handler_err.InfoErr{
			Message: err.Error(),
			Err: handler_err.ErrInvalidInput,
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return TokenInfoDTO{}, &handler_err.InfoErr{
			Message: "invalid token",
			Err: handler_err.ErrInvalidInput,
		}
	}

	return TokenInfoDTO{
		ID: claims["id"].(string),
		Name: claims["name"].(string),
		Username: claims["username"].(string),
		Type: claims["type"].(string),
	}, &handler_err.InfoErr{}
}

func (t *TokenGenerator) removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}