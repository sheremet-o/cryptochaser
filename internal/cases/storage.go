package cases

import (
	"context"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

//go:generate mockgen -source=./storage.go -destination=./testdata/storage.go --package=testdata
type Storage interface {
	GetCurrent(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMax(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetMin(ctx context.Context, titles []string) ([]*entities.Coin, error)
	GetAvg(ctx context.Context, titles []string) ([]*entities.Coin, error)
	StoreCoin(ctx context.Context, coin *entities.Coin) error
	GetAllTitles(ctx context.Context) ([]string, error)
}
