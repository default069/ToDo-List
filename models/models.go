package models

// это модел для разроса на ДатаБазу в json формате
type Todo struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Titles      string `json:"title" gorm:"varchar(55)"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// UpdateTodoInput - поля, которые можно изменить (все необязательные)
type UpdateTodoInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Completed   *bool   `json:"completed"`
}