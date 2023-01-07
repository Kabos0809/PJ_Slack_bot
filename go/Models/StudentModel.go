package Models

type Student struct {
	ID uint64 `gorm:"AUTO_INCREMENT;"`
	FirstName string `gorm:"not null;" binding:"required"`
	LastName string `gorm:"not null;" binding:"required"`
	Name string `gorm:"not null;" binding:"required"`
	Grade string `gorm:"not null;" binding:"required"`
	//振替予定のリストはリリース後拡張予定
	//Schedules []TransferSchedules `gorm:"foreignKey: ID"`
	MathCounts uint64 `gorm:"not null; default: 0;"`
	JpnCounts uint64 `gorm:"not null; default: 0;"`
	EngCounts uint64 `gorm:"not null; default: 0;"`
}

type TransferCounts struct {
	MathCounts uint64
	JpnCounts uint64
	EngCounts uint64
}

func (s *Student) TableName() string {
	return "student"
}