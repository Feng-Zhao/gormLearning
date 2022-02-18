package imp

import (
	"booklibrary/service/teacher"
	"log"
)

type Service struct{}

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmsgprefix)
}
func NewService() teacher.IService {
	s := &Service{}
	return s
}
