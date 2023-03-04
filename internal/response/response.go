package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	ctx *fiber.Ctx
}

type Data struct {
	Ok         bool           `json:"ok"`
	StatusCode int            `json:"status_code"`
	Message    string         `json:"message"`
	Data       any            `json:"data,omitempty"`
	Meta       map[string]any `json:"meta,omitempty"`
	Err        error          `json:"-"`
}

func New(c *fiber.Ctx) *Response {
	return &Response{
		ctx: c,
	}
}

func (r *Response) Success(data *Data) error {
	data.Ok = true

	if data.StatusCode == 0 {
		data.StatusCode = 200
	}

	return r.ctx.Status(data.StatusCode).JSON(data)
}

func (r *Response) Error(data *Data) error {
	data.Ok = false

	if data.Message == "" {
		data.Message = "An error occurred."
	}

	if data.StatusCode == 0 {
		data.StatusCode = 500
	}

	if data.Err != nil {
		data.Meta = map[string]interface{}{
			"trace_msg": data.Err.Error(),
		}
	}

	return r.ctx.Status(data.StatusCode).JSON(data)
}
