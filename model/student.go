package model

type Student struct {
	ID        uint `gorm:"column:id"`
	Name      string
	ClassID   uint
	ClassName string
	TID       uint   `gorm:"column:teacher_id"`
	TName     string `gorm:"column:teacher_name"`
}

func (Student) TableName() string {
	return "student"
}
