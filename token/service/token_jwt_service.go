package service

import (
	"context"
	"time"

	"github.com/giovanisilqueirasantos/e-commerce-go-clean-arch/domain"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Info string
	jwt.StandardClaims
}

type tokenService struct{}

func NewTokenService() *tokenService {
	return &tokenService{}
}

func (t *tokenService) Sign(ctx context.Context, info domain.TokenInfo, expirationInMinutes int64) (domain.Token, error) {
	expirationTime := time.Now().Add(time.Duration(expirationInMinutes) * time.Minute)

	claims := &Claims{
		Info: info.Info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return domain.Token(tokenString), err
}

func (t *tokenService) IsValid(ctx context.Context, token domain.Token) (domain.IsValid, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(string(token), claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if !tkn.Valid {
		return false, nil
	}

	return true, nil
}
