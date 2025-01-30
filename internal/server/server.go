package server

import (
	"gorm.io/driver/postgres"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewServer(dsn string) (*Server, error) {
	// Initialize database
	db, err := initializeDB(dsn)
	if err != nil {
		return nil, err
	}

	// Create server instance
	server := &Server{
		DB:     db,
		Router: gin.Default(),
	}

	// Setup routes
	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	// Health check route
	s.Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	// Add your other routes here
	// Example:
	// userRoutes := s.Router.Group("/users")
	// userRoutes.GET("/", handlers.GetUsers)
}

func (s *Server) Run(addr string) error {
	return s.Router.Run(addr)
}

func initializeDB(dsn string) (*gorm.DB, error) {
	// Initialize your database connection here
	// Example for PostgreSQL:
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run migrations
	// db.AutoMigrate(&models.User{})

	log.Println("Connected to database")
	return db, nil
}
