package service

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// UploadImageService handles file upload to Supabase via S3-compatible storage
func UploadImageService(file *multipart.FileHeader) (string, error) {
	// Open the file for reading
	src, err := file.Open()
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return "", fmt.Errorf("failed to open file")
	}
	log.Println("File opened successfully")
	defer src.Close()

	// Generate a unique file name
	fileID := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	log.Printf("Generated unique file name: %s\n", fileID)

	// Set up AWS session using the S3-compatible endpoint and credentials
	log.Println("Creating AWS session using S3-compatible endpoint...")
	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String(os.Getenv("SUPABASE_REGION")),
		Endpoint:         aws.String(os.Getenv("SUPABASE_S3_ENDPOINT")),
		DisableSSL:       aws.Bool(true), // Set to true for non-SSL endpoint
		Credentials:      credentials.NewStaticCredentials(os.Getenv("SUPABASE_ACCESS_KEY"), os.Getenv("SUPABASE_SECRET_KEY"), ""),
		S3ForcePathStyle: aws.Bool(true), // Required for S3-compatible services
	})
	if err != nil {
		log.Printf("Error creating AWS session: %v\n", err)
		return "", fmt.Errorf("failed to create AWS session")
	}
	log.Println("AWS session created successfully")

	// Create an S3 client
	svc := s3.New(sess)

	// Prepare upload input for S3
	uploadInput := &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("SUPABASE_BUCKET")), // The name of your Supabase bucket
		Key:         aws.String(fileID),                       // Unique file name
		Body:        src,                                      // File content
		ACL:         aws.String("public-read"),                // Set file visibility (can be "private" or "public-read")
		ContentType: aws.String(file.Header.Get("Content-Type")),
	}

	// Debug: Log the details of the upload
	log.Printf("Uploading file to bucket: %s, file key: %s\n", os.Getenv("SUPABASE_BUCKET"), fileID)

	// Upload the file to S3-compatible storage
	_, err = svc.PutObject(uploadInput)
	if err != nil {
		log.Printf("Error uploading file to S3: %v\n", err)
		return "", fmt.Errorf("failed to upload file")
	}
	log.Printf("File uploaded successfully: %s\n", fileID)

	// Construct the public URL for the file (if you want public access to the file)
	fileURL := fmt.Sprintf("%s/object/public/%s/%s", os.Getenv("SUPABASE_S3_URL"), os.Getenv("SUPABASE_BUCKET"), fileID)
	log.Printf("File URL: %s\n", fileURL)

	// Return the file URL
	return fileURL, nil
}
