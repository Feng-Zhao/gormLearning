package service

import (
	"booklibrary/service/student"
	"booklibrary/service/student/imp"
	"fmt"
)

type Service struct {
	StudentService *student.IService
}

type API struct {
	Name string
	Func interface{}
}

var APIS map[string]API

func init() {
	APIS = make(map[string]API)
}

func (o *Service) InitApi() {
	APIS["listStudent"] = API{Name: "listStudent", Func: fmt.Println}
}

func NewService() *Service {
	s := Service{}
	s.StudentService = imp.NewService()
	return &s
}
