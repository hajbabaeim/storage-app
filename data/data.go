package data

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type Promotion struct {
	ID             string  `json:"id"`
	Price          float64 `json:"price"`
	ExpirationDate string  `json:"expiration_date"`
}

type PromotionWithLineNumber struct {
	Promotion
	LineNumber uint64
}

var (
	PromotionData = make(map[uint64]Promotion)
	DataLock      sync.RWMutex
)

func LoadDataEvery30Min() {
	for {
		absPath, _ := filepath.Abs("data/promotions.csv")
		LoadData(absPath)
		time.Sleep(30 * time.Minute)
	}
}

func LoadData(filePath string) {
	lines := readCsv(filePath)
	promotions := processLines(lines)
	storeData(promotions)
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
			if err != nil {
				close(out)
				break
			}
			out <- line
		}
	}()

	return out
}

func processLines(lines <-chan []string) <-chan PromotionWithLineNumber {
	out := make(chan PromotionWithLineNumber)

	go func() {
		var lineNumber uint64 = 1 // Start from 1, because CSV header is not a record.
		for line := range lines {
			price, _ := strconv.ParseFloat(line[1], 64)
			out <- PromotionWithLineNumber{
				Promotion: Promotion{
					ID:             line[0],
					Price:          price,
					ExpirationDate: line[2],
				},
				LineNumber: lineNumber,
			}
			lineNumber++
		}
		close(out)
	}()

	return out
}

func storeData(promotions <-chan PromotionWithLineNumber) {
	DataLock.Lock()
	PromotionData = make(map[uint64]Promotion) // Clear the data before loading the new one
	for p := range promotions {
		PromotionData[p.LineNumber] = p.Promotion
	}
	DataLock.Unlock()
}
