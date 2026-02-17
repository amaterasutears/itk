package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r *Router) Register(g *echo.Group) {
	// ping
	g.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
