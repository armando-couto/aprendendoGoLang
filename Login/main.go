package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"login/controllers"
	"login/database"
	"login/middlewares"
	"login/models"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	api := r.Group("/api")
	{
		public := api.Group("/public")
		{
			public.POST("/login", controllers.Login)
			public.POST("/signup", controllers.Signup)
		}

		// here
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			protected.GET("/profile", controllers.Profile)
		}
	}

	return r
}
func main() {
	err := database.InitDatabase()
	if err != nil {
		log.Fatalln("could not create database", err)
	}

	database.DB.AutoMigrate(&models.User{})

	r := setupRouter()
	r.Run(":8080")
}