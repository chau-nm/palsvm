package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
		return
	}
	if c.Request.Method == "POST" {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{})
		return
	}
}
