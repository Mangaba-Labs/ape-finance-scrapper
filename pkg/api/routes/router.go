package router

import (
	"github.com/Mangaba-Labs/ape-finance-scrapper/database"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/api/handler"
	stock "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/handler"
	stockRepository "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/repository"
	stockService "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/stock/service"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router pkg
func SetupRoutes(app *fiber.App) error {
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	repositoryStock := stockRepository.Repository{
		DB: db,
	}
	serviceStock := stockService.NewUserService(repositoryStock)
	stockHandler := stock.NewStockHandler(serviceStock)
	// Api base
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Health
	health := v1.Group("/health")
	health.Get("/", handler.HealthCheck)

	stock := v1.Group("/shares")
	stock.Get("/", stockHandler.GetAllStocks)
	stock.Get("/:id", stockHandler.GetStockByID)
	stock.Delete("/:id", stockHandler.DeleteStock)
	stock.Post("/:bvmf", stockHandler.AddStock)

	return nil
}