package handler

import (
	"net/http"

	 "github.com/OlegKapat/ListOfWorks"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	id,ok:=c.Get(userCtx)
	if!ok{
		newErrorResponse(c,http.StatusInternalServerError,"id not found")
		return
	}
	var input todo.TodoList
	if err:= c.BindJSON(&input); 
	err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	// call service method
	h.services.TodoList.Create(id.(int),input)

}
func (h *Handler) getAllLists(c *gin.Context) {}
func (h *Handler) getListById(c *gin.Context) {}
func (h *Handler) updateListById(c *gin.Context) {}
func (h *Handler) deleteListById(c *gin.Context) {}