package handler

import (
	"VK-Quest/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary createUser
// @Tags create
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body model.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users/ [post]
func (h *Handler) createUser(c *gin.Context) {
	var input *model.User
	if err := c.BindJSON(&input); err != nil {
		return
	}

	err := h.services.UserService.CreateUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})
}

func (h *Handler) getUserAchievements(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	achievement, err := h.services.UserService.GetAchievements(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, achievement)
}
