package main

import (
	"github.com/gin-gonic/gin"
)

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
	s.engine.GET("/hello", Wrap(helloRouteHandler))
}

func NewServer(config Config) *Server {
	engine := gin.New()
	return &Server{engine: engine, Config: config}
}
