package student

import (
	"booklibrary/model"

	"gorm.io/gorm"
)

type IService interface {
	ListStudent(db *gorm.DB) []model.Student
}
