package public_controller

import (
	"github.com/gofiber/fiber/v2"

	"Template/models/errors"
	"Template/models/response"
)

func checkHealth() response.ResponseModel {
	return response.ResponseModel{
		RetCode: "100",
		Message: "Request success!",
		Data: errors.ErrorModel{
			Message:   "Service is available!",
			IsSuccess: true,
			Error:     nil,
		},
	}
}

func CheckServiceHealth(c *fiber.Ctx) error {
	health := checkHealth()
	response := errors.ErrorModel{}
	response = health.Data.(errors.ErrorModel)
	if !response.IsSuccess {
		return c.JSON(health)
	}
	return c.JSON(health)
}
