package wallet

import (
	"errors"
	"net/http"

	"github.com/amaterasutears/itk/internal/contract/wallet"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var (
	ErrWalletIDIsEmpty   error = errors.New("wallet id is empty")
	ErrWalletIDIsInvalid error = errors.New("wallet id is invalid")
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
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, wallet.NewCreateWalletResponse(w))
}

func (h *Handler) Balance(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, ErrWalletIDIsEmpty.Error())
	}

	wid, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, ErrWalletIDIsInvalid.Error())
	}

	b, err := h.ws.Balance(c.Request().Context(), wid)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, wallet.NewBalanceResponse(b))
}
