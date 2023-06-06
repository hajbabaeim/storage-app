package handler

import (
	"storage-app/internal/service"

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

	promotion, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "promotion not found",
		})
	}

	return c.JSON(promotion)
}
