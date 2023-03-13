package api

import "blog/api/settings_api"

/**
*@Author: whh
*@CreateTime: 2023-03-13  22:27
 */

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
}

var ApiGroupApp = new(ApiGroup)
