package handler

import (
	"go-astro/internal/response"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) checkHealth(c *fiber.Ctx) error {
	return response.New(c).Success(&response.Data{
		Message: "I'm alive!",
		Data: map[string]any{
			"time": time.Now().String(),
		},
	})
}
