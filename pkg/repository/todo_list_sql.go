package repository

import (
	"crypto/internal/nistec"
	"fmt"

	todo "github.com/OlegKapat/ListOfWorks"
	"gorm.io/gorm"
)

type TodoListSql struct {
	db *gorm.DB
}
func NewTodoListSql(db *gorm.DB)*TodoListSql{
	return &TodoListSql{db: db}
}

func (r *TodoListSql) Create(userId int,list todo.TodoList)(int,error){
	tx,err:=r.db.Begin()
	if err!=nil{
		return 0,err
	}
	var id int
	createListQuery:=fmt.Sprintf("INSERT INTO %s (title,description) VALUES ($1,$2) RETURNING id",todoListsTable)
	row := tx.QueryRow(createListQuery,list.Title,list.Description)
	if err:= row.Scan(&id);
	   err!=nil{
			tx.Rollback()
			return 0,err
	}
	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id,list_id) VALUES ($1,$2)",userListsTable)
	_,err = tx.Exec(createUsersListQuery,userId,id)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}
	return id,tx.Commit()
}