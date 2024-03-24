package handler

import (
	"VK-Quest/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	actorsGroup := router.Group("/users")
	{
		actorsGroup.POST("/", h.createUser)
		actorsGroup.GET("/:id", h.getUserAchievements)
	}

	moviesGroup := router.Group("/quests")
	{
		moviesGroup.POST("/", h.createQuest)
		moviesGroup.POST("/:id", h.setCompleteQuest)
	}

	return router
}
