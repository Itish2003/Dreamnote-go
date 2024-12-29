package route

import (
	"github.com/gin-gonic/gin"
	"itish.github.io/dreamnote/controller"
)

func Route(r *gin.Engine) *gin.RouterGroup {
	routes := r.Group("/v1")
	{
		routes.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "base endpoint working",
			})
		})
		routes.POST("/getDetails", controller.GetUser)
		routes.POST("/signup", controller.SignUp)
		routes.POST("/login", controller.Login)
		routes.PUT("/update", controller.Update)
		routes.POST("/uploadImage", controller.UploadImage)
		routes.POST("/getBlogs", controller.GetBlogs)
		routes.POST("/updateBlog", controller.UpdateBlogs) //future update, the backend is made but not the frontend.
		routes.POST("/deleteBlog", controller.DeleteBlogs)
		routes.POST("/createBlog", controller.CreateBlogs)
	}

	return routes
}
