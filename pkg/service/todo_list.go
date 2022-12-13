package service

import (
	todo "github.com/OlegKapat/ListOfWorks"
	"github.com/OlegKapat/ListOfWorks/pkg/repository"
)

type TodoListService struct{
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList ) *TodoListService{
	return &TodoListService{repo:repo}
}
func (s *TodoListService) Create(userId int,list todo.TodoList)(int,error){
	return s.repo.Create(userId,list)
}