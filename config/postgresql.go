package config

import (
	"github.com/odeassis/goodoor/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgreSQL() (*gorm.DB, error) {

	logger := GetLogger("postgresql")

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5300 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("postgreSQL opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("postgreSQL automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
