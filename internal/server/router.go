package server

import (
	"github.com/chau-nm/palsvm/internal/handlers"
)

func (s *Server) setupRoutes() {
	s.router.GET("/health", handlers.CheckHealth)

	// auth
	s.router.GET("/login", authorizedMiddleware(), handlers.LoginViewHandler)
	s.router.POST("/login", authorizedMiddleware(), handlers.LoginHandler)
	s.router.GET("/logout", handlers.LogoutHandler)

	// dashboard
	s.router.GET("/", authRequiredMiddleware(), handlers.HomeViewHandler)
	s.router.GET("/palworld-setting", authRequiredMiddleware(), handlers.PalworldSettingViewHandler)
}
