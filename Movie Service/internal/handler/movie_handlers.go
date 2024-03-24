package handler

import (
	"VK-Test_Ex/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) != "admin" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	var input *model.Movie
	if err = c.BindJSON(&input); err != nil {
		return
	}

	err = h.services.MovieService.CreateMovie(c, input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.ID,
	})
}

func (h *Handler) updateMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) != "admin" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input *model.UpdateMovie
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	upMovie := &model.Movie{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		ReleaseDate: input.ReleaseDate,
		Rating:      input.Rating,
		Actors:      input.Actors,
	}

	if err := h.services.MovieService.UpdateMovie(c, upMovie); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteMovie(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) != "admin" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.MovieService.DeleteMovie(c, int64(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) listMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) == "" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	sort := c.DefaultQuery("sort", "rating_desc")
	movies, err := h.services.MovieService.GetMovies(c, sort)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (h *Handler) searchMovies(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) == "" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	search := c.Param("id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	movies, err := h.services.MovieService.SearchMovies(c, search)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, movies)
}
