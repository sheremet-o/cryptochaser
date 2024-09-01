package entities

import (
	"time"

	"github.com/pkg/errors"
)

type Coin struct {
	Title     string
	Cost      float64
	CreatedAt time.Time
}

func NewCoin(title string, cost float64) (*Coin, error) {
	if title == "" {
		return nil, errors.Wrap(ErrInvalidParam, "'title' field can't be blank")
	}
	if cost <= 0 {
		return nil, errors.Wrap(ErrInvalidParam, "cost value must be positive")
	}

	return &Coin{
		Title: title,
		Cost:  cost,
	}, nil
}

func (c *Coin) SetTime(t time.Time) {
	c.CreatedAt = t
}

func (c *Coin) SetCurrentTime() {
	c.CreatedAt = time.Now()
}
