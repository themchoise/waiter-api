package http

import (
	"github.com/gin-gonic/gin"
	ws "github.com/waiter/back/infrastructure/websocket"
)

type WSHandler struct {
	hub *ws.Hub
}

func NewWSHandler(hub *ws.Hub) *WSHandler {
	return &WSHandler{hub: hub}
}

func (h *WSHandler) Connect(c *gin.Context) {
	restaurantID := c.Param("restaurantId")
	h.hub.Subscribe(c.Writer, c.Request, restaurantID)
}
