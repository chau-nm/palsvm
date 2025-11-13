package server

import (
	"net/http"

	"github.com/chau-nm/palsvm/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// authorizedMiddleware prevents authenticated users from accessing login/register pages
// Redirects logged-in users to dashboard/home page
func authorizedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.Next()
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

// authRequiredMiddleware ensures only authenticated users can access protected routes
// Redirects unauthenticated users to login page
func authRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		username := session.Get("username")
		if username == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		if username != config.PalSvmServerConfig.ServerUsername {
			session.Clear()
			_ = session.Save()
			c.Redirect(http.StatusFound, "/login")
		}
		c.Next()
	}
}
