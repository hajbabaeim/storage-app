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
)

func ImportCSV(filePath string, svc *service.PromotionService) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Batch size for batch processing
	const batchSize = 1000
	batch := make([]*model.Promotion, 0, batchSize)

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
			ID:             record[0],
			Price:          price,
			ExpirationDate: expirationDate,
		})

		if len(batch) >= batchSize {
			if err := insertBatch(context.Background(), svc, batch); err != nil {
				return err
			}
			batch = batch[:0]
		}
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
