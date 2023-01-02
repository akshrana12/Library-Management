package controller

import (
	"fmt"
	"net/http"

	"github.com/akshrana12/Library-Management/config"
	models "github.com/akshrana12/Library-Management/model"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.Users
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user)
	err := config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &user)
}

func ShowUsers(c *gin.Context) {
	fmt.Println("inside")
	var users []models.Users
	result := config.DB.Find(&users).Error
	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
		return
	}
	c.JSON(200, &users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	var result models.Users
	err := config.DB.Where("student_id = ?", id).First(&result)
	if err.RowsAffected == 0 {
		c.String(200, "No Record Found with this id")
		return
	}
	c.JSON(200, &result)
}
