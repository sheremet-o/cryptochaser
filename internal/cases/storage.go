package cases

import (
	"context"
	"math"

	"cryptochaser\internal\entities\coin.go"
)

type Storage interface {
	GetCurrent(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMax(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMin(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetAvg(ctx context.Context, titles []string) (float64, error)
	StoreCoin(coin *entities.Coin) error
}

type Storage struct {
	Coins []*entities.Coin
}

func (s *Storage) GetCurrent(ctx context.Context, titles []string) ([]*entities.Coin, error) {
	coins, err := s.storage.GetCurrent(ctx, titles)
    if err != nil{
        return nil, erors.Wrap(err, "GetCurrent")
    }
}

func (s *Storage) GetMax(ctx context.Context, title string) (*entities.Coin, error) {
    var maxCoin *entities.Coin

    for _, coin := range s.Coins {
        if coin.Title == title {
            if maxCoin == nil || coin.Cost > maxCoin.Cost {
                maxCoin = coin
            }
        }
    }

    if maxCoin == nil {
        return nil, erorrs.Wrap(err, "GetMax")
    }

    return maxCoin, nil
}

func (s *Storage) GetMin(ctx context.Context, title string) (*entities.Coin, error) {
    var minCoin *entities.Coin

    for _, coin := range s.Coins {
        if coin.Title == title {
            if minCoin == nil || coin.Cost > minCoin.Cost {
                minCoin = coin
            }
        }
    }

    if minCoin == nil {
        return nil, erorrs.Wrap(err, "GetMin")
    }

    return minCoin, nil
}

func (s *Storage) GetAvg(title string) (float64, error) {
    sum := 0.0
    count := 0
    for _, coin := range s.Coins {
        if coin.Title == title {
            sum += coin.Cost
            count++
        }
    }

    if count == 0 {
        return 0.0, errors.Wrap(err, "GetAvg")
    }

    return sum / float64(count), nil
}

func (s *Storage) StoreCoin(coin *entities.Coin) error {
    s.Coins = append(s.Coins, coin)
    return nil
}