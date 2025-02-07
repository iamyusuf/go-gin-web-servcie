package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
	"my-service/internal/models"
	"my-service/internal/services"
	"net/http"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	if users, err := h.userService.GetUsers(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

type user struct {
	Name     string `json:"name" binding:"required" required:"$field is required"`
	Password string `json:"password" binding:"required" required:"$field is required"`
	Email    string `json:"email" binding:"required" required:"$field is required"`
}

var (
	g         = galidator.G()
	validator = g.Validator(models.User{})
)

func (h *UserHandler) CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": validator.DecryptErrors(err)})
	} else if userId, err := h.userService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"id": userId})
	}
}

func (h *UserHandler) FindUser(c *gin.Context) {
	if user, err := h.userService.GetUser(c); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	if err := h.userService.UpdateUser(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	if err := h.userService.DeleteUser(c); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
