package test

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"github.com/kabos0809/slack_bot/go/Models"
	"github.com/kabos0809/slack_bot/go/Config"
)

func TestConnDB(t *testing.T) {
	godotenv.Load("../../.env")
	
	dsn := Config.DbUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatalf("[FATAL] Failed to Open DB: %s", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	err = db.AutoMigrate(&Models.RestDate{}, &Models.Student{}, &Models.School{})
	if err != nil {
		t.Fatalf("[FATAL] Failed to Migrate: %s", err)
	}
}