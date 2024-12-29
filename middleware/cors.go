package middleware

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set allowed origins (your frontend services)
		allowedOrigins := []string{
			"",                      // Production React frontend
			"http://localhost:5173", // Local React development
		}

		// Check the Origin header from the request
		origin := c.Request.Header.Get("Origin")
		isAllowed := false

		// Validate if the origin is allowed
		for _, o := range allowedOrigins {
			if o == origin {
				isAllowed = true
				break
			}
		}

		// Set CORS headers if the origin is allowed
		if isAllowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		// Common CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests (OPTIONS method)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		c.Next()
	}
}
