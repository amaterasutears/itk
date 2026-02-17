package wallet

import (
	"context"
	"time"

	"github.com/amaterasutears/itk/internal/model/wallet"
	"github.com/google/uuid"
)

type Service struct {
	ws WalletStorage
}

func New(ws WalletStorage) *Service {
	return &Service{
		ws: ws,
	}
}

func (s *Service) Create(ctx context.Context) (*wallet.Wallet, error) {
	now := time.Now().UTC()

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return s.ws.Insert(ctx, wallet.New(uuid, now))
}
