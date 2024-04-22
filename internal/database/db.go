package database

import (
	"fmt"
	"log"
	"time"

	"aitrainer-api/config"
	"aitrainer-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	var database *gorm.DB
	var err error

	db_hostname := cfg.Postgress.Host
	db_name := cfg.Postgress.Name
	db_user := cfg.Postgress.User
	db_pass := cfg.Postgress.Password
	db_port := cfg.Postgress.Port

	dbURl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_user, db_pass, db_hostname, db_port, db_name)

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(postgres.Open(dbURl), &gorm.Config{})
		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize database. Retrying...", i)
			time.Sleep(3 * time.Second)
		}
	}
	database.AutoMigrate(&models.Book{})
	database.AutoMigrate(&models.User{})

	DB = database
}
