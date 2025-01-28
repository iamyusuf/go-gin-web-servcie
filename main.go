package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Config struct {
	port string
}

type Server struct {
	Config
	engine *gin.Engine
}

func (s *Server) Start() error {
	s.registerRoutes()
	return s.engine.Run(":8080")
}

func (s *Server) registerRoutes() {
	s.engine.GET("/hello", helloRouteHandler)
}

func NewServer(config Config) *Server {
	engine := gin.New()
	return &Server{engine: engine, Config: config}
}

func main() {
	server := NewServer(Config{port: ":8080"})
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func helloRouteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
