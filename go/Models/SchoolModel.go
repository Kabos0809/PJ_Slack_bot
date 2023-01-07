package Models

type School struct {
	ID uint64 `gorm:"AUTO_INCREMENT;"`
	Name string `gorm:"not null;"`
}

func (s *School) TableName() string {
	return "school"
}