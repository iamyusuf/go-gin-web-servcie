package server

import (
	"gorm.io/driver/postgres"
	"log"
	"my-service/internal/handlers"
	"my-service/internal/models"
	"my-service/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer(dsn string) (*Server, error) {
	db, err := initializeDB(dsn)
	if err != nil {
		return nil, err
	}

	server := &Server{
		DB:     db,
		Router: gin.Default(),
	}

	server.setupRoutes()
	return server, nil
}

func (s *Server) setupRoutes() {
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	userHandler := handlers.NewUserHandler(services.NewUserService(s.DB))
	apis := s.Router.Group("/api")
	apis.POST("users", userHandler.CreateUser)
	apis.GET("users/:id", userHandler.FindUser)
	apis.PUT("users/:id", userHandler.UpdateUser)
	apis.DELETE("users/:id", userHandler.DeleteUser)
}

func (s *Server) Run(addr string) error {
	return s.Router.Run(addr)
}

func initializeDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	_ = db.AutoMigrate(&models.User{})

	log.Println("Connected to database")
	return db, nil
}
