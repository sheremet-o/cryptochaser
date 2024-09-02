package cases

import (
	"context"
	
	"cryptochaser\internal\entities\coin.go"
)

type ExternalClient interface {
	GetActualCoin(ctx context.Context, titles []string) ([]*entities.Coin, error)
}

func (c *Client)GetActualCoin(ctx context.Context, titles []string)([]*entities.Coin, error){
	coins, err := s.storage.GetCurrent(ctx, titles)
	if err != nil{
		return nil, errors.Wrap(err, "GetActualCoin")
	}
	return coins.nil
}