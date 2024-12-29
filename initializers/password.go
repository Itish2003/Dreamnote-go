package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // automigrate is also using this var

func ConnectUser() error {
	log.Println("Connecting to database")
	dsn := os.Getenv("DIRECT_URL")
	if dsn == "" {
		log.Println("DIRECT_URL variable not loading...")
		return fmt.Errorf("env variable showing empty")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	DB = DB.Debug()

	log.Println("Database connection successful")

	return nil
}
