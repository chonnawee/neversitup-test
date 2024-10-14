package handlers

import (
	"net/http"
	services_interfaces "neversitup-test/internal/core/interfaces/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PermutationsHandler struct {
	service services_interfaces.PermutationsServiceInterface
}

func NewPermutationsHandler(service services_interfaces.PermutationsServiceInterface) PermutationsHandler {
	return PermutationsHandler{service: service}
}

func (h *PermutationsHandler) Generate(c *fiber.Ctx) error {
	input := c.Query("alphabet")
	input = strings.TrimSpace(input)
	answer := h.service.Generate(input)
	response := map[string]interface{}{
		"status": "ok",
		"answer": answer,
	}
	return c.Status(http.StatusOK).JSON(response)
}
