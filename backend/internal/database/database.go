package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"backend/internal/models"
)

func InitDB() *gorm.DB {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "mydb")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.RoomParticipant{},
		&models.Pertanyaan{},
		&models.QuestionOption{},
		&models.SesiUjian{},
		&models.Answer{},
		&models.HasilUjian{},
		&models.Feedback{},
		&models.ActivityLog{},
		&models.PasswordReset{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}