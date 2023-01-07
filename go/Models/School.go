package Models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

func GetAllSchool(db *gorm.DB) (*[]School, error) {
	var schools []School
	tx := db.Preload("Students").Begin()
	if err := tx.Find(&schools).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &schools, nil
}

func GetSchool(id uuid.UUID, db *gorm.DB) (*School, error) {
	var school *School
	tx := db.Preload("Students").Begin()
	if err := tx.Where("school_id = ?", id).Find(school).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return school, nil
}

func AddSchool(school *School, db *gorm.DB) error {
	tx := db.Begin()
	if err := tx.Create(school).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteSchool(id uuid.UUID, db *gorm.DB) error {
	tx := db.Preload("Students").Begin()
	if err := tx.Where("school_id = ?", id).Delete(&School{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetStudentsFromSchool(id uuid.UUID, db *gorm.DB) ([]Student, error) {
	var school *School
	tx := db.Preload("Students").Begin()
	if err := tx.Find(school).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return school.Students, nil
}