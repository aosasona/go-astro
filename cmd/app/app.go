package app

import (
	"go-astro/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type App struct {
	app    *fiber.App
	config config.Config
}

func New(config config.Config) *App {
	fiber := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})
	return &App{
		app:    fiber,
		config: config,
	}
}

func (a *App) Run() error {
	app := a.app

	app.Use(cors.New(cors.Config{
		AllowOrigins: a.config.AllowedOrigins,
	}))

	return nil
}

func (a *App) Kill() error {
	if err := a.app.Shutdown(); err != nil {
		return err
	}
	return nil
}
