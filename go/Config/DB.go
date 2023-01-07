package Config

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

type DbConfig struct {
	User string
	DB string
	Pass string
	Host string
	TZ string
	Port string
}

func BuildDBConfig() *DbConfig {
	dbconfig := DbConfig{
		User: os.Getenv("POSTGRES_USER"),
		DB: os.Getenv("POSTGRES_DB"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		Host: os.Getenv("DB_HOST"),
		TZ: os.Getenv("TZ"),
		Port: "5432",
	}
	return &dbconfig
}

func DbUrl() string {
	dbConfig := BuildDBConfig()
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.DB,
		dbConfig.Port,
		dbConfig.TZ,
	)
}
