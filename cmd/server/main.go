package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/waiter/back/application/usecase"
	"github.com/waiter/back/config"
	_ "github.com/waiter/back/docs"
	"github.com/waiter/back/infrastructure/persistence"
	ws "github.com/waiter/back/infrastructure/websocket"
	handler "github.com/waiter/back/interfaces/http"
)

// @title          Waiter API
// @version        1.0
// @description    Sistema de llamado de mozos para restaurantes
// @host           localhost:8080
// @BasePath       /api/v1
func main() {
	_ = godotenv.Load()

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))

	cfg := config.Load()

	db, err := persistence.NewDatabase(cfg.DSN)
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}

	// Repositories
	restaurantRepo := persistence.NewRestaurantRepo(db)
	tableRepo := persistence.NewTableRepo(db)
	requestRepo := persistence.NewRequestRepo(db)
	feedbackRepo := persistence.NewFeedbackRepo(db)

	// WebSocket hub
	hub := ws.NewHub()

	// Use cases
	requestUC := usecase.NewRequestUseCase(requestRepo, tableRepo, hub)
	feedbackUC := usecase.NewFeedbackUseCase(feedbackRepo, tableRepo)
	restaurantUC := usecase.NewRestaurantUseCase(restaurantRepo, tableRepo)

	// Handlers
	requestHandler := handler.NewRequestHandler(requestUC)
	feedbackHandler := handler.NewFeedbackHandler(feedbackUC)
	restaurantHandler := handler.NewRestaurantHandler(restaurantUC)

	// Router
	router := handler.SetupRouter(requestHandler, feedbackHandler, restaurantHandler, hub)

	slog.Info("server starting", "port", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
