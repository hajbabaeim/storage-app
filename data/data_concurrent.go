package data

import (
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

const (
	BatchSize     uint64 = 1000 // Process 1000 lines at a time
	MaxGoroutines uint64 = 100  // Maximum number of concurrent goroutines
)

var sem = make(chan struct{}, MaxGoroutines) // Semaphore to limit concurrent goroutines
var once sync.Once

func LoadDataConcurrentEvery30Min(signal chan<- struct{}) {
	for {
		absPath, _ := filepath.Abs("data/promotions.csv")
		LoadDataConcurrent(absPath, signal)
		time.Sleep(30 * time.Minute)
	}
}

func LoadDataConcurrent(filePath string, signal chan<- struct{}) {
	lines := readCsv(filePath)
	promotions := processBatches(lines)
	storeBatches(promotions, signal)
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

func storeBatches(promotions <-chan PromotionWithLineNumber, signal chan<- struct{}) {
	temp := make(map[string]Promotion)
	for p := range promotions {
		temp[strconv.FormatUint(p.LineNumber, 10)] = p.Promotion
		if len(temp) == int(BatchSize) {
			DataLock.Lock()
			for k, v := range temp {
				value, _ := strconv.ParseUint(k, 10, 64)
				PromotionData[value] = v
			}
			DataLock.Unlock()
			temp = make(map[string]Promotion)
			once.Do(func() { close(signal) })
		}
	}
	if len(temp) > 0 {
		DataLock.Lock()
		for k, v := range temp {
			value, _ := strconv.ParseUint(k, 10, 64)
			PromotionData[value] = v
		}
		DataLock.Unlock()
	}
}
