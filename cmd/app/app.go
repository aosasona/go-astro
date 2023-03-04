package app

import (
	"fmt"
	"go-astro/internal/config"
	"go-astro/internal/handler"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type App struct {
	app    *fiber.App
	config config.Config
}

func New(config config.Config) *App {
	fiber := fiber.New(fiber.Config{
		AppName:               config.AppName,
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
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

	app.Use(logger.New())
	app.Use(recover.New())

	h := handler.New(app)
	h.ServeUI()
	h.ServeAPI()

	return app.Listen(fmt.Sprintf("0.0.0.0:%v", a.config.Port))
}

func (a *App) Kill() error {
	if err := a.app.Shutdown(); err != nil {
		return err
	}
	return nil
}
