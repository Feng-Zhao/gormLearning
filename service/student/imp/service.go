package imp

import (
	"booklibrary/service/student"
	"log"
)

type Service struct{}

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
}
func NewService() student.IService {
	s := &Service{}
	return s
}
