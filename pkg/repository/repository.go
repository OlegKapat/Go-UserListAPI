package repository

import (
	todo "github.com/OlegKapat/ListOfWorks"
	"gorm.io/gorm"
)

type Authorization  interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username,password string) (todo.User,error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList)(int,error)
}

type TodoItem interface {}

type Repository struct {
	Authorization 
	TodoList
	TodoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization : NewAuthSql(db),
		TodoList: NewTodoListSql(db),
	}
}