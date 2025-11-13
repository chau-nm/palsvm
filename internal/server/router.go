package server

import (
	"github.com/chau-nm/palsvm/internal/handlers"
)

func (s *Server) setupRoutes() {
	s.router.GET("/health", handlers.CheckHealth)

	s.setupViewRouter()
}

func (s *Server) setupViewRouter() {
	s.router.GET("/login", handlers.Login)
}
