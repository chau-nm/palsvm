package handlers

import (
	"github.com/gin-gonic/gin"
)

func PalworldSettingViewHandler(c *gin.Context) {
	renderPage(c, "templates/page/palworld_setting.tmpl", gin.H{})
}
