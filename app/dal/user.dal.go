package dal

import (
	"go-todo-api/database"

	"gorm.io/gorm"
)

// User struct defines the user
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Todos    []Todo `gorm:"foreignKey:User"`
}

// CreateUser create a user entry in the user's table
func CreateUser(user *User) *gorm.DB {
	return database.DB.Db.Create(user)
}

// FindUser searches the user's table with the condition given
func FindUser(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Db.Model(&User{}).Take(dest, conds...)
}

// FindUserByEmail searches the user's table with the email given
func FindUserByEmail(dest interface{}, email string) *gorm.DB {
	return FindUser(dest, "email = ?", email)
}
