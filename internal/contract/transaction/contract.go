package transaction

import (
	"errors"

	"github.com/amaterasutears/itk/internal/model/transaction"
	"github.com/google/uuid"
)

var (
	ErrInvalidOperationType error = errors.New("invalid operation type")
	ErrInvalidAmount        error = errors.New("invalid amount")
)

type OperationType string

const (
	DepositOperationType  OperationType = "deposit"
	WithdrawOperationType OperationType = "withdraw"
)

type CreateTranasctionRequest struct {
	WalletID      uuid.UUID     `json:"wallet_id"`
	OperationType OperationType `json:"operation_type"`
	Amount        int           `json:"amount"`
}

func (c *CreateTranasctionRequest) Validate() error {
	err := c.validateOperationType()
	if err != nil {
		return err
	}

	return c.validateAmount()
}

func (c *CreateTranasctionRequest) validateOperationType() error {
	switch c.OperationType {
	case DepositOperationType, WithdrawOperationType:
		return nil
	}

	return ErrInvalidOperationType
}

func (c *CreateTranasctionRequest) validateAmount() error {
	if c.Amount <= 0 {
		return ErrInvalidAmount
	}

	return nil
}

func (c *CreateTranasctionRequest) ToModel() *transaction.Transaction {
	return &transaction.Transaction{
		WalletID:      c.WalletID,
		OperationType: transaction.OperationType(c.OperationType),
		Amount:        c.Amount,
	}
}
