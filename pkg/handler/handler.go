package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/romon267/go-rest/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Any("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PATCH("/:id", h.updateListById)
			lists.DELETE("/:id", h.deleteListById)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PATCH("/:item_id", h.updateItemById)
				items.DELETE("/:item_id", h.deleteItemById)
			}
		}
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"statusCode": 404, "message": "Page not found"})
	})

	return router
}
