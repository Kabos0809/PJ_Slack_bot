package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type ModelInterface interface {
	GetRestDatebyID(id uuid.UUID) (*RestDate, error)
	GetRestDateFromStudent(studentID uuid.UUID) (*[]RestDate, error)
	CreateRestDate(rdate *RestDate) error
	DeleteRestDate(id uuid.UUID) error
	GetAllSchool() (*[]School, error)
	GetSchoolbyID(id uuid.UUID) (*School, error)
	AddSchool(school *School) error
	DeleteSchool(id uuid.UUID) error
	TestGetFirstSchool(id uuid.UUID) (*School, error)
	GetStudentbySchoolAndGrade(school string, grade string) (*[]Student, error)
	GetStudentbyID(id uuid.UUID) (*Student, error)
	CreateStudent(student *Student) error
	DeleteStudent(id uuid.UUID) error
	AddRestDate4Student(rdate *RestDate, id uuid.UUID) error
	DeleteRestFromStudent(rdate *RestDate, id uuid.UUID) error
	TransferCount(id uuid.UUID) (TransferCounts, error)
}

type Model struct {
	Db *gorm.DB
}