package global

import (
	"blog/config"
	"gorm.io/gorm"
)

/**
*@Author: whh
*@CreateTime: 2023-03-04  21:38
 */

var (
	Config *config.Config
	DB     *gorm.DB
)
