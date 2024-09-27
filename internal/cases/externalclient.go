package cases

import (
	"context"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

//go:generate mockgen -source=./externalclient.go -destination=./testdata/externalclient.go --package=testdata
type ExternalClient interface {
	GetActualCoin(ctx context.Context, titles []string) ([]*entities.Coin, error)
}
