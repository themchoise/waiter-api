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

// Create godoc
// @Summary      Crear restaurante
// @Description  Crea un nuevo restaurante
// @Tags         restaurants
// @Accept       json
// @Produce      json
// @Param        request  body      usecase.CreateRestaurantInput  true  "Datos del restaurante"
// @Success      201      {object}  entity.Restaurant
// @Failure      400      {object}  map[string]string
// @Failure      422      {object}  map[string]string
// @Router       /api/v1/restaurants [post]
func (h *RestaurantHandler) Create(c *gin.Context) {
	var input usecase.CreateRestaurantInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation_error", "message": err.Error()})
		return
	}

	r, err := h.uc.CreateRestaurant(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable_entity", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, r)
}

// Get godoc
// @Summary      Obtener restaurante
// @Description  Retorna un restaurante por su ID
// @Tags         restaurants
// @Produce      json
// @Param        restaurantId  path      string  true  "ID del restaurante"
// @Success      200           {object}  entity.Restaurant
// @Failure      404           {object}  map[string]string
// @Router       /api/v1/restaurants/{restaurantId} [get]
func (h *RestaurantHandler) Get(c *gin.Context) {
	id := c.Param("restaurantId")

	r, err := h.uc.GetRestaurant(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not_found", "message": "restaurant not found"})
		return
	}

	c.JSON(http.StatusOK, r)
}

type createTableInput struct {
	Number int `json:"number" binding:"required"`
}

// CreateTable godoc
// @Summary      Crear mesa
// @Description  Crea una nueva mesa para un restaurante
// @Tags         tables
// @Accept       json
// @Produce      json
// @Param        restaurantId  path      string            true  "ID del restaurante"
// @Param        request       body      createTableInput  true  "Datos de la mesa"
// @Success      201           {object}  entity.Table
// @Failure      400           {object}  map[string]string
// @Failure      422           {object}  map[string]string
// @Router       /api/v1/restaurants/{restaurantId}/tables [post]
func (h *RestaurantHandler) CreateTable(c *gin.Context) {
	restaurantID := c.Param("restaurantId")

	var input createTableInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation_error", "message": err.Error()})
		return
	}

	t, err := h.uc.CreateTable(usecase.CreateTableInput{
		Number:       input.Number,
		RestaurantID: restaurantID,
	})
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "unprocessable_entity", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t)
}

// GetTables godoc
// @Summary      Listar mesas
// @Description  Retorna las mesas de un restaurante
// @Tags         tables
// @Produce      json
// @Param        restaurantId  path      string  true  "ID del restaurante"
// @Success      200           {array}   entity.Table
// @Failure      500           {object}  map[string]string
// @Router       /api/v1/restaurants/{restaurantId}/tables [get]
func (h *RestaurantHandler) GetTables(c *gin.Context) {
	restaurantID := c.Param("restaurantId")

	tables, err := h.uc.GetTables(restaurantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_error", "message": "failed to retrieve tables"})
		return
	}

	c.JSON(http.StatusOK, tables)
}
