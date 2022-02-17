package student

import (
	"booklibrary/model"

	"gorm.io/gorm"
)

type IService interface {
	ListStudent(db *gorm.DB, args map[string]string) []model.Student
}
