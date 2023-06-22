package database

import (
	"fmt"
	"github.com/farhodm/alif-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// DBInit initializes a database connection and gorm.DB object
func DBInit() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s timezone=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("TIMEZONE"),
		os.Getenv("DB_SSL_MODE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
