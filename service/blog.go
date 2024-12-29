package service

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/models"
)

// CreateBlog creates a new blog post for the user identified by their email.
func CreateBlog(email string, blog *models.Blog) (*models.Blog, error) {
	// Find the user by email
	var user models.User
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("User not found:", email)
		return nil, errors.New("user not found")
	}
	log.Println("Data was successfully binded...")
	// Set the user ID for the blog
	blog.UserID = user.ID
	log.Println("Blog User ID is set:", blog.UserID)

	// Generate a new UUID for the blog (if not already provided)
	log.Println("Initial Blog ID generated by the frontend", blog.ID)
	if blog.ID == uuid.Nil {
		blog.ID = uuid.New()
		log.Println("Blog ID generated by the server-side if empty", blog.ID)
	}

	// Create the blog record in the database
	if err := initializers.DB.Create(&blog).Error; err != nil {
		log.Println("Type:", reflect.TypeOf(blog.ID))
		log.Println("Error creating blog:", err)
		return nil, errors.New("error creating blog")
	}
	log.Println("Blog created successfully...")

	// Return the created blog
	return blog, nil
}

func GetAllBlogs(email string) ([]models.Blog, error) {
	var existingUser models.User
	if err := initializers.DB.Where("email=?", email).First(&existingUser).Error; err != nil {
		log.Println("Having trouble finding the user...")
		return nil, fmt.Errorf("error finding the user")
	}
	log.Println("User successfully found...")
	// Fetch all blogs linked to the user's ID
	var blogs []models.Blog
	if err := initializers.DB.Where("user_id = ?", existingUser.ID).Find(&blogs).Error; err != nil {
		log.Println("Error fetching blogs...")
		return nil, fmt.Errorf("error fetching blogs: %v", err)
	}
	log.Println("Blogs successfully fetched:", blogs)

	return blogs, nil
}

func DeleteBlog(email string, blogID uuid.UUID) error {
	var user models.User

	// Find the user by email
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("Error finding the user...")
		return fmt.Errorf("user not found: %v", err)
	}
	log.Println("User found with ID:", user.ID)

	// Find the blog by ID and verify it belongs to the user
	var blog models.Blog
	if err := initializers.DB.Where("id = ? AND user_id = ?", blogID, user.ID).First(&blog).Error; err != nil {
		log.Println("Error finding the blog or blog doesn't belong to the user...")
		return fmt.Errorf("blog not found or unauthorized: %v", err)
	}
	log.Println("Blog found with ID:", blog.ID)

	// Delete the blog
	if err := initializers.DB.Delete(&blog).Error; err != nil {
		log.Println("Error deleting the blog...")
		return fmt.Errorf("error deleting the blog: %v", err)
	}
	log.Println("Blog deleted successfully!")
	return nil
}

// future update, the backend is made but not the frontend.
func UpdateBlog(email string, blogID uuid.UUID, newTitle string, newContent string) error {
	var user models.User

	// Find the user by email
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("Error finding the user...")
		return fmt.Errorf("user not found: %v", err)
	}
	log.Println("User found with ID:", user.ID)

	// Find the blog by ID and verify it belongs to the user
	var blog models.Blog
	if err := initializers.DB.Where("id = ? AND user_id = ?", blogID, user.ID).First(&blog).Error; err != nil {
		log.Println("Error finding the blog or blog doesn't belong to the user...")
		return fmt.Errorf("blog not found or unauthorized: %v", err)
	}
	log.Println("Blog found with ID:", blog.ID)

	// Update the blog fields
	blog.Title = newTitle
	blog.Content = newContent

	if err := initializers.DB.Save(&blog).Error; err != nil {
		log.Println("Error updating the blog...")
		return fmt.Errorf("error updating the blog: %v", err)
	}
	log.Println("Blog updated successfully!")
	return nil
}