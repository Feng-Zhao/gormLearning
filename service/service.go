package service

import (
	"booklibrary/service/student"
	studentImp "booklibrary/service/student/imp"
)

type Service struct {
	APIS           map[string]API
	StudentService student.IService
}

type API struct {
	Name string
	Func interface{}
}

func (o *Service) InitApi() {
	o.APIS = make(map[string]API)
	o.APIS["listStudent"] = API{Name: "listStudent", Func: o.StudentService.ListStudent}
	o.APIS["addStudent"] = API{Name: "addStudent", Func: o.StudentService.AddStudent}
}

func NewService() *Service {
	s := Service{}
	s.StudentService = studentImp.NewService()
	return &s
}
