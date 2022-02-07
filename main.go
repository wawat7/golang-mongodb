package main

import (
	"github.com/gin-gonic/gin"
	"golang-mongodb/app"
	"golang-mongodb/user"
)

func main() {

	// Setup DB
	database := app.NewMongoDatabase()

	userRepository := user.NewRepository(database)
	userService := user.NewService(userRepository)
	userController := user.NewController(userService)

	route := gin.Default()

	userController.Route(route)

	route.Run(":8080")

}
