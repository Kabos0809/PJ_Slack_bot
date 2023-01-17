package Models

import (
	"errors"
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
	if err := tx.Where("id = ?", id).First(restdate).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return restdate, nil
}

//休んだ日の登録
func (m Model) CreateRestDate(rdate *RestDate) error {
	tx := m.Db.Begin()
	if rdate.Subject != "国語" && rdate.Subject != "数学" && rdate.Subject != "英語" {
		err := errors.New("[ERROR] Subjects must be Japanese, Mathematics or English.")
		return err
	}
	if err := tx.Create(rdate).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//休んだ日の削除
func (m Model) DeleteRestDate(id uuid.UUID) error {
	var restdate *RestDate
	tx := m.Db.Begin()

	if err := tx.Where("id = ?", id).First(restdate).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("id = ?", id).Delete(&RestDate{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}