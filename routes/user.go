package routes

import (
	"github.com/akshrana12/Library-Management/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	bookroute := router.Group("api/user")
	bookroute.POST("/addUser", controller.AddUser)
	bookroute.GET("/getAllUsers", controller.ShowUsers)
	bookroute.GET("/getUser/:id", controller.GetUserById)
}
