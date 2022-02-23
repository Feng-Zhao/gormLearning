package student

import (
	"booklibrary/model"

	"gorm.io/gorm"
)

type IService interface {
	ListStudent(db *gorm.DB) []model.Student
	AddStudent(db *gorm.DB, id uint, name string, cid uint, cname string, tid uint, tname string) (model.Student, error)
}
