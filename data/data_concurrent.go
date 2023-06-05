package data

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	BatchSize     uint64 = 1000 // Process 1000 lines at a time
	MaxGoroutines uint64 = 100  // Maximum number of concurrent goroutines
)

var sem = make(chan struct{}, MaxGoroutines) // Semaphore to limit concurrent goroutines

type PromotionWithLineNumber struct {
	Promotion
	LineNumber uint64
}

func LoadDataConcurrentEvery30Min() {
	for {
		absPath, _ := filepath.Abs("data/promotions.csv")
		LoadDataConcurrent(absPath)
		time.Sleep(30 * time.Minute)
	}
}

func LoadDataConcurrent(filePath string) {
	lines := readCsv(filePath)
	promotions := processBatches(lines)
	storeBatches(promotions)
}

func readCsv(filePath string) <-chan []string {
	out := make(chan []string)

	go func() {
		csvFile, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)

		for {
			line, err := reader.Read()
			if err == io.EOF {
				close(out)
				break
			} else if err != nil {
				panic(err)
			}

			out <- line
		}
	}()

	return out
}

func processBatches(lines <-chan []string) <-chan PromotionWithLineNumber {
	out := make(chan PromotionWithLineNumber)

	go func() {
		var lineNumber uint64 = 1 // Start from 1, because CSV header is not a record.
		batch := make([]Promotion, 0, BatchSize)

		for line := range lines {
			price, _ := strconv.ParseFloat(line[1], 64)

			promotion := Promotion{
				ID:             line[0],
				Price:          price,
				ExpirationDate: line[2],
			}

			batch = append(batch, promotion)
			lineNumber++

			if uint64(len(batch)) >= BatchSize {
				go processBatchConcurrent(batch, lineNumber-BatchSize, out)
				batch = make([]Promotion, 0, BatchSize)
			}
		}

		if len(batch) > 0 {
			go processBatchConcurrent(batch, lineNumber-uint64(len(batch)), out)
		}
	}()

	return out
}

func processBatchConcurrent(batch []Promotion, startLineNumber uint64, out chan<- PromotionWithLineNumber) {
	sem <- struct{}{}
	defer func() { <-sem }()
	for i, promotion := range batch {
		out <- PromotionWithLineNumber{
			Promotion:  promotion,
			LineNumber: startLineNumber + uint64(i),
		}
	}
}

func storeBatches(promotions <-chan PromotionWithLineNumber) {
	DataLock.Lock()
	defer DataLock.Unlock()

	for p := range promotions {
		PromotionData[p.LineNumber] = p.Promotion
	}
}
