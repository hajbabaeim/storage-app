package data

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
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

var PromotionData = make(map[uint64]Promotion)
var DataLock = &sync.RWMutex{}

func AddTestPromotion(id uint64, promotionId string, price float64, expirationDate string) {
	promotion := Promotion{
		ID:             promotionId,
		Price:          price,
		ExpirationDate: expirationDate,
	}

	DataLock.Lock()
	PromotionData[id] = promotion
	DataLock.Unlock()
}

func LoadDataEvery30Min() {
	for {
		absPath, _ := filepath.Abs("data/promotions.csv")
		LoadData(absPath)
		time.Sleep(30 * time.Minute)
	}
}

func LoadData(filePath string) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(errors.New(fmt.Sprintf("Error during opening the file: %s", err)))
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	var count uint64 = 0
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(errors.New(fmt.Sprintf("Error during reading the file: %s", err)))
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		promotion := Promotion{
			ID:             line[0],
			Price:          price,
			ExpirationDate: line[2],
		}

		DataLock.Lock()
		PromotionData[count] = promotion
		count++
		DataLock.Unlock()
	}
}
