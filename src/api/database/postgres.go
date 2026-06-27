package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ayush-0192/SyncPad/src/api/config"
)

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHOST,
		cfg.DBPORT,
		cfg.DBUSER,
		cfg.DBPASS,
		cfg.DBNAME,
		cfg.DBSSL,

	)

	return gorm.Open(postgres.Open(dsn),&gorm.Config{})
}