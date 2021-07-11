package handler

import (
	"strconv"

	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/models"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service service.StockService
}

func (h *Handler) AddStock(c *fiber.Ctx) error {
	bvmf := c.Params("bvmf")
	res := h.service.Create(bvmf)
	return c.Status(res.HttpCode).JSON(fiber.Map{"status": res.Status, "message": res.Message})
}

func (h *Handler) GetAllStocks(c *fiber.Ctx) error {
	stocks, res := h.service.GetAll()
	return c.Status(res.HttpCode).JSON(fiber.Map{"status": res.Status, "message": res.Message, "result": stocks})
}

func (h *Handler) GetStockByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": models.ERROR, "message": "Bad request!"})
	}
	stock, res := h.service.GetByID(id)
	return c.Status(res.HttpCode).JSON(fiber.Map{"status": res.Status, "message": res.Message, "result": stock})
}

func (h *Handler) DeleteStock(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": models.ERROR, "message": "Bad request!"})
	}
	res := h.service.Delete(id)
	return c.Status(res.HttpCode).JSON(fiber.Map{"status": res.Status, "message": res.Message})
}