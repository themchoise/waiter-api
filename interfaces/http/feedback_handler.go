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

// Create godoc
// @Summary      Registrar feedback
// @Description  Registra el feedback de un cliente sobre el servicio
// @Tags         feedback
// @Accept       json
// @Produce      json
// @Param        request  body      usecase.CreateFeedbackInput  true  "Datos del feedback"
// @Success      201      {object}  entity.Feedback
// @Failure      400      {object}  map[string]string
// @Failure      422      {object}  map[string]string
// @Router       /api/v1/feedback [post]
func (h *FeedbackHandler) Create(c *gin.Context) {
	var input usecase.CreateFeedbackInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation_error", "message": err.Error()})
		return
	}

	fb, err := h.uc.CreateFeedback(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable_entity", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fb)
}
