package http

import (
	"github.com/gin-gonic/gin"
	ws "github.com/waiter/back/infrastructure/websocket"
)

func SetupRouter(
	requestHandler *RequestHandler,
	feedbackHandler *FeedbackHandler,
	restaurantHandler *RestaurantHandler,
	hub *ws.Hub,
) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")
	{
		// Client endpoints
		api.POST("/requests", requestHandler.Create)
		api.GET("/table/:id/status", requestHandler.GetTableStatus)
		api.POST("/feedback", feedbackHandler.Create)

		// Restaurant endpoints
		api.GET("/restaurants/:restaurantId/requests/active", requestHandler.GetActive)
		api.PATCH("/requests/:id/complete", requestHandler.Complete)

		// Restaurant management
		api.POST("/restaurants", restaurantHandler.Create)
		api.GET("/restaurants/:id", restaurantHandler.Get)
		api.POST("/tables", restaurantHandler.CreateTable)
		api.GET("/restaurants/:id/tables", restaurantHandler.GetTables)

		// WebSocket
		wsHandler := NewWSHandler(hub)
		api.GET("/ws/:restaurantId", wsHandler.Connect)
	}

	return r
}
