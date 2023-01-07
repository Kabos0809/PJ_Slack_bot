package Models

import "github.com/google/uuid"

type School struct {
	SchoolID uuid.UUID `gorm:"primaryKey; type:uuid; default:uuid_generate_v4();"`
	Students []Student `gorm:"foreignKey: SchoolID references: SchoolID default: nil"`
	Name string `gorm:"not null;"`
}

func (s *School) TableName() string {
	return "school"
}