package main

import (
	"todo/db"
	_ "todo/cmd/docs" 
	"todo/internal/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           ToDo API
// @version         1.0
// @description     Простой ToDo List на Gin + PostgreSQL
// @host            localhost:8080
// @BasePath        /

func main() {
	db := db.InitPostgres()
	h := handlers.NewHandler(db)
	r := gin.Default()
	
	r.GET("/todos", h.GetTodos)
	r.GET("/todo/:id", h.GetTodo)
	r.POST("/todos", h.CreateTask)
	r.PATCH("/todos/:id", h.UpdateTodo)
	r.DELETE("/todos/:id", h.DelateTask)
	//Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	r.Run(":8080")
}
