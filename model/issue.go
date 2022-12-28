package models

type BookIssue struct {
	UserId int
	BookId int
	// Users  Users `gorm:"foreignKey:UserId;references:UserId"`
	// Books  Books `gorm:"foreignKey:BookId;references:BookId"`
}
