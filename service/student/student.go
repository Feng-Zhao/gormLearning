package student

import (
	"booklibrary/model"

	"gorm.io/gorm"
)

type IService interface {
	ListStudent(db *gorm.DB) []model.Student
	AddStudent(db *gorm.DB, student *model.Student) error
}

// type Student struct {
// 	ID          int64
// 	Name        string
// 	ClassID     int64
// 	ClassName   string
// 	TeacherID   int64
// 	TeacherName string
// }
