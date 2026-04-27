// routes/routes.go
package routes

import (
	"srv-api/medsos/configs"
	h_medsos "srv-api/medsos/handlers/medsos"
	r_medsos "srv-api/medsos/repositories/medsos"
	s_medsos "srv-api/medsos/services/medsos"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srv-api/middlewares/middlewares"
)

func New() *echo.Echo {
	e := echo.New()
	DB := configs.InitDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	JWT := middlewares.NewJWTService()

	medsosRepo := r_medsos.NewMedsosRepository(DB)
	medsosService := s_medsos.NewMedsosService(medsosRepo, JWT)
	handler := h_medsos.NewMedsosHandler(medsosService)

	e.GET("/picture/*", handler.GetPicture)

	medsos := e.Group("/medsos", middlewares.AuthorizeJWT(JWT))
	{
		medsos.POST("/create", handler.Create)
		medsos.GET("/get", handler.Get)

	}

	return e
}
