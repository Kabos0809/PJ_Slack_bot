package Models

import (
	"github.com/google/uuid"
)

//以下二つはWebページに対応した場合に使うかも
/*
func (m Model) GetAllRestDate() (*[]RestDate, error) {
	var restdates []RestDate
	tx := m.Db.Begin()
	if err := tx.Find(&restdates).Error; err != nil {
		tx.Rollback()
		return &restdates, err
	}
	tx.Commit()
	return &restdates, nil
}
*/

func (m Model) GetRestDatebyID(id uuid.UUID) (*RestDate, error) {
	var restdate *RestDate
	tx := m.Db.Begin()
	if err := tx.Where("id = ?", id).First(&restdate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return restdate, nil
}

func (m Model) GetRestDateFromStudent(studentID uuid.UUID) (*[]RestDate, error) {
	var restdates []RestDate
	var student *Student
	tx := m.Db.Preload("RestDates").Begin()
	if err := tx.Where("id = ?", studentID).First(&student).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	for _, r := range student.RestDates {
		restdates = append(restdates, r)
	}

	return &restdates, nil
}

//休んだ日の登録
func (m Model) CreateRestDate(rdate *RestDate) error {
	tx := m.Db.Begin()
	rdate.ID = uuid.New()
	if err := tx.Create(rdate).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

//休んだ日の削除
func (m Model) DeleteRestDate(id uuid.UUID) error {
	tx := m.Db.Begin()

	if err := tx.Where("id = ?", id).Delete(&RestDate{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}