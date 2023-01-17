package Models

func (m Model) GetAllSchool() (*[]School, error) {
	var schools []School
	tx := m.Db.Preload("Students").Begin()
	if err := tx.Find(&schools).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &schools, nil
}

func (m Model) GetSchoolbyID(id uint64) (*School, error) {
	var school *School
	tx := m.Db.Preload("Students").Begin()
	if err := tx.Where("id = ?", id).Find(school).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return school, nil
}

func (m Model) AddSchool(school *School) error {
	tx := m.Db.Begin()
	if err := tx.Create(school).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) DeleteSchool(id uint64) error {
	tx := m.Db.Preload("Students").Begin()
	if err := tx.Where("id = ?", id).Delete(&School{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) AddStudent4School(student *Student, id uint64) error {
	var school *School
	tx := m.Db.Begin()

	if err := tx.Where("id = ?", id).First(school).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := m.Db.Model(school).Association("Students").Append(student); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(school).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}