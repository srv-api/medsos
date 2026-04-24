package medsos

import (
	"srv-api/medsos/dto"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) Create(c echo.Context) error {
	var req dto.MedsosRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body",
		})
	}

	medsos, err := h.serviceMedsos.Create(req)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create medsos",
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"status": "success",
		"data":   medsos,
	})
}
