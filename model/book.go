package model

import (
	"time"
)

type Book struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	BelongID    uint
	Description string
	CreatedTime time.Time `gorm:"autoCreateTime"`
	UpdatedTime time.Time `gorm:"autoUpdateTime"`
}

func GetFixedBook(id uint) *Book {
	return &Book{ID: 0002, Name: "new book", BelongID: 2022010101, Description: "new book"}
}
