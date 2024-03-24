package handler

import (
	"VK-Test_Ex/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createActor(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) != "admin" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	var input *model.Actor
	if err := c.BindJSON(&input); err != nil {
		return
	}

	err = h.services.ActorService.CreateActor(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.ID,
	})
}

func (h *Handler) updateActor(c *gin.Context) {
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

	var input *model.UpdateActor
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	upActor := &model.Actor{
		ID:        id,
		Name:      input.Name,
		Gender:    input.Gender,
		BirthDate: input.BirthDate,
	}

	if err := h.services.ActorService.UpdateActor(c, upActor); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteActor(c *gin.Context) {
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

	err = h.services.ActorService.DeleteActor(c, int64(id))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) listActors(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if h.services.GetRole(userId) == "" {
		newErrorResponse(c, http.StatusNotFound, "permission denied")
		return
	}

	items, err := h.services.ActorService.GetActors(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
