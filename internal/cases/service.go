package cases

import (
	"context"
	"math"

	"cryptochaser\internal\entities\coin.go"
)

type Service struct {
	storage Storage
	extClient ExternalClient
}

func NewService (s Storage, ec ExternalClient) (*Service, error) {
	if s == nil || s == Storage(nil){
		return nil, errors.Wrap(entities.ErrInvalidParam, "storage must not be nil")
	}

	if ec == nil || ec == ExternalClient(nil){
		return nil, errors.Wrap(entities.ErrInvalidParam, "external client must not be nil")
	}

	return &Service{
		storage: s,
		extClient: ec,
	}, nil
}

func (s *Service)GetCurrentRates(ctx context.Context, titles []string)([]*entities.Coin, error){
	coins, err := s.storage.GetCurrent(ctx, titles)
	if err != nil{
		return nil, errors.Wrap(err, "GetCurrentRates")
	}
	return coins.nil
}

func (s *Service) GetMaxRates(ctx context.Context, title string) (*entities.Coin, error) {
    var maxCoin *entities.Coin

    for _, coin := range s.Coins {
        if coin.Title == title {
            if maxCoin == nil || coin.Cost > maxCoin.Cost {
                maxCoin = coin
            }
        }
    }

    if maxCoin == nil {
        return nil, erorrs.Wrap(err, "GetMaxRates")
    }

    return maxCoin, nil
}

func (s *Service) GetMinRates(ctx context.Context, title string) (*entities.Coin, error) {
    var minCoin *entities.Coin

    for _, coin := range s.Coins {
        if coin.Title == title {
            if minCoin == nil || coin.Cost > minCoin.Cost {
                minCoin = coin
            }
        }
    }

    if minCoin == nil {
        return nil, erorrs.Wrap(err, "GetMinRates")
    }

    return minCoin, nil
}

func (s *Service) GetAvgRates(title string) (float64, error) {
    sum := 0.0
    count := 0
    for _, coin := range s.Coins {
        if coin.Title == title {
            sum += coin.Cost
            count++
        }
    }

    if count == 0 {
        return 0.0, errors.Wrap(err, "GetAvgRates")
    }

    return sum / float64(count), nil
}