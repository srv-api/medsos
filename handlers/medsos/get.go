package medsos

import (
	"srv-api/medsos/dto"

	"github.com/labstack/echo/v4"
)

func (h *domainHandler) Get(c echo.Context) error {
	var req dto.MatchFeedRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(400, "Invalid request")
	}
	userID := c.Get("UserId").(string)
	req.UserID = userID

	medsos, err := h.serviceMedsos.Get(req)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to get medsos")
	}

	return c.JSON(200, echo.Map{
		"status": "success",
		"data":   medsos,
	})
}
