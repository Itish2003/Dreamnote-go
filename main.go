package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"itish.github.io/dreamnote/initializers"
	"itish.github.io/dreamnote/middleware"
	"itish.github.io/dreamnote/route"
)

func init() {
	// if err := initializers.LoadEnv(); err != nil {
	// 	log.Fatalf("[CRITICAL] Failed to load env: %s", err)
	// }
	if err := initializers.ConnectUser(); err != nil {
		log.Fatalf("[CRITICAL] Failed to initialize database connection: %s", err)
	}
	if err := initializers.Migrate(initializers.DB); err != nil {
		log.Fatalf("[CRITICAL] Failed to make migrate models on supabase: %s", err)
	}
}

func main() {
	router := gin.New() // creating a router

	router.Use(gin.Logger()) // creating a logger for debug

	router.Use(middleware.CORSMiddleware()) // adding middleware for cors

	f, err := os.Create("logFile.log") // creating a logger file to store the logger output
	if err != nil {
		log.Println("Trouble creating a logger")
	}
	defer f.Close()

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout) // connecting gin default writer to write to file and and terminal

	route.Route(router) // accessing the endpoints

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	router.Run(":" + port) // default port: 8080
}
