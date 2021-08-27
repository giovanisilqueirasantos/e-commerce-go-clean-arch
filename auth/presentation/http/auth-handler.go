package http

import (
	"github.com/skeey/e-commerce-go-clean-arch/domain"
)

type AuthHandler struct {
	AuthUseCase domain.AuthUseCase
}

// func NewAuthHandler(c *http.ServeMux, auc domain.AuthUseCase) {
// 	c.HandleFunc("")

// 	handler := &authHandler{
// 		AuthUseCase: auc,
// 	}
// }
