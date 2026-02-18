package transaction

import (
	"errors"

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
