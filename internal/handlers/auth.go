package handlers

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(200, "login.tmpl", gin.H{})
		return
	}
	if c.Request.Method == "POST" {
		c.HTML(200, "login.tmpl", gin.H{})
		return
	}
}
