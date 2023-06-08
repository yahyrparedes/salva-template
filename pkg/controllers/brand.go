package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yahyrparedes/salva-template/pkg/services"
	"strconv"
)

// GetBrands func gets all exists brands.
// @Description Get all exists brands.
// @Summary get all exists brands
// @Tags Brands
// @Accept json
// @Produce json
// @Success 200 {array} mappers.Brand
// @Router /v1/brands [get]
func GetBrands(c *fiber.Ctx) error {
	brands := services.GetBrands()
	return c.Status(fiber.StatusOK).JSON(brands)
}

// GetBrand func gets brand by given ID or 404 error.
// @Description Get brand by given ID.
// @Summary get brand by given ID
// @Tags Brand
// @Accept json
// @Produce json
// @Param id path string true "Brand ID"
// @Success 200 {object} mappers.Brand
// @Router /v1/brands/{id} [get]
func GetBrand(c *fiber.Ctx) error {
	id, _ := strconv.Atoi((c.Params("id")))
	brand := services.GetBrand(id)
	return c.Status(fiber.StatusOK).JSON(brand)
}
