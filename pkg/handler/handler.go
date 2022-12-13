package handler

import (
	"github.com/OlegKapat/ListOfWorks/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth:= router.Group("/auth")
	{
		auth.POST("/singUp",h.singUp)
		auth.POST("/singIn",h.singIn)
	}
	api:= router.Group("/api",h.userIdentity)
	{
		lists:= api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/",h.getAllLists)
			lists.GET("/:id",h.getListById)
			lists.DELETE("/:id",h.deleteListById)
			lists.PUT("/:id",h.updateListById)
		}
		items:= lists.Group(":id/items")
		{
			items.POST("/",h.createItem)
			items.GET("/",h.getAllItems)
			items.GET("/:item_id",h.getItemById)
			items.PUT("/:item_id",h.updateItemById)
			items.DELETE("/d:item_id",h.deleteItemById)
			

		}
	}
	return router
}
