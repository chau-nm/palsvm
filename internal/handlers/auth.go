package handlers

import (
	"fmt"
	"net/http"

	"github.com/chau-nm/palsvm/internal/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginTemplate = "login.tmpl"
)

func LoginViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, loginTemplate, gin.H{})
}

func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username != config.PalSvmServerConfig.ServerUsername || password != config.PalSvmServerConfig.ServerPassword {
		c.HTML(http.StatusUnauthorized, loginTemplate, gin.H{
			"Message": "Invalid username or password",
		})
		return
	}
	session.Set("username", username)
	if err := session.Save(); err != nil {
		fmt.Printf("Error saving session: %v\n", err)
		c.HTML(http.StatusUnauthorized, loginTemplate, gin.H{
			"Message": "Something went wrong",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}
	c.Redirect(http.StatusFound, "/login")
}
