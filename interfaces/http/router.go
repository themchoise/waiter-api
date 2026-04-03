package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ws "github.com/waiter/back/infrastructure/websocket"
)

func SetupRouter(
	requestHandler *RequestHandler,
	feedbackHandler *FeedbackHandler,
	restaurantHandler *RestaurantHandler,
	hub *ws.Hub,
) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		// Client endpoints
		api.POST("/requests", requestHandler.Create)
		api.GET("/tables/:tableId/status", requestHandler.GetTableStatus)
		api.POST("/feedback", feedbackHandler.Create)

		// Restaurant endpoints
		api.GET("/restaurants/:restaurantId/requests/active", requestHandler.GetActive)
		api.PATCH("/requests/:requestId", requestHandler.Complete)

		// Restaurant management
		api.POST("/restaurants", restaurantHandler.Create)
		api.GET("/restaurants/:restaurantId", restaurantHandler.Get)
		api.POST("/restaurants/:restaurantId/tables", restaurantHandler.CreateTable)
		api.GET("/restaurants/:restaurantId/tables", restaurantHandler.GetTables)

		// WebSocket
		wsHandler := NewWSHandler(hub)
		api.GET("/ws/:restaurantId", wsHandler.Connect)
	}

	return r
}
