package Models

type School struct {
	ID uint64 `gorm:"primaryKey; AUTO_INCREMENT;"`
	Students []Student `gorm:"foreignKey:SchoolID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name string `gorm:"not null;"`
}

func (s *School) TableName() string {
	return "school"
}