package global

import (
	"blog/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

/**
*@Author: whh
*@CreateTime: 2023-03-04  21:38
 */

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
