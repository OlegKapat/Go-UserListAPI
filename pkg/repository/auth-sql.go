package repository

import (
	"github.com/OlegKapat/ListOfWorks"

	"gorm.io/gorm"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	userListsTable  = "user_lists"
	todoItemsTable  = "todo_items"
	ListsItemsTable = "lists_items"
)

type AuthSql struct {
	db *gorm.DB
}

func NewAuthSql(db *gorm.DB) *AuthSql {
	return &AuthSql{db: db}
}

func (r *AuthSql) CreateUser(user todo.User) (int, error) {
	var id int
	row := r.db.Create(&user)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}
func (r *AuthSql) GetUser(username, password string) (todo.User, error) {
	var user todo.User
	query := r.db.Table("users").Where("username=?", username, "password=?", password).Select("id").Scan(&user)
	error := query.Error
	return user, error
}
