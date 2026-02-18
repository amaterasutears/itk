-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  wallet_id UUID NOT NULL REFERENCES wallets (id) ON DELETE RESTRICT,
  operation_type VARCHAR(20) NOT NULL CHECK (operation_type IN ('deposit', 'withdraw')),
  amount INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS transactions;
-- +goose StatementEnd
