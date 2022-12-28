package models

type Books struct {
	BookId     int    `json:"bookId" gorm:"primaryKey"`
	BookName   string `json:"bookName" gorm:"unique" binding:"required"`
	BookAuthor string `json:"bookAuthor" binding:"required"`
	BookGenre  string `json:"bookGenre" binding:"required"`
}
