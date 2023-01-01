package controller

import (
	"net/http"

	"github.com/akshrana12/Library-Management/cache"
	"github.com/akshrana12/Library-Management/config"
	models "github.com/akshrana12/Library-Management/model"
	"github.com/gin-gonic/gin"
)

var bookCache *cache.RedisCache

func Initialize() {
	bookCache = cache.NewRedisCache("localhost:6379", 0, 100)
}
func AddBook(c *gin.Context) {
	var book models.Books
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := config.DB.Create(&book).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, &book)
}

func ShowBooks(c *gin.Context) {
	var books []models.Books
	cachedResult := bookCache.GetBookArray("allBooks")
	if cachedResult == nil {
		result := config.DB.Find(&books).Error
		if result != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error()})
			return
		}
		bookCache.SetBookByArray("allBooks", books)
		c.JSON(200, books)
	} else {
		c.JSON(200, cachedResult)
	}

}

func GetBookById(c *gin.Context) {
	id := c.Param("id")
	var result *models.Books = bookCache.GetBook(id)
	if result == nil {
		err := config.DB.Where("book_id = ?", id).First(&result)
		if err.RowsAffected == 0 {
			c.String(200, "No Record Found with this id")
			return
		}
		bookCache.SetBook(id, result)
		c.JSON(200, result)
		return
	} else {
		c.JSON(200, result)
	}
}

func GetBookByGenre(c *gin.Context) {
	genre := c.Param("genre")
	var result []models.Books
	err := config.DB.Where("book_genre = ?", genre).Find(&result)
	if err.RowsAffected == 0 {
		c.String(200, "No Record Found with this genre")
		return
	}
	c.JSON(200, &result)
}
