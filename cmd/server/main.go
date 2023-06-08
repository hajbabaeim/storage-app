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
	"time"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const (
	csvPath = "data/promotions.csv"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	redisConnURL, errCreateRedisUrl := utils.ConnectionURLBuilder("redis")
	if errCreateRedisUrl != nil {
		log.Fatal().Err(errCreateRedisUrl).Msg("failed to get database uri")
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConnURL,
		Password: "",
		DB:       0,
	})

	app := fiber.New()
	app.Use(logger.New())
	// Use rate limiter for api to perform in peak periods
	app.Use(limiter.New(limiter.Config{
		Max:        50,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
		SkipFailedRequests:     false,
		SkipSuccessfulRequests: false,
		LimiterMiddleware:      limiter.FixedWindow{},
	}))

	postgresConnURL, errCreateDBUrl := utils.ConnectionURLBuilder("postgres")
	if errCreateDBUrl != nil {
		log.Fatal().Err(errCreateDBUrl).Msg("failed to get database uri")
	}
	dbClient, err := repository.NewPostgresDb(postgresConnURL)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	promotionRepository := repository.NewPromotionRepository(dbClient)
	promotionService := service.NewPromotionService(promotionRepository, rdb)
	promotionHandler := handler.NewPromotionHandler(promotionService)

	app.Get("/promotions/:id", promotionHandler.GetPromotion)

	log.Info().Msg("Starting CSV import process")
	absPath, _ := filepath.Abs(csvPath)
	if !fileExists(absPath) {
		log.Info().Msg("File does not exist")
	} else {
		log.Info().Msg("File exists")
	}
	ticker := time.NewTicker(30 * time.Minute)
	quit := make(chan struct{})
	go func() {
		var previousChecksum string
		for {
			select {
			case <-ticker.C:
				currentChecksum, err := utils.CalculateMD5(absPath)
				if err != nil {
					log.Error().Err(err).Msg("Failed to calculate MD5 checksum of the file")
					continue
				}

				if currentChecksum == previousChecksum {
					log.Info().Msg("CSV file does not change, just update the info if does not exist")
					err = csvimporter.ImportCSV(context.Background(), absPath, promotionService, false)
					if err != nil {
						log.Error().Err(err).Msg("Failed to import CSV")
						continue
					}
					continue
				}

				log.Info().Msg("CSV file has been changed, so it should be replaced")
				err = csvimporter.ImportCSV(context.Background(), absPath, promotionService, true)
				if err != nil {
					log.Error().Err(err).Msg("Failed to import CSV")
					continue
				}

				// Update the previousChecksum with the current one after a successful import
				previousChecksum = currentChecksum

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// promotionBatches, err := csvimporter.ImportCSVWithoutDeletion(absPath)
	// if err != nil {
	// 	log.Error().Err(err).Msg("Failed to import CSV")
	// 	return
	// }
	// go csvimporter.BatchInsert(context.Background(), promotionService, promotionBatches)

	port := os.Getenv("SERVER_PORT")
	log.Info().Msg(fmt.Sprintf("Starting server on port %s", port))
	if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
