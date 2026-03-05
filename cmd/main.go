package main

import (
	"log/slog"
	"os"
	"pizza-tracker-go/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cnfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cnfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	setupRoutes(router, h)

	slog.Info("Server starting", "url", "http://localhost:"+cnfg.Port)

	if err := router.Run(":" + cnfg.Port); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}
}
