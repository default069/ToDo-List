package db
// Это для чегота не пон но нужная веш 
import "gorm.io/gorm"

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB)handler {
	return handler{db}
}