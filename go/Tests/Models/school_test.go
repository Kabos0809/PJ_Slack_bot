package test

import (
	"testing"
)

func TestAddSchool(t *testing.T) {
	if err := m.AddSchool(&testSchool_1); err != nil {
		t.Fatalf("[FAIL] Failed to add school: %s", err)
	}
	if err := m.AddSchool(&testSchool_2); err != nil {
		t.Fatalf("[FAIL] Failed to add school: %s", err)
	}
}

func TestDeleteSchool(t *testing.T) {
	if err := m.DeleteSchool(testSchool_1.ID); err != nil {
		t.Fatalf("[FAIL] Failed to delete school: %s", err)
	}
	if err := m.DeleteSchool(testSchool_2.ID); err != nil {
		t.Fatalf("[FAIL] Failed to delete school: %s", err)
	}
}