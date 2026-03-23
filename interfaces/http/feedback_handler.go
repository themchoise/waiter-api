package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waiter/back/application/usecase"
)

type FeedbackHandler struct {
	uc *usecase.FeedbackUseCase
}

func NewFeedbackHandler(uc *usecase.FeedbackUseCase) *FeedbackHandler {
	return &FeedbackHandler{uc: uc}
}

func (h *FeedbackHandler) Create(c *gin.Context) {
	var input usecase.CreateFeedbackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fb, err := h.uc.CreateFeedback(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fb)
}
