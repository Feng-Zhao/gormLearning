package imp

import (
	"booklibrary/model"
	"fmt"

	"gorm.io/gorm"

	"log"
)

func (s *Service) ListStudent(db *gorm.DB, args map[string]string) []model.Student {
	var studnets []model.Student
	db.Find(&studnets)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil
	}
	fmt.Println(db.Statement)
	for _, v := range studnets {
		// fmt.Printf("student ID : %v\t|student name : %s\t|class_id : %v\t|class_name : %v\t|teacher_id : %v|teacher_name : %v\t\n",
		// 	v.ID, v.Name, v.ClassID, v.ClassName, v.TID, v.TName)

		log.Printf("student ID : %v\t|student name : %s\t|class_id : %v\t|class_name : %v\t|teacher_id : %v|teacher_name : %v\t\n",
			v.ID, v.Name, v.ClassID, v.ClassName, v.TID, v.TName)
	}
	return studnets
}
