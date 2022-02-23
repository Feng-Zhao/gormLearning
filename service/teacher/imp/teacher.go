package imp

import (
	"booklibrary/model"
	"fmt"
	"gorm.io/gorm"
)

func (o *Service) ListTeacher(db *gorm.DB) []model.Teacher {
	var teachers []model.Teacher
	db.Find(&teachers)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil
	}
	for _, v := range teachers {
		fmt.Printf("techer ID : %v |teacger name : %s\t\n",
			v.ID, v.Name)
	}
	return teachers
}
