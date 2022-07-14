package routes

import "github.com/gin-gonic/gin"
import "github.com/raulkkonen/first_gin/controller"

func UserRoute(router *gin.Engine){
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUser)
	router.DELETE("/:id", controller.DeleteUser)
	router.PUT("/:id", controller.UpdateUser)
}