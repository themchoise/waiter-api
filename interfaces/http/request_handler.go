package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waiter/back/application/usecase"
)

type RequestHandler struct {
	uc *usecase.RequestUseCase
}

func NewRequestHandler(uc *usecase.RequestUseCase) *RequestHandler {
	return &RequestHandler{uc: uc}
}

func (h *RequestHandler) Create(c *gin.Context) {
	var input usecase.CreateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := h.uc.CreateRequest(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

func (h *RequestHandler) GetActive(c *gin.Context) {
	restaurantID := c.Param("restaurantId")

	requests, err := h.uc.GetActiveRequests(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}

func (h *RequestHandler) Complete(c *gin.Context) {
	id := c.Param("id")

	if err := h.uc.CompleteRequest(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "request completed"})
}

func (h *RequestHandler) GetTableStatus(c *gin.Context) {
	tableID := c.Param("id")

	requests, err := h.uc.GetTableStatus(tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, requests)
}
