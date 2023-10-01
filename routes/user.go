package routes

import (
	"sanberhub-test/handlers"
	postgre "sanberhub-test/pkg"
	"sanberhub-test/repository"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	u := repository.RepositoryUser(postgre.DB)
	h := handlers.UserHandler(u)

	e.GET("/users", h.FindUsers)
	e.GET("/mutasi/:no_rekening", h.Mutasi)
	e.GET("/saldo/:no_rekening", h.Saldo)
	e.POST("/daftar", h.Daftar)
	e.POST("/tabung", h.Tabung)
	e.POST("/tarik", h.Tarik)
}
