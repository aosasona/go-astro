package handlers

import (
	"go-astro/web"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type Handler struct {
	app *fiber.App
}

func New(app *fiber.App) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) ServeAPI() {
	api := h.app.Group("/api")
	r := api.Group("/v1")

	r.Get("/health", h.checkHealth)
}

func (h *Handler) ServeUI() {
	h.app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(web.UI),
		Browse:     false,
		PathPrefix: "dist",
	}))
}
