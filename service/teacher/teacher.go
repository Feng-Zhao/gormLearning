package teacher

import (
	"booklibrary/model"
	"gorm.io/gorm"
)

type IService interface {
	ListTeacher(db *gorm.DB) []model.Teacher
}
