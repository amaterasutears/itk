package server

import (
	"context"
	"fmt"

	"github.com/amaterasutears/itk/config"
	"github.com/amaterasutears/itk/internal/http/router"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
	c *config.Server
}

func New(r *router.Router, c *config.Server) *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	rootg := e.Group("")

	r.Register(rootg)

	return &Server{
		e: e,
		c: c,
	}
}

func (s *Server) Start() error {
	return s.e.Start(fmt.Sprintf(":%d", s.c.Port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
