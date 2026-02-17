package wallet

import (
	"log"
	"net/http"

	wc "github.com/amaterasutears/itk/internal/contract/wallet"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	ws WalletService
}

func New(ws WalletService) *Handler {
	return &Handler{
		ws: ws,
	}
}

func (h *Handler) Create(c echo.Context) error {
	w, err := h.ws.Create(c.Request().Context())
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, wc.NewCreateWalletResponse(w.ID, w.CreatedAt))
}
