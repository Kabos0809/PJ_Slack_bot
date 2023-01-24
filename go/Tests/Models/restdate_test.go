package test

import (
	"fmt"
	"testing"
	"database/sql"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/kabos0809/slack_bot/go/Config"
	"github.com/kabos0809/slack_bot/go/Models"
)
/*type metadata struct {
	schid uuid.UUID
	stuid uuid.UUID
	resid uuid.UUID
}*/
var m Models.Model
var sqldb *sql.DB

func TestCreate(t *testing.T) {
	m, sqldb = ConnDB()

	if err := m.AddSchool(&testSchool_1); err != nil {
		t.Fatalf("[FAIL] Failed to Add School: %s", err)
	}
	if err := m.AddSchool(&testSchool_2); err != nil {
		t.Fatalf("[FAIL] Failed to Add School: %s", err)
	}

	for _, testStudent := range testStudents_2 {
		if err := m.CreateStudent(&testStudent); err != nil {
		t.Fatalf("[FAIL] Failed to CreateStudent: %s", err)
		}

		if err := m.AddStudent4School(&testStudent, testStudent.SchoolID); err != nil {
		t.Fatalf("[FAIL] Failed to Add Student for School: %s", err)
		}
	}
	
	for _, testRestDate := range testRestDates {
		if err := m.CreateRestDate(&testRestDate); err != nil {
		t.Fatalf("[FAIL] Failed to CreateRestDate: %s", err)
		}

		if err := m.AddRestDate4Student(&testRestDate, testRestDate.StudentID); err != nil {
		t.Fatalf("[FAIL] Failed to Add RestDate for Student: %s", err)
		}
	}
}

func TestGetRestDate(t *testing.T) {
	restdate, err := m.GetRestDatebyID(testRestDates[2].ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get RestDate by ID: %s", err)
	}
	
	fmt.Println(*restdate)

	restdates_1, err := m.GetRestDateFromStudent(testStudents_2[0].ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get RestDate from Student: %s", err)
	}

	restdates_2, err := m.GetRestDateFromStudent(testStudents_2[2].ID)
	if err != nil {
		t.Fatalf("[FAIL] Failed to Get RestDate from Student: %s", err)
	}

	fmt.Println(*restdates_1)
	fmt.Println(*restdates_2)
}

func TestDeleteRestDate(t *testing.T) {
	for _, restdate := range testRestDates{
		if err := m.DeleteRestFromStudent(&restdate, restdate.StudentID); err != nil {
		t.Fatalf("[FAIL] Failed to Delete RestDate from Student Association: %s", err)
		}

		if err := m.DeleteRestDate(restdate.ID); err != nil {
		t.Fatalf("[FAIL] Failed to Delete RestDate: %s", err)
		}
	}
}

func ConnDB() (Models.Model, *sql.DB) {
	godotenv.Load("../../.env")
	
	dsn := Config.DbUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	
	err = db.AutoMigrate(&Models.RestDate{}, &Models.Student{}, &Models.School{})
	if err != nil {
		panic(err)
	}

	m := Models.Model{Db: db}
	return m, sqlDB
}