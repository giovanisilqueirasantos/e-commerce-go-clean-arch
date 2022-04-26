package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type authService struct{}

func NewAuthService() *authService {
	return &authService{}
}

func (a authService) EncodePass(ctx context.Context, pass string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes)
}

func (a authService) PassIsEqualHashedPass(ctx context.Context, pass string, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
