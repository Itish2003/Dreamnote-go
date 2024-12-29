package service

import (
	"errors"

	"gorm.io/gorm"
	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/models"
)

func GetUserDetails(user *string) (models.User, error) {
	var usr models.User
	// Query the database to find a user with the specified email
	err := initializers.DB.Where("email = ?", user).First(&usr).Error
	if err != nil {
		// If an error occurs (no user found), return an error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return usr, nil // User not found
		}
		return usr, err // Database error
	}
	return usr, nil // Return the user object if found
}
