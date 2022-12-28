package routes

import (
	"github.com/akshrana12/Library-Management/controller"
	"github.com/gin-gonic/gin"
)

func BookIssueRoutes(router *gin.Engine) {
	bookroute := router.Group("api/action")
	bookroute.POST("/issueBook", controller.IssueBook)
	bookroute.DELETE("/returnBook", controller.ReturnBook)
	bookroute.GET("/getUserDetail", controller.GetUserDetail)

}
