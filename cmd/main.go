package main

import (
	"todo/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB = gorm.DB{}

func main() {
	h := handlers.NewHandler(&gorm.DB{})
	r := gin.Default()

	r.GET("/todo", h.GetTodo)
	r.GET("/todos", h.GetTodos)
	r.POST("/todos", h.CreateTask)
	r.DELETE("/todos:id", h.DelateTask)

	r.Run(":8080")
}
