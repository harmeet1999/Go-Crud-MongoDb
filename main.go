package main

import (
	"gotest/connections"
	"gotest/controllers"
	errorHandler "gotest/middlerware"

	// "gotest/cronJobs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// cronJobs.StartCronJob()

	r := gin.Default()
	connections.ConnectDb()
	r.Use(errorHandler.ErrorHandlerMiddleware())
	r.GET("/all", controllers.GetAllData)
	r.POST("/post", controllers.CreateUser)
	r.PUT("/update/:id", controllers.UpdateProduct)
	r.DELETE("/delete/:id", controllers.DeleteProductById)

	r.Run(":" + os.Getenv("PORT"))

}
