package handlers

import (
	"net/http"
	"todo/models"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)
type Handler struct{
	DB *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler{
	return &Handler{DB: db}
}
//Логоки работи Get Метода чтобы получить все и толка одну по ID
//Get Methor writing --> /todos: список задач
func (h *Handler) GetTodos(c *gin.Context){
	var todos []models.Todo
	h.DB.Find(&todos)
	c.JSON(http.StatusOK,todos)
}
//Get метой что бы п sолучить толка одно по ид --> /todos/:id
func (h *Handler) GetTodo(c *gin.Context){
	var todo []models.Todo
	if err := h.DB.First(&todo, c.Param("id")).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"task not faund"})
		return 
	}
	c.JSON(http.StatusOK, todo)
}
//POST что бы добавить task на тодо лист -->  /todos
func (h *Handler) CreateTask(c * gin.Context){
	var todo []models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"failedcreate new task "})
		return
	}
	c.JSON(http.StatusOK,todo)
}
//DELATE что бы удалить task по ID --> /todos/:id
func (h *Handler)DelateTask(c * gin.Context){
	h.DB.Delete(&models.Todo{}, c.Param("id"))
	c.Status(http.StatusNoContent)
}