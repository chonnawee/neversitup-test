package handlers

import (
	"net/http"
	services_interfaces "neversitup-test/internal/core/interfaces/services"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type OddHandler struct {
	service services_interfaces.OddServiceInterface
}

func NewOddHander(service services_interfaces.OddServiceInterface) OddHandler {
	return OddHandler{service: service}
}

func (h *OddHandler) FindIntInSlice(c *fiber.Ctx) error {
	input := c.Query("numbers")
	input = strings.TrimSpace(input)
	numbersStr := strings.Split(input, ",")
	var slice []int
	for _, numberStr := range numbersStr {
		numberInt, _ := strconv.Atoi(numberStr)
		slice = append(slice, numberInt)
	}
	answer := h.service.FindIntInSlice(slice)
	response := map[string]interface{}{
		"status": "ok",
		"answer": answer,
	}
	return c.Status(http.StatusOK).JSON(response)
}
