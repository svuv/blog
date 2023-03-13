package settings_api

import "github.com/gin-gonic/gin"

/**
*@Author: whh
*@CreateTime: 2023-03-06  22:39
 */

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "xxx"})
}
