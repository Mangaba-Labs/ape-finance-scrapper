package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Mangaba-Labs/ape-finance-scrapper/database"
	router "github.com/Mangaba-Labs/ape-finance-scrapper/pkg/api/routes"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/config"
	"github.com/Mangaba-Labs/ape-finance-scrapper/pkg/domain/cron"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Setting up environment variables
	config.SetupEnvVars()

	// Setting up database connection
	db, err := database.NewDatabase()

	if err != nil {
		log.Fatalf("cannot migrate database, err: %s", err.Error())
	}

	// Running migrations
	migrations := config.Migrate{DB: db}

	err = migrations.MigrateAll()

	if err != nil {
		log.Fatalf("cannot migrate database, err: %s", err.Error())
	}

	app := fiber.New()
	// Helmet Security
	app.Use(helmet.New())

	//Handle Cors
	app.Use(cors.New())

	//Handle panics
	app.Use(recover.New())

	// //Handle logs
	app.Use(logger.New())

	//Request ID
	app.Use(requestid.New())

	err = router.SetupRoutes(app)

	if err != nil {
		log.Fatalf("Cannot setup routes")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8081"
	} else {
		port = fmt.Sprintf(":%s", port)
	}

	// Starting cron job
	cron.InitCron()

	log.Fatal(app.Listen(port))
}
