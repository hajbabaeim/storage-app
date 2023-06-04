package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"storage/data"
	"strconv"

	"github.com/labstack/echo/v4"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "concurrent", "define running mode with concurrency or not")
	if mode == "concurrent" {
		go data.LoadDataConcurrentEvery30Min() // using for data heavy processes (large csv file)
	} else {
		go data.LoadDataEvery30Min()
	}
	e := echo.New()
	e.GET("/promotions/:id", getPromotionByID)
	e.Start(":1321")
}

func getPromotionByID(c echo.Context) error {
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Fatalln(errors.New(fmt.Sprintf("Input id is not valid: %s", err)))
	}
	data.DataLock.RLock()
	promotion, ok := data.PromotionData[uintId]
	data.DataLock.RUnlock()

	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Not found"})
	}

	return c.JSON(http.StatusOK, promotion)
}
