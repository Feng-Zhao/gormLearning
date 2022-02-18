package imp

import (
	"booklibrary/model"
	"gorm.io/gorm"
)

func (o *Service) ListTeacher(db *gorm.DB) []model.Teacher {
	var teachers []model.Teacher
	db.Find(&teachers)
	return teachers
}
