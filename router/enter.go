package router

import (
	"blog/api"
	"blog/global"
	"github.com/gin-gonic/gin"
)

/**
*@Author: whh
*@CreateTime: 2023-03-06  21:49
 */

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("", settingsApi.SettingsInfoView)
	return router
}
