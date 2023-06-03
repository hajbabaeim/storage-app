package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"storage/data"

	"github.com/labstack/echo/v4"
)

func main() {
	go data.LoadDataEvery30Min()
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
