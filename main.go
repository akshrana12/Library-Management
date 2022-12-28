package main

import (
	"github.com/akshrana12/Library-Management/config"
	"github.com/akshrana12/Library-Management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.BookIssueRoutes(router)
	router.Run()
}
