package main

import "github.com/gin-gonic/gin"
import "github.com/raulkkonen/first_gin/routes"
import "github.com/raulkkonen/first_gin/config"

func main()  {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}