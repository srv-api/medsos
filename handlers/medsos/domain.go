package medsos

import (
	s "srv-api/medsos/services/medsos"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Create(c echo.Context) error
}

type domainHandler struct {
	serviceMedsos s.MedsosService
}

func NewMedsosHandler(service s.MedsosService) DomainHandler {
	return &domainHandler{
		serviceMedsos: service,
	}
}
