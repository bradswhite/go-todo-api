package dal

import (
	"go-todo-api/database"

	"gorm.io/gorm"
)

// Todo struct defines the Todo Model
type Todo struct {
	gorm.Model
	Task      string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	User      uint  `gorm:"not null" gorm:"index"`
	// this is a pointer because int == 0,
}

// CreateTodo create a todo entry in the todo's table
func CreateTodo(todo *Todo) *gorm.DB {
	return database.DB.Db.Create(todo)
}

// FindTodo finds a todo with given condition
func FindTodo(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Db.Model(&Todo{}).Take(dest, conds...)
}

// FindTodoByUser finds a todo with given todo and user identifier
func FindTodoByUser(dest interface{}, todoIden interface{}, userIden interface{}) *gorm.DB {
	return FindTodo(dest, "id = ? AND todos.user = ?", todoIden, userIden)
}

// FindTodosByUser finds the todos with user's identifier given
func FindTodosByUser(dest interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Db.Model(&Todo{}).Find(dest, "todos.user = ?", userIden)
}

// DeleteTodo deletes a todo from todos' table with the given todo and user identifier
func DeleteTodo(todoIden interface{}, userIden interface{}) *gorm.DB {
	return database.DB.Db.Unscoped().Delete(&Todo{}, "id = ? AND todos.user = ?", todoIden, userIden)
}

// UpdateTodo allows to update the todo with the given todoID and userID
func UpdateTodo(todoIden interface{}, userIden interface{}, data interface{}) *gorm.DB {
	return database.DB.Db.Model(&Todo{}).Where("id = ? AND todos.user = ?", todoIden, userIden).Save(data)
}
