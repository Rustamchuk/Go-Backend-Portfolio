package handler

import (
	"VK-Test_Ex/internal/service"
	"errors"
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	actorsGroup := router.Group("/actors")
	{
		actorsGroup.POST("/", h.createActor)
		actorsGroup.PUT("/:id", h.updateActor)
		actorsGroup.DELETE("/:id", h.deleteActor)
		actorsGroup.GET("/", h.listActors)
	}

	moviesGroup := router.Group("/movies")
	{
		moviesGroup.POST("/", h.createMovie)
		moviesGroup.PUT("/:id", h.updateMovie)
		moviesGroup.DELETE("/:id", h.deleteMovie)
		moviesGroup.GET("/", h.listMovies)
		moviesGroup.GET("/search", h.searchMovies)
	}

	return router
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
