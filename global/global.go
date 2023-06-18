package global

import (
	"gorm.io/gorm"
	"colly_music/config"
)

var (
	Config config.Config
	Db     *gorm.DB
)
