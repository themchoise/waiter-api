package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waiter/back/application/usecase"
)

type RestaurantHandler struct {
	uc *usecase.RestaurantUseCase
}

func NewRestaurantHandler(uc *usecase.RestaurantUseCase) *RestaurantHandler {
	return &RestaurantHandler{uc: uc}
}

func (h *RestaurantHandler) Create(c *gin.Context) {
	var input usecase.CreateRestaurantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := h.uc.CreateRestaurant(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, r)
}

func (h *RestaurantHandler) Get(c *gin.Context) {
	id := c.Param("id")

	r, err := h.uc.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, r)
}

func (h *RestaurantHandler) CreateTable(c *gin.Context) {
	var input usecase.CreateTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := h.uc.CreateTable(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}

func (h *RestaurantHandler) GetTables(c *gin.Context) {
	restaurantID := c.Param("id")

	tables, err := h.uc.GetTables(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tables)
}
