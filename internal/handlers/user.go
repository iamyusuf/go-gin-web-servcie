package handlers

import (
	"github.com/gin-gonic/gin"
	"my-service/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	userId, err := h.userService.CreateUser(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": userId})
}

func (h *UserHandler) FindUser(c *gin.Context) {
	user, err := h.userService.GetUser(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	err := h.userService.UpdateUser(c)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
