package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ModelInterface interface {
	GetRestDatebyID(id uint64) (*RestDate, error)
	CreateRestDate(rdate *RestDate) error
	DeleteRestDate(id uint64) error
	GetAllSchool() (*[]School, error)
	GetSchoolbyID(id uint64) (*School, error)
	AddSchool(school *School) error
	DeleteSchool(id uint64) error
	TestGetFirstSchool(id uuid.UUID) (*School, error)
	GetStudentbySchoolAndGrade(school string, grade string) (*[]Student, error)
	GetStudentbyID(id uint64) (*Student, error)
	CreateStudent(student *Student) error
	DeleteStudent(id uint64) error
	AddRestDate4Student(rdate *RestDate, id uint64) error
	DeleteRestFromStudent(rdate *RestDate, id uint64) error
}

type Model struct {
	Db *gorm.DB
}