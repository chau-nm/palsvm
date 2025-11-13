package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PalworldSettingViewHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}
