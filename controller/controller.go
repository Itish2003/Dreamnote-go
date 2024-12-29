package controller

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"itish.github.io/dreamnote/models"
	"itish.github.io/dreamnote/service"
)

func SignUp(c *gin.Context) {
	var user models.User                            // creating a variable of type User
	if err := c.ShouldBindJSON(&user); err != nil { // Binding the data coming from the request
		log.Println("Unable to Bind the data from frontend...")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input...",
		})
		return
	}
	log.Println("Bind Successfully from the frontend...")

	err := service.CreateUser(&user) // accessing the service layer and passing on the processed data to the service
	if err != nil {
		log.Println("Unable to create the user on the server side...")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to create user on the server side...",
		})
		return
	}
	log.Println("User created successfully on the server side...")
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var user models.User                            // creating a variable of type User
	if err := c.ShouldBindJSON(&user); err != nil { // Binding the data coming from the request
		log.Println("Unable to Bind the data from frontend...")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input...",
		})
		return
	}
	log.Println("Bind Successfully from the frontend...")

	err := service.UserLogin(&user)
	if err != nil {
		log.Println("Unable to login the user on the server side...")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to login user on the server side...",
		})
		return
	}
	log.Println("User successfully logged in on the server side...")
	c.JSON(http.StatusOK, gin.H{"message": "User login successfully"})
}

func Update(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error Binding the body...")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input...",
		})
		return
	}
	log.Println("Bind Successfully from the frontend...")
	err := service.UserUpdate(&user)
	if err != nil {
		log.Println("Unable to update the user on the server side...")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to update user on the server side...",
		})
		return
	}
	log.Println("User successfully updated on the server side...")
	c.JSON(http.StatusOK, gin.H{"message": "User update successfully"})
}

// future update, the backend is made but not the frontend.
func UpdateBlogs(c *gin.Context) {
	var request struct {
		Email   string `json:"email"`
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// Bind the JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Invalid request body...")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Convert ID to uuid.UUID
	blogID, err := uuid.Parse(request.ID)
	if err != nil {
		log.Println("Invalid blog ID format...")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid blog ID"})
		return
	}

	// Update the blog
	err = service.UpdateBlog(request.Email, blogID, request.Title, request.Content)
	if err != nil {
		log.Println("Error updating the blog:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func CreateBlogs(c *gin.Context) {
	var requestPayload struct {
		Email string      `json:"email"`
		Blog  models.Blog `json:"blog"`
	}

	// Bind the incoming JSON payload to the struct
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		log.Println("Binding was not successful...")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	log.Println("Binding was successful...")
	log.Println("User Email:", requestPayload.Email)
	log.Println("Blog Title:", requestPayload.Blog.Title)
	log.Println("Blog Content:", requestPayload.Blog.Content)
	log.Println("Blog ID (as string):", requestPayload.Blog.ID)
	log.Println("Blog ID (type):", reflect.TypeOf(requestPayload.Blog.ID))
	log.Println("Blog User ID:", requestPayload.Blog.UserID)
	log.Println("Blog CreatedAt:", requestPayload.Blog.CreatedAt)
	log.Println("Blog UpdatedAt:", requestPayload.Blog.UpdatedAt)
	log.Println("Blog DeletedAt:", requestPayload.Blog.DeletedAt)

	// Convert the Blog ID from string to uuid.UUID if it's not the zero value
	if requestPayload.Blog.ID == uuid.Nil {
		// If the ID is the zero value, it was likely not provided in the request, so generate a new one
		requestPayload.Blog.ID = uuid.New()
	} else {
		// If the ID is provided as string, ensure it's parsed into a valid UUID
		parsedID, err := uuid.Parse(requestPayload.Blog.ID.String())
		if err != nil {
			log.Println("Invalid UUID format for Blog ID:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Blog ID format"})
			return
		}
		requestPayload.Blog.ID = parsedID
	}
	log.Println("Blog ID (after parsing):", requestPayload.Blog.ID)
	log.Println("Blog ID (after parsing type):", reflect.TypeOf(requestPayload.Blog.ID))

	// Call the service to create the blog
	createdBlog, err := service.CreateBlog(requestPayload.Email, &requestPayload.Blog)
	if err != nil {
		log.Println("Error creating blog:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating blog"})
		return
	}
	log.Println("Blog created successfully... (controller)")

	// Respond with the created blog details
	c.JSON(http.StatusOK, createdBlog)

}

func DeleteBlogs(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		ID    string `json:"id"`
	}

	// Bind the JSON request
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Invalid request body...")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Convert ID to uuid.UUID
	blogID, err := uuid.Parse(request.ID)
	if err != nil {
		log.Println("Invalid blog ID format...")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid blog ID"})
		return
	}

	// Delete the blog
	err = service.DeleteBlog(request.Email, blogID)
	if err != nil {
		log.Println("Error deleting the blog:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func GetBlogs(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Invalid request body...")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	log.Println(user.Email)

	blogs, err := service.GetAllBlogs(user.Email)
	if err != nil {
		log.Println("Error fetching blogs:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func GetUser(c *gin.Context) {
	var userRequest struct {
		Email string `json:"email" binding:"required"`
	}

	// Bind the request body to get the email
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Println("Error Binding the body...")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input. Email is required.",
		})
		return
	}

	log.Println("Bind Successfully from the frontend...")

	var existingUser models.User
	// Fetch the user from the database
	existingUser, err := service.GetUserDetails(&userRequest.Email)
	if err != nil {
		log.Println("Unable to get the user on the server side...")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to get user on the server side.",
		})
		return
	}

	log.Println("Got User successfully on the server side...")

	// Return only the necessary fields to the frontend
	c.JSON(http.StatusOK, gin.H{
		"name":      existingUser.Name,
		"email":     existingUser.Email,
		"password":  existingUser.Password,
		"age":       existingUser.Age,
		"sex":       existingUser.Sex,
		"bio":       existingUser.Bio,
		"linkedin":  existingUser.Linkedin,
		"instagram": existingUser.Instagram,
		"github":    existingUser.Github,
		"photo":     existingUser.Photo,
	})
}

func UploadImage(c *gin.Context) {
	// Parse the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error parsing file:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse file"})
		return
	}

	// Call the service to upload the file
	fileURL, err := service.UploadImageService(file)
	if err != nil {
		log.Println("Error uploading file:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	// Respond with the file URL
	c.JSON(http.StatusOK, gin.H{"url": fileURL})
}
