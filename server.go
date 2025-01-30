package main

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	port string
}

type Server struct {
	Config
	Router *gin.Engine
}

func (s *Server) Start() error {
	s.registerRoutes()
	return s.Router.Run(":8080")
}

func (s *Server) registerRoutes() {
	s.Router.GET("/hello", Wrap(helloRouteHandler))
}

func NewServer(config Config) *Server {
	engine := gin.New()
	return &Server{Router: engine, Config: config}
}
