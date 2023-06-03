package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"storage/data"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetPromotionByID(t *testing.T) {
	e := echo.New()
	var testId uint64 = 1
	testPromotionId := "test-id"
	data.AddTestPromotion(testId, testPromotionId, 60.683466, "2018-08-04 05:32:31 +0200 CEST")

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/promotions/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.FormatUint(testId, 10))

	if assert.NoError(t, getPromotionByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), testPromotionId)
	}
}
