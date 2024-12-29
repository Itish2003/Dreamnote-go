package service

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/models"
)

func CreateUser(user *models.User) error {
	var existingUser models.User
	result := initializers.DB.Where("email=?", user.Email).First(&existingUser)

	if result.Error == nil {
		log.Println("User might already exist...")
		return fmt.Errorf("user might already exits")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err)
		return err
	}
	user.Password = string(hashedPassword) // Replace plaintext password with hashed version

	if err := initializers.DB.Create(user).Error; err != nil { // creating user check in supabase
		log.Println("User not created in the database layer...", err)
		return err
	}
	log.Println("User successfully created in the database layer...")
	return nil
}