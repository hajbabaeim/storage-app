package data

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	BatchSize     = 1000 // Process 1000 lines at a time
	MaxGoroutines = 100  // Maximum number of concurrent goroutines
)

var sem = make(chan struct{}, MaxGoroutines) // Semaphore to limit concurrent goroutines

func LoadDataConcurrentEvery30Min() {
	for {
		LoadDataConcurrent("promotions.csv")
		time.Sleep(30 * time.Minute)
	}
}

func LoadDataConcurrent(filePath string) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	batch := make([]Promotion, 0, BatchSize)
	var lineNumber uint64 = 0

	for {
		line, err := reader.Read()
		if err == io.EOF {
			go func(b []Promotion, startLineNumber uint64) {
				sem <- struct{}{}
				processBatchConcurrent(b, startLineNumber) // Pass the start line number for this batch
				<-sem
			}(batch, lineNumber-uint64(len(batch)))
			break
		} else if err != nil {
			panic(err)
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		promotion := Promotion{
			ID:             line[0], // You may want to remove this line if you're not using the ID field anymore
			Price:          price,
			ExpirationDate: line[2],
		}

		batch = append(batch, promotion)

		if uint64(len(batch)) >= BatchSize {
			go func(b []Promotion, startLineNumber uint64) {
				sem <- struct{}{}
				processBatchConcurrent(b, startLineNumber) // Pass the start line number for this batch
				<-sem
			}(batch, lineNumber-uint64(len(batch))+1)
			batch = make([]Promotion, 0, BatchSize) // Reset batch
		}

		lineNumber++
	}

	// Wait for all goroutines to finish
	for i := 0; i < cap(sem); i++ {
		sem <- struct{}{}
	}
}

func processBatchConcurrent(batch []Promotion, startLineNumber uint64) {
	DataLock.Lock()
	for i, promotion := range batch {
		PromotionData[startLineNumber+uint64(i)] = promotion
	}
	DataLock.Unlock()
	println("Processed a batch of size:", len(batch))
}
