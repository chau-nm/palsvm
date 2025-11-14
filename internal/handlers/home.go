package handlers

import (
	"github.com/gin-gonic/gin"
)

func HomeViewHandler(c *gin.Context) {
	renderPage(c, "templates/page/index.tmpl", gin.H{})
}
