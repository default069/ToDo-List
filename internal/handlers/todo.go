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

// GetTodos godoc
// @Summary      Список всех задач
// @Tags         todos
// @Produce      json
// @Success      200  {array}   models.Todo
// @Router       /todos [get]
func (h *Handler) GetTodos(c *gin.Context){
	var todos []models.Todo
	h.DB.Find(&todos)
	c.JSON(http.StatusOK,todos)
}
//Get метой что бы п sолучить толка одно по ид --> /todos/:id

// GetTodo godoc
// @Summary      Получить задачу по ID
// @Tags         todos
// @Produce      json
// @Param        id   path      int  true  "ID задачи"
// @Success      200  {object}  models.Todo
// @Failure      404  {object}  map[string]string
// @Router       /todos/{id} [get]
func (h *Handler) GetTodo(c *gin.Context){
	var todo []models.Todo
	if err := h.DB.First(&todo, c.Param("id")).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":"task not faund"})
		return 
	}
	c.JSON(http.StatusOK, todo)
}
//POST что бы добавить task на тодо лист -->  /todos

// CreateTask godoc
// @Summary      Создать задачу
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo  body      models.Todo  true  "Новая задача"
// @Success      200   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /todos [post]
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

// DelateTask godoc
// @Summary      Удалить задачу по ID
// @Tags         todos
// @Param        id   path  int  true  "ID задачи"
// @Success      204
// @Router       /todos/{id} [delete]
func (h *Handler) DelateTask(c *gin.Context){
	h.DB.Delete(&models.Todo{}, c.Param("id"))
	c.Status(http.StatusNoContent)
}

// UpdateTodo godoc
// @Summary      Обновить задачу (частично)
// @Description  Меняются только те поля, которые переданы в теле запроса
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id    path      int                     true  "ID задачи"
// @Param        todo  body      models.UpdateTodoInput  true  "Поля для обновления"
// @Success      200   {object}  models.Todo
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /todos/{id} [patch]
func (h *Handler) UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := h.DB.First(&todo, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	var input models.UpdateTodoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != nil {
		todo.Titles = *input.Title
	}
	if input.Description != nil {
		todo.Description = *input.Description
	}
	if input.Completed != nil {
		todo.Completed = *input.Completed
	}

	if err := h.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update task"})
		return
	}
	c.JSON(http.StatusOK, todo)
}