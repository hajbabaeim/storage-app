package csvimporter

import (
	"context"
	"encoding/csv"
	"io"
	"os"
	"storage-app/internal/model"
	"storage-app/internal/service"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

// Batch size for batch processing
const batchSize = 10000

func parsePrice(priceStr string) (float64, error) {
	return strconv.ParseFloat(priceStr, 64)
}

func parseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05 -0700 MST", dateStr)
}

func ImportCSVWithoutDeletion(filePath string) (<-chan []*model.Promotion, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	out := make(chan []*model.Promotion)
	go func() {
		defer close(out)
		defer file.Close()

		batch := make([]*model.Promotion, 0, batchSize)

		lineNumber := 1
		for {
			record, err := reader.Read()
			if err == io.EOF {
				if len(batch) > 0 {
					out <- batch
				}
				break
			}
			if err != nil {
				log.Error().Err(err).Msg("Error during read CSV")
				continue
			}

			price, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Error().Err(err).Msg("Error during read CSV")
				continue
			}

			expirationDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", record[2])
			if err != nil {
				log.Error().Err(err).Msg("Error during read CSV")
				continue
			}

			batch = append(batch, &model.Promotion{
				ID:             lineNumber,
				PID:            record[0],
				Price:          price,
				ExpirationDate: expirationDate,
			})

			if len(batch) >= batchSize {
				out <- batch
				batch = make([]*model.Promotion, 0, batchSize)
			}
		}
	}()

	return out, nil
}

func ImportCSV(ctx context.Context, filePath string, svc *service.PromotionService, reset bool) error {
	tx, err := svc.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	if reset {
		_, err := svc.DeleteAll(ctx)
		if err != nil {
			_ = svc.RollbackTransaction(ctx, tx)
			return err
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		_ = svc.RollbackTransaction(ctx, tx)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	batch := make([]*model.Promotion, 0, batchSize)

	lineNumber := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			_ = svc.RollbackTransaction(ctx, tx)
			return err
		}

		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			_ = svc.RollbackTransaction(ctx, tx)
			return err
		}

		expirationDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", record[2])
		if err != nil {
			_ = svc.RollbackTransaction(ctx, tx)
			return err
		}

		batch = append(batch, &model.Promotion{
			ID:             lineNumber,
			PID:            record[0],
			Price:          price,
			ExpirationDate: expirationDate,
		})

		if len(batch) >= batchSize {
			if err := svc.BatchInsert(ctx, batch); err != nil {
				_ = svc.RollbackTransaction(ctx, tx)
				return err
			}
			batch = batch[:0]
		}
		lineNumber++
	}

	if len(batch) > 0 {
		if err := svc.BatchInsert(ctx, batch); err != nil {
			_ = svc.RollbackTransaction(ctx, tx)
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func BatchInsert(ctx context.Context, svc *service.PromotionService, batches <-chan []*model.Promotion) error {
	for batch := range batches {
		err := svc.InsertBatch(ctx, batch)
		if err != nil {
			return err
		}
	}
	return nil
}
