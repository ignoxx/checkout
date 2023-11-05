package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct {
	app  *echo.Echo
	port string
}

func NewApi(port string) *Api {
	return &Api{
		port: port,
	}
}

func (a *Api) Start() error {
	a.app = echo.New()

    a.RegisterRoutes()

	return a.app.Start(a.port)
}

func (a *Api) RegisterRoutes() {
    a.app.GET("/ping", func(c echo.Context) error {
        return c.String(http.StatusOK, "pong")
    })
}

func (a *Api) Stop() error {
	context := context.Background()
	return a.app.Shutdown(context)
}
