package test

import (
	"time"
	"testing"
	
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/kabos0809/slack_bot/go/Models"
	"github.com/kabos0809/slack_bot/go/Config"
)

var format string = "2006-01-02"
var testdate string = "2023-01-17"

func TestCreate(t *testing.T) {
	godotenv.Load("../../.env")
	
	dsn := Config.DbUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	
	err = db.AutoMigrate(&Models.RestDate{}, &Models.Student{}, &Models.School{})
	if err != nil {
		panic(err)
	}

	m := Models.Model{Db: db}

	testSchool := Models.School{
		Name: "TestSchool1",
	}
	if err := m.AddSchool(&testSchool); err != nil {
		t.Fatalf("[FATAL] Failed to AddSchool: %s", err)
	}

	school, err := m.TestGetFirstSchool()
	if err != nil {
		t.Fatalf("[FATAL] Failed to Get School: %s", err)
	}

	testStudent := Models.Student{
		ID: uuid.New(),
		FirstName: "hoge",
		LastName: "henyo",
		Name: "hogehenyo",
		Grade: "テスト",
		SchoolID: school.ID,
	}
	if err := m.CreateStudent(&testStudent); err != nil {
		t.Fatalf("[FATAL] Failed to CreateStudent: %s", err)
	}

	if err := m.AddStudent4School(&testStudent, testStudent.SchoolID); err != nil {
		t.Fatalf("[FATAL] Failed to Add Student for School: %s", err)
	}

	date, _ := time.Parse(testdate, format)

	testRestDate := Models.RestDate{
		ID: uuid.New(),
		StudentID: testStudent.ID,
		Date: date,
		Subject: "英語",
	}

	if err := m.CreateRestDate(&testRestDate); err != nil {
		t.Fatalf("[FATAL] Failed to CreateRestDate: %s", err)
	}

	if err := m.AddRestDate4Student(&testRestDate, testRestDate.StudentID); err != nil {
		t.Fatalf("[FATAL] Failed to Add RestDate for Student: %s", err)
	}
}