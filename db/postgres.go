package db

import (
	"fmt"
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
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		"localhost",
		"5432",
		"postgres",
		"1111",
		"ToDoDB",
	)
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
