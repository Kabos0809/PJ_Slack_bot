package test

import (
	"fmt"
	"testing"
)

/*type metadata struct {
	schid uuid.UUID
	stuid uuid.UUID
	resid uuid.UUID
}*/


func TestCrateStudent(t *testing.T) {
	if err := m.AddSchool(&testSchool_1); err != nil {
	t.Fatalf("[FATAL] Failed to Add School: %s", err)
	}	
	if err := m.AddSchool(&testSchool_2); err != nil {
	t.Fatalf("[FATAL] Failed to Add School: %s", err)
	}	
	for _, s := range testStudents {
		if err := m.CreateStudent(&s); err != nil {
			t.Fatalf("[FATAL] Failed to Create Student: %s", err)
		}
		if err := m.AddStudent4School(&s, s.SchoolID); err != nil {
			t.Fatalf("[FATAL] Failed to Add Student for Schools Association: %s", err)
		}
	}
}

func TestGetStudent(t *testing.T) {
	students_1, err := m.GetStudentbySchoolAndGrade(testSchool_1.ID, "テスト")
	if err != nil {
		t.Fatalf("[FATAL] Failed to Get Student From School: %s", err)
	}
	students_2, err := m.GetStudentbySchoolAndGrade(testSchool_2.ID, "はずれ")
	if err != nil {
		t.Fatalf("[FATAL] Failed to Get Student From School: %s", err)
	}

	fmt.Println(*students_1)
	fmt.Println(*students_2)
}

func TestDeleteStudent(t *testing.T) {
	defer sqldb.Close()

	for _, s := range testStudents {
		if err := m.DeleteStudentFromSchool(&s, s.SchoolID); err != nil {
			t.Fatalf("[FATAL] Failed to Delete Student from School: %s", err)
		}
		if err := m.DeleteStudent(s.ID); err != nil {
			t.Fatalf("[FATAL] Failed to Delete Student: %s", err)
		}
	}
}