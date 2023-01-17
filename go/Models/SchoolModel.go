package Models

import "github.com/google/uuid"

type School struct {
	ID uuid.UUID `gorm:"primaryKey; type:uuid;"`
	Students []Student `gorm:"foreignKey:SchoolID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name string `gorm:"not null;"`
}

func (s *School) TableName() string {
	return "school"
}