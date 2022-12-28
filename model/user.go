package models

type Users struct {
	StudentId   int    `json:"studentId" gorm:"primaryKey"`
	StudentName string `json:"studentName" binding:"required"`
	EmailId     string `json:"emailId" binding:"required,email" gorm:"unique"`
}
