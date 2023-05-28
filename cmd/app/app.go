package app

import (
	"fmt"
	"go-astro/configs"
	"go-astro/handlers"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pocketbase/dbx"
)

type App struct {
	app    *fiber.App
	db     *dbx.DB
	config configs.Config
}

func New(config configs.Config, db *dbx.DB) *App {
	fiber := fiber.New(fiber.Config{
		AppName:               config.AppName,
		DisableStartupMessage: true,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
	})
	return &App{
		app:    fiber,
		db:     db,
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

	h := handlers.New(app)
	h.ServeAPI()
	h.ServeUI()

	return app.Listen(fmt.Sprintf("0.0.0.0:%v", a.config.Port))
}

func (a *App) Kill() error {
	if err := a.app.Shutdown(); err != nil {
		return err
	}
	return nil
}
