package routes

import (
	"github.com/akshrana12/Library-Management/controller"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	bookroute := router.Group("api/books")
	bookroute.POST("/addBook", controller.AddBook)
	bookroute.GET("/getAllBooks", controller.ShowBooks)
	bookroute.GET("/getBook/:id", controller.GetBookById)
	bookroute.GET("/getBook/genre/:genre", controller.GetBookByGenre)
}
