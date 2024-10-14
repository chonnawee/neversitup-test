package handlers

import (
	"net/http"
	services_interfaces "neversitup-test/internal/core/interfaces/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CounterHandler struct {
	service services_interfaces.CounterServiceInterface
}

func NewCounterHandler(service services_interfaces.CounterServiceInterface) CounterHandler {
	return CounterHandler{service: service}
}

func (h *CounterHandler) CountSmileys(c *fiber.Ctx) error {
	input := c.Query("smileys")
	input = strings.TrimSpace(input)
	var slice []string
	smileys := strings.Split(input, ",")
	slice = append(slice, smileys...)
	answer := h.service.CountSmileyInSlice(slice)
	response := map[string]interface{}{
		"status": "ok",
		"answer": answer,
	}
	return c.Status(http.StatusOK).JSON(response)
}
