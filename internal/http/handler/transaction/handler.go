package transaction

import (
	"log"
	"net/http"

	"github.com/amaterasutears/itk/internal/contract/transaction"
	mtranasction "github.com/amaterasutears/itk/internal/model/transaction"
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
	var req transaction.CreateTranasctionRequest

	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = req.Validate()
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.ts.Create(c.Request().Context(), mtranasction.New(req.WalletID, mtranasction.OperationType(req.OperationType), req.Amount))
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
