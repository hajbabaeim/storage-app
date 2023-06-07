package handler

import (
	"fmt"
	"storage-app/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler struct {
	service *service.PromotionService
}

func NewPromotionHandler(service *service.PromotionService) *PromotionHandler {
	return &PromotionHandler{
		service: service,
	}
}

func (h *PromotionHandler) GetPromotion(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Println("---", id)
	i, err := strconv.Atoi(id)
	promotion, err := h.service.GetByID(c.Context(), i)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "promotion not found",
		})
	}

	return c.JSON(promotion)
}

// func (h *PromotionHandler) GetPromotionByDBID(c *fiber.Ctx) error {
// 	dbID, err := strconv.ParseInt(c.Params("id"), 10, 64)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
// 			"error": "Invalid database ID",
// 		})
// 	}

// 	promotion, err := h.service.GetByDBID(dbID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
// 				"error": "Promotion not found",
// 			})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
// 			"error": "Error getting promotion",
// 		})
// 	}

// 	return c.JSON(promotion)
// }
