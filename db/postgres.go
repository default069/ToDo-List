package db

import (
	"log"
	"todo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Файл для подклучения в дата базе PostgresSql базе
//
// Это для додклучения к Дата База
func InitPostgres() *gorm.DB {
	dns := "postgres://postgres:1111@localhost:5432/todo_db?sslmode=disable"

	//А это чтобы открить форти или Логи я фулл не понел
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//Это для обработки ошибк
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}
	log.Println("connected")
	db.AutoMigrate(&models.Todo{})
	return db
}
