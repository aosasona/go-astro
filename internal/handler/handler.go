package handler

import (
	"go-astro/ui"
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
	r := h.app.Group("/api")

	r.Get("/health", h.checkHealth)
}

func (h *Handler) ServeUI() {
	h.app.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(ui.UI),
		Browse:     false,
		PathPrefix: "dist",
	}))
}
