package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waiter/back/application/usecase"
	"github.com/waiter/back/domain/entity"
)

type RequestHandler struct {
	uc *usecase.RequestUseCase
}

func NewRequestHandler(uc *usecase.RequestUseCase) *RequestHandler {
	return &RequestHandler{uc: uc}
}

// Create godoc
// @Summary      Crear solicitud
// @Description  Crea una nueva solicitud de mesa (llamar mozo, pedir cuenta, ayuda)
// @Tags         requests
// @Accept       json
// @Produce      json
// @Param        request  body      usecase.CreateRequestInput  true  "Datos de la solicitud"
// @Success      201      {object}  entity.Request
// @Failure      400      {object}  map[string]string
// @Failure      422      {object}  map[string]string
// @Router       /api/v1/requests [post]
func (h *RequestHandler) Create(c *gin.Context) {
	var input usecase.CreateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation_error", "message": err.Error()})
		return
	}

	req, err := h.uc.CreateRequest(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable_entity", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// GetActive godoc
// @Summary      Obtener solicitudes activas
// @Description  Retorna las solicitudes activas de un restaurante
// @Tags         requests
// @Produce      json
// @Param        restaurantId  path      string  true  "ID del restaurante"
// @Success      200           {array}   entity.Request
// @Failure      500           {object}  map[string]string
// @Router       /api/v1/restaurants/{restaurantId}/requests/active [get]
func (h *RequestHandler) GetActive(c *gin.Context) {
	restaurantID := c.Param("restaurantId")

	requests, err := h.uc.GetActiveRequests(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error", "message": "failed to retrieve requests"})
		return
	}

	c.JSON(http.StatusOK, requests)
}

type updateRequestInput struct {
	Status entity.RequestStatus `json:"status" binding:"required"`
}

// Complete godoc
// @Summary      Actualizar estado de solicitud
// @Description  Actualiza el estado de una solicitud (marcar como atendida)
// @Tags         requests
// @Accept       json
// @Produce      json
// @Param        requestId  path      string              true  "ID de la solicitud"
// @Param        request    body      updateRequestInput   true  "Nuevo estado"
// @Success      204
// @Failure      400        {object}  map[string]string
// @Failure      422        {object}  map[string]string
// @Router       /api/v1/requests/{requestId} [patch]
func (h *RequestHandler) Complete(c *gin.Context) {
	id := c.Param("requestId")

	var input updateRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation_error", "message": err.Error()})
		return
	}

	if input.Status != entity.Done {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid_status", "message": "only DONE status transition is supported"})
		return
	}

	if err := h.uc.CompleteRequest(id); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable_entity", "message": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetTableStatus godoc
// @Summary      Obtener estado de mesa
// @Description  Retorna las solicitudes de una mesa específica
// @Tags         requests
// @Produce      json
// @Param        tableId  path      string  true  "ID de la mesa"
// @Success      200      {array}   entity.Request
// @Failure      500      {object}  map[string]string
// @Router       /api/v1/tables/{tableId}/status [get]
func (h *RequestHandler) GetTableStatus(c *gin.Context) {
	tableID := c.Param("tableId")

	requests, err := h.uc.GetTableStatus(tableID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error", "message": "failed to retrieve table status"})
		return
	}

	c.JSON(http.StatusOK, requests)
}
