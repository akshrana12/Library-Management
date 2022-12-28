package controller

import (
	"net/http"
	"strconv"

	"github.com/akshrana12/Library-Management/config"
	models "github.com/akshrana12/Library-Management/model"
	"github.com/gin-gonic/gin"
)

func IssueBook(c *gin.Context) {
	userId := c.Query("userId")
	bookId := c.Query("bookId")
	var result models.Books
	err := config.DB.Where("book_id = ?", bookId).First(&result)
	if err.RowsAffected == 0 {
		c.String(400, "Not valid book or user id")
		return
	}
	var result1 models.Users
	err1 := config.DB.Where("student_id = ?", userId).First(&result1)
	if err1.RowsAffected == 0 {
		c.String(400, "Not valid book or user id")
		return
	}
	var finalResult models.BookIssue
	err2 := config.DB.Where(" user_id = ? AND book_id = ?", userId, bookId).First(&finalResult)
	if err2.RowsAffected > 0 {
		c.String(400, "Book is already issued to this user")
		return
	}
	validUserId, _ := strconv.Atoi(userId)
	validBookId, _ := strconv.Atoi(bookId)
	var issueThisBook models.BookIssue = models.BookIssue{
		UserId: validUserId,
		BookId: validBookId,
	}
	error := config.DB.Create(&issueThisBook).Error
	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	c.JSON(200, &issueThisBook)
}

func ReturnBook(c *gin.Context) {
	userId := c.Query("userId")
	bookId := c.Query("bookId")
	var finalResult models.BookIssue
	err2 := config.DB.Where(" user_id = ? AND book_id = ?", userId, bookId).Delete(&finalResult)
	if err2.RowsAffected == 0 {
		c.String(400, "Following book was not issued to this person")
		return
	}
	c.String(200, "Return of book was Successfull")
}

func GetUserDetail(c *gin.Context) {
	userId := c.Query("userId")
	var finalResult []models.BookIssue
	err := config.DB.Where(" user_id = ?", userId).Find(&finalResult)
	if err.RowsAffected == 0 {
		c.String(400, "No book was issued by this person")
		return
	}
	issuedBooks := make([]models.Books, 0)
	for _, book := range finalResult {
		var res models.Books
		config.DB.Where("book_id=?", book.BookId).First(&res)
		issuedBooks = append(issuedBooks, res)
	}
	c.JSON(200, issuedBooks)
}
