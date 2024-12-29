package service

import (
	"fmt"
	"log"

	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/models"
)

func UserUpdate(user *models.User) error {
	var existingUser models.User
	result := initializers.DB.Where("email=?", user.Email).First(&existingUser)
	if result.Error != nil {
		log.Println("User doesn't exist...")
		return fmt.Errorf("user doesn't exits")
	}
	log.Println("User does exist...")

	existingUser.Name = user.Name
	existingUser.Age = user.Age
	existingUser.Sex = user.Sex
	existingUser.Linkedin = user.Linkedin
	existingUser.Instagram = user.Instagram
	existingUser.Bio = user.Bio
	existingUser.Github = user.Github
	existingUser.Photo = user.Photo

	// Save the updated user object
	updateResult := initializers.DB.Save(&existingUser)
	if updateResult.Error != nil {
		log.Println("Error updating user:", updateResult.Error)
		return fmt.Errorf("could not update user")
	}

	log.Println("User updated successfully")
	return nil
}
