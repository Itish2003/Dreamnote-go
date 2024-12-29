package initializers

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"itish.github.io/dreamnote/models"
)

func Migrate(db *gorm.DB) error { // accessing the global variable DB defined in password.go
	log.Println("Starting database migration...")
	err := db.AutoMigrate(&models.User{}, &models.Blog{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
		return fmt.Errorf("migration failed...")
	} else {
		log.Println("Migration completed successfully!")
	}
	return nil
}
