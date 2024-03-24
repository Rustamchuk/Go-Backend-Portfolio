package handler

import (
	"VK-Quest/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createQuest(c *gin.Context) {
	var input *model.Quest
	if err := c.BindJSON(&input); err != nil {
		return
	}

	err := h.services.QuestService.CreateQuest(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": input.Id,
	})
}

func (h *Handler) setCompleteQuest(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		return
	}
	questId, err := strconv.Atoi(c.Param("quest_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quest id param")
		return
	}

	quest, err := h.services.QuestService.GetQuest(c, questId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid quest id param")
		return
	}

	err = h.services.UserService.IncreaseUserBalance(c, userId, quest.Cost)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.CompletedQuestService.SetQuestComplete(c, &model.CompletedQuest{
		UserID:  userId,
		QuestID: questId,
	})
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
