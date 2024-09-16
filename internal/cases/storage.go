package cases

import (
	"context"

	"D:\Dev\_go\cryptochaser\cryptochaser\internal\entities\coin.go"

	"github.com/pkg/errors"
)

type Storage interface {
	GetCurrent(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMax(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMin(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetAvg(ctx context.Context, titles []string) ([]*entities.Coin, error)
	StoreCoin(ctx context.Context, coin *entities.Coin) error
	GetAllTitles(ctx context.Context)([]string, error)
}

type Storage struct {
	Coins []*entities.Coin
}

func (s *Storage) GetCurrent(ctx context.Context, titles []string) ([]*entities.Coin, error) {
	coins, err := s.storage.GetCurrent(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(err, "GetCurrent")
	}
}

func (s *Storage) GetMax(ctx context.Context, title string) (*entities.Coin, error) {
	coins, err := s.storage.GetMax(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(err, "GetMax")
	}
}

func (s *Storage) GetMin(ctx context.Context, title string) (*entities.Coin, error) {
	coins, err := s.storage.GetMin(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(err, "GetMin")
	}
}

func (s *Storage) GetAvg(ctx context.Context, title string) (*entities.Coin, error) {
	coins, err := s.storage.GetAvg(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(err, "GetAvgn")
	}
}

func (s *Storage) StoreCoin(coin *entities.Coin) error {
	s.Coins = append(s.Coins, coin)
	return nil
}
