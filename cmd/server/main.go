package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/ghv061101/RestApiAge/config"
	handlers "github.com/ghv061101/RestApiAge/internal/handler"
	"github.com/ghv061101/RestApiAge/internal/logger"
	"github.com/ghv061101/RestApiAge/internal/middleware"
	"github.com/ghv061101/RestApiAge/internal/models"
	"github.com/ghv061101/RestApiAge/internal/repository"
	routes "github.com/ghv061101/RestApiAge/internal/routes"
	"github.com/ghv061101/RestApiAge/internal/service"
	"github.com/ghv061101/RestApiAge/storage"
)

func main() {
	cfg := config.Load()

	dbCfg := &storage.Config{URL: cfg.DatabaseURL}
	db, err := storage.NewConnection(dbCfg)
	if err != nil {
		log.Fatalf("could not load the database: %v", err)
	}

	if err := models.MigrateUsers(db); err != nil {
		log.Fatalf("could not migrate db: %v", err)
	}

	repo := repository.New(db)
	svc := service.New(repo)
	h := handlers.NewUserHandler(svc)

	logg, _ := logger.New()
	defer logg.Sync()

	app := fiber.New()
	// middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger(logg))

	routes.Register(app, h)

	log.Println("starting server on :8080")
	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

