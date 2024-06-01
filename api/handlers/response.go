package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Response - Response data
type Response struct {
	Ctx    *fiber.Ctx
	Status int
	Msg    string
	Data   fiber.Map
}

// ResponseData - Response data fields
type ResponseData struct {
	Msg  string    `json:"msg,omitempty"`
	Data fiber.Map `json:"data,omitempty"`
}

// Send - Send output status, message and data
func (r Response) Send() error {
	// Create response data
	responseData := ResponseData{
		Msg:  r.Msg,
		Data: r.Data,
	}

	return r.Ctx.Status(r.Status).JSON(responseData)
}
