package wallet

import (
	"context"

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
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return s.ws.Insert(ctx, wallet.New(uuid))
}

func (s *Service) Balance(ctx context.Context, wid uuid.UUID) (int, error) {
	return s.ws.Balance(ctx, wid)
}
