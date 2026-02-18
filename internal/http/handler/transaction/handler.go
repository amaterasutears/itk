package transaction

import (
	"net/http"

	"github.com/amaterasutears/itk/internal/contract/transaction"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	ts TransactionService
}

func New(ts TransactionService) *Handler {
	return &Handler{
		ts: ts,
	}
}

func (h *Handler) Create(c echo.Context) error {
	var requestBody transaction.CreateTranasctionRequest

	err := c.Bind(&requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = requestBody.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.ts.Create(c.Request().Context(), requestBody.ToModel())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
