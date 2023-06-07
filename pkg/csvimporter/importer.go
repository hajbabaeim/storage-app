package csvimporter

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"storage-app/internal/model"
	"storage-app/internal/service"
	"strconv"
	"time"
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

	fmt.Println("-----> ImportCSV 1")
	// Delete and insert in batches
	for i := 0; i < len(records); i += batchSize {
		end := i + batchSize
		if end > len(records) {
			end = len(records)
		}

		fmt.Println("-----> ImportCSV 2")
		// Start a transaction
		tx, err := promotionSvc.BeginTransaction(ctx)
		if err != nil {
			return err
		}

		fmt.Println("-----> ImportCSV 3")
		// Delete a batch
		for _, record := range records[i:end] {
			i, err := strconv.Atoi(record[0])
			_, err = promotionSvc.DeleteByID(ctx, i) // Assume record[0] is the id
			if err != nil {
				promotionSvc.RollbackTransaction(ctx, tx) // Rollback transaction on error
				return err
			}
		}

		fmt.Println("-----> ImportCSV 4")
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

func ImportCSV(filePath string, svc *service.PromotionService) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	batch := make([]*model.Promotion, 0, batchSize)

	lineNumber := 1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return err
		}

		expirationDate, err := time.Parse("2006-01-02 15:04:05 -0700 MST", record[2])
		if err != nil {
			return err
		}

		batch = append(batch, &model.Promotion{
			ID:             lineNumber,
			PID:            record[0],
			Price:          price,
			ExpirationDate: expirationDate,
		})

		if len(batch) >= batchSize {
			if err := insertBatch(context.Background(), svc, batch); err != nil {
				return err
			}
			batch = batch[:0]
		}

		lineNumber++
	}

	// Insert the remaining data
	if len(batch) > 0 {
		if err := insertBatch(context.Background(), svc, batch); err != nil {
			return err
		}
	}

	return nil
}

func insertBatch(ctx context.Context, svc *service.PromotionService, batch []*model.Promotion) error {
	for _, promotion := range batch {
		if err := svc.AddPromotion(ctx, promotion); err != nil {
			return err
		}
	}
	return nil
}
