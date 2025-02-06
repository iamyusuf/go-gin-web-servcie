package handlers

import (
	"github.com/gin-gonic/gin"
	"my-service/internal/services"
	"net/http"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: *service,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	if token, err := h.authService.Login(c); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}
