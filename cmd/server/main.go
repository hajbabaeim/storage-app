package main

//go:generate go run entgo.io/ent/cmd/ent generate ./ent/schema

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"storage-app/internal/handler"
	"storage-app/internal/repository"
	"storage-app/internal/service"
	"storage-app/internal/utils"
	"storage-app/pkg/csvimporter"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	csvPath = "data/promotions.csv"
	dbUri   = "postgresql://<username>:<password>@localhost:5432/<db>?sslmode=disable"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	app := fiber.New()
	app.Use(logger.New())
	postgresConnURL, errCreateDBUrl := utils.ConnectionURLBuilder("postgres")
	if errCreateDBUrl != nil {
		log.Fatal().Err(errCreateDBUrl).Msg("failed to get database uri")
	}
	dbClient, err := repository.NewPostgresDb(postgresConnURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	promotionRepository := repository.NewPromotionRepository(dbClient)
	promotionService := service.NewPromotionService(promotionRepository)
	promotionHandler := handler.NewPromotionHandler(promotionService)

	app.Get("/promotions/:id", promotionHandler.GetPromotion)

	log.Info().Msg("Starting CSV import process")
	absPath, _ := filepath.Abs(csvPath)

	promotionBatches, err := csvimporter.ImportCSV(absPath)
	if err != nil {
		log.Error().Err(err).Msg("Failed to import CSV")
		return
	}

	go csvimporter.BatchInsert(context.Background(), promotionService, promotionBatches)

	// if err := csvimporter.ImportCSV(absPath, promotionService); err != nil {
	// 	log.Fatal().Err(err).Msg("failed to import CSV data")
	// }

	port := os.Getenv("SERVER_PORT")
	log.Info().Msg(fmt.Sprintf("Starting server on port %s", port))
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
