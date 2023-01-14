package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabse()
}

func main() {

	router := gin.Default()

	// auth
	router.POST("/singup", controllers.SingUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// post
	router.POST("/posts", middleware.RequireAuth, controllers.PostCreate)
	router.PUT("/posts/:id", middleware.RequireAuth, controllers.PostUpdate)
	router.DELETE("/posts/:id", middleware.RequireAuth, controllers.PostDelete)
	router.GET("/posts", controllers.PostIndex)
	router.GET("/posts/:id", controllers.PostShow)
	router.Run()
}
