package entities_test

import (
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

func TestNewCoin(t *testing.T) {
	tests := []struct {
		name        string
		title       string
		cost        float64
		expectedErr error
	}{
		{name: "Valid parameters", title: "Bitcoin", cost: 5000, expectedErr: nil},
		{name: "Empty title", title: "", cost: 5000, expectedErr: errors.Wrap(entities.ErrInvalidParam, "StoreCoin")},
		{name: "Zero cost", title: "Ethereum", cost: 0, expectedErr: errors.Wrap(entities.ErrInvalidParam, "cost value must be positive")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coin, err := entities.NewCoin(tt.title, tt.cost)

			if err != tt.expectedErr {
				t.Errorf("Expected error to be '%v', but got '%v'", tt.expectedErr, err)
			}

			if err == nil {
				if coin.Title != tt.title {
					t.Errorf("Expected title to be '%s', but got '%s'", tt.title, coin.Title)
				}

				if coin.Cost != tt.cost {
					t.Errorf("Expected cost to be '%f', but got '%f'", tt.cost, coin.Cost)
				}
			}
		})
	}
}

func TestSetTime(t *testing.T) {
	coin := &entities.Coin{}
	tm := time.Now()
	coin.SetTime(tm)

	if !coin.CreatedAt.Equal(tm) {
		t.Errorf("Expected CreatedAt to be '%v', but got '%v'", tm, coin.CreatedAt)
	}
}

func TestSetCurrentTime(t *testing.T) {
	coin := &entities.Coin{}
	coin.SetCurrentTime()
	now := time.Now()

	if now.Sub(coin.CreatedAt) > time.Second {
		t.Errorf("Expected CreatedAt to be current time, but got '%v'", coin.CreatedAt)
	}
}
