package Models

import "github.com/google/uuid"

type Student struct {
	ID uuid.UUID `gorm:"primaryKey; type:uuid;"`
	FirstName string `gorm:"not null;" binding:"required"`
	LastName string `gorm:"not null;" binding:"required"`
	Name string `gorm:"not null;" binding:"required"`
	Grade string `gorm:"not null;" binding:"required"`
	SchoolID uuid.UUID
	RestDates []RestDate `gorm:"foreignKey:StudentID;references:ID; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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