package cases

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sheremet-o/cryptochaser/internal/entities"
)

type Service struct {
	storage   Storage
	extClient ExternalClient
}

func NewService(s Storage, ec ExternalClient) (*Service, error) {
	if s == nil || s == Storage(nil) {
		return nil, errors.Wrap(entities.ErrInvalidParam, "storage must not be nil")
	}

	if ec == nil || ec == ExternalClient(nil) {
		return nil, errors.Wrap(entities.ErrInvalidParam, "external client must not be nil")
	}

	return &Service{
		storage:   s,
		extClient: ec,
	}, nil
}

func (s *Service) GetCoinByTitle(ctx context.Context, title string) (*entities.Coin, error) {
	titles := make([]string, 1)
	titles[0] = title

	coins, err := s.storage.GetCurrent(ctx, titles)
	if err != nil && !errors.Is(err, entities.ErrNotFound) {
		return nil, errors.Wrap(err, "GetCurrent")
	}
	if errors.Is(err, entities.ErrNotFound) {
		coins, err = s.extClient.GetActualCoin(ctx, titles)
		if err != nil {
			return nil, errors.Wrap(err, "GetActualCoin")
		}
		err = s.storage.StoreCoin(ctx, coins[0])
		if err != nil {
			return nil, errors.Wrap(err, "StoreCoin")
		}
		return coins[0], nil
	}
	return coins[0], nil
}

func (s *Service) GetCurrentRates(ctx context.Context, titles []string) ([]*entities.Coin, error) {
	coins, err := s.storage.GetCurrent(ctx, titles)
	if err != nil {
		return nil, errors.Wrap(err, "GetCurrentRates")
	}
	return coins, nil
}

func (s *Service) GetMaxRates(ctx context.Context, title string) (*entities.Coin, error) { // применить к остальным методам
	var CoinsList = []string{title}
	coins, err := s.storage.GetMax(ctx, CoinsList)
	if err != nil {
		return nil, errors.Wrap(err, "GetMax")
	}
	return coins[0], nil
}

func (s *Service) GetMinRates(ctx context.Context, title string) (*entities.Coin, error) { // применить к остальным методам
	var CoinsList = []string{title}
	coins, err := s.storage.GetMin(ctx, CoinsList)
	if err != nil {
		return nil, errors.Wrap(err, "GetMin")
	}
	return coins[0], nil
}

func (s *Service) GetAvgRates(ctx context.Context, title string) (*entities.Coin, error) { // применить к остальным методам
	var CoinsList = []string{title}
	coins, err := s.storage.GetAvg(ctx, CoinsList)
	if err != nil {
		return nil, errors.Wrap(err, "GetAvg")
	}
	return coins[0], nil
}
