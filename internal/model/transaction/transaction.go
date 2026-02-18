package transaction

import (
	"time"

	"github.com/google/uuid"
)

type OperationType string

const (
	DepositOperationType  OperationType = "deposit"
	WithdrawOperationType OperationType = "withdraw"
)

type Transaction struct {
	WalletID      uuid.UUID     `db:"wallet_id"`
	OperationType OperationType `db:"operation_type"`
	Amount        int           `db:"amount"`
	CreatedAt     time.Time     `db:"created_at"`
}
