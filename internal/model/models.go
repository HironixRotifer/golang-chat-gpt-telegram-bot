package models

import (
	database "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/database"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&database.UserStorage{})
	return err
}
