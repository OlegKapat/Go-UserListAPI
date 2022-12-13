package service

import (
	"github.com/OlegKapat/ListOfWorks"
	"github.com/OlegKapat/ListOfWorks/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenereteToken(username,password string) (string,error)
	ParseToken(token string)(int,error)
}

type TodoList interface  {
	Create(userId int, list todo.TodoList)(int,error)
}

type TodoItem interface{}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	    TodoList:      NewTodoListService(repos.TodoList),
		// TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
