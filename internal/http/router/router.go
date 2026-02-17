package router

import (
	"net/http"

	"github.com/amaterasutears/itk/internal/http/handler/wallet"
	"github.com/labstack/echo/v4"
)

type Router struct {
	w *wallet.Handler
}

func New(w *wallet.Handler) *Router {
	return &Router{
		w: w,
	}
}

func (r *Router) Register(g *echo.Group) {
	// ping
	g.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	// api v1
	apiv1g := g.Group("/api/v1")
	// wallets
	apiv1g.POST("/wallets", r.w.Create)
}
