package service

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/models"
)

func UserLogin(user *models.User) error {
	var existingUser models.User                                                // creating a variable to hold record
	result := initializers.DB.Where("email=?", user.Email).First(&existingUser) // searching the email for a valid entry then storing the data from that entry

	if result.Error == nil {
		log.Println("User email exists...")

		// Compare the provided password with the hashed password
		if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err == nil {
			log.Println("Log in successful...")
			return nil
		}
	}
	log.Println("Incorrect Password...")
	return errors.New("invalid email or password")
}
