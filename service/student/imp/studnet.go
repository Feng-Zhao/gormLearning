package imp

import (
	"booklibrary/model"
	"fmt"
	"log"
	"reflect"

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

func (s *Service) AddStudent(db *gorm.DB, id uint, name string, cid uint, cname string, tid uint, tname string) (model.Student, error) {
	if addArgCheck(id, cid, tid) {
		log.Println("wrong type of arg")
		return model.Student{}, fmt.Errorf("wrong type of args")
	}
	student := model.Student{ID: id, Name: name, ClassID: cid, ClassName: cname, TID: tid, TName: tname}
	return student, nil
}

func addArgCheck(id uint, cid uint, tid uint) bool {
	if reflect.TypeOf(id).Kind() == reflect.Uint &&
		reflect.TypeOf(cid).Kind() == reflect.Uint &&
		reflect.TypeOf(cid).Kind() == reflect.Uint {
		return true
	}
	return false
}
