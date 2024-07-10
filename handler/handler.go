package handler

import (
	"github.com/odeassis/goodoor/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Looger
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetPostgreSQL()
}
