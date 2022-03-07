package imp

import (
	"booklibrary/model"
	"fmt"

	"gorm.io/gorm"
)

// ListStudent select * from student
func (s *Service) ListStudent(db *gorm.DB) []model.Student {
	var students []model.Student
	db.Find(&students)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil
	}
	for _, v := range students {
		// log.Printf("student ID : %v |student name : %s\t|class_id : %v |class_name : %v |teacher_id : %v teacher_name : %v\t\n",
		// 	v.ID, v.Name, v.ClassID, v.ClassName, v.TID, v.TName)

		fmt.Printf("student ID : %v |student name : %s\t|class_id : %v |class_name : %v |teacher_id : %v |teacher_name : %v\t\n",
			v.ID, v.Name, v.ClassID, v.ClassName, v.TID, v.TName)
	}
	return students
}

// insert
func (s *Service) AddStudent(db *gorm.DB, student *model.Student) error {
	if err := db.Create(student).Error; err != nil {
		return err
	}
	return nil
}

// func addArgCheck(id uint, cid uint, tid uint) bool {
// 	if reflect.TypeOf(id).Kind() == reflect.Uint &&
// 		reflect.TypeOf(cid).Kind() == reflect.Uint &&
// 		reflect.TypeOf(cid).Kind() == reflect.Uint {
// 		return true
// 	}
// 	return false
// }
