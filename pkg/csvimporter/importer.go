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

func ImportCSVConcurrent(file string, promotionSvc *service.PromotionService) error {
	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a new reader
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Delete and insert in batches
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		// Start a transaction
		tx, err := promotionSvc.BeginTransaction(ctx)
		if err != nil {
			return err
		}

		// Delete a batch
		for _, record := range records[i:end] {
			i, err := strconv.Atoi(record[0])
			_, err = promotionSvc.DeleteByID(ctx, i) // Assume record[0] is the id
			if err != nil {
				promotionSvc.RollbackTransaction(ctx, tx) // Rollback transaction on error
				return err
			}
		}

		// Insert a batch
		for _, record := range records[i:end] {
			price, err := parsePrice(record[1])
			if err != nil {
				promotionSvc.RollbackTransaction(ctx, tx) // Rollback transaction on error
				return err
			}
			expirationDate, err := parseDate(record[2])
			if err != nil {
				promotionSvc.RollbackTransaction(ctx, tx) // Rollback transaction on error
				return err
			}
			promotion := &model.Promotion{
				PID:            record[0],
				Price:          price,
				ExpirationDate: expirationDate,
			}
			err = promotionSvc.Insert(ctx, promotion)
			if err != nil {
				promotionSvc.RollbackTransaction(ctx, tx) // Rollback transaction on error
				return err
			}
		}

		// Commit the transaction
		err = promotionSvc.CommitTransaction(ctx, tx)
		if err != nil {
			return err
		}
	}

	return nil
}

func ImportCSV(filePath string) (<-chan []*model.Promotion, error) {
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

// func insertBatch(ctx context.Context, svc *service.PromotionService, batch []*model.Promotion) error {
// 	for _, promotion := range batch {
// 		if err := svc.AddPromotion(ctx, promotion); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func BatchInsert(ctx context.Context, svc *service.PromotionService, batches <-chan []*model.Promotion) error {
	for batch := range batches {
		err := svc.InsertBatch(ctx, batch)
		if err != nil {
			return err
		}
	}
	return nil
}
