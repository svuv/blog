package router

import (
	"blog/api"
	"github.com/gin-gonic/gin"
)

/**
*@Author: whh
*@CreateTime: 2023-03-14  22:28
 */

func SettingsRouter(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
}
