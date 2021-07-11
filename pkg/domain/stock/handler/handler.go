package handler

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/service"
	"github.com/gofiber/fiber/v2"
)

// StockHandler contract
type StockHandler interface {
	AddStock(c *fiber.Ctx) error
	GetAllStock(c *fiber.Ctx) error
	GetStockByID(c *fiber.Ctx) error
	DeletStock(c *fiber.Ctx) error
}

// NewStockHandler returns a pointer to an handler impl
func NewStockHandler(s service.StockService) Handler {
	return Handler{
		service: s,
	}
}
