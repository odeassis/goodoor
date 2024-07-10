package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Looger
)

func Init() error {
	var err error

	db, err = InitializePostgreSQL()

	if err != nil {
		return fmt.Errorf("error initalize postgreSQL: %v", err)
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}

func GetLogger(p string) *Looger {
	logger = NewLogger(p)
	return logger
}
