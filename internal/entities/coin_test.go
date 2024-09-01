package entities_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"D:\Dev\_go\cryptochaser\cryptochaser\internal\entities\coin.go"
)

func TestNewCoinCorrect(t *testing.T) {
	coin, err := NewCoin("Testcoin", 10.0)
	assert.NoError(t, err)
	assert.Equal(t, coin.Cost, 10.0)
}

func TestNewCoinEmptyTitle(t *testing.T) {
	_, err := NewCoin("", 10.0)
	assert.Error(t, err)
}

func TestNewCoinNegativeCost(t *testing.T) {
	_, err := NewCoin("Testcoin", -10.0)
	assert.Error(t, err)
}

func TestNewCoinZeroCost(t *testing.T) {
	_, err := NewCoin("Testcoin", 0)
	assert.Error(t, err)
}

func TestSetTime(t *testing.T) {
	coin := &Coin{}
	mockTime := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
	coin.SetTime(mockTime)
	assert.Equal(t, coin.CreatedAt, mockTime)
}

func TestSetCurrentTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockTime := time.Date(2024, time.September, 1, 0, 0, 0, 0, time.UTC)
	timeMock := NewMockTime(ctrl)

	coin := &Coin{}
	SetTimeFunc(timeMock)
	c.SetCurrentTime()

	assert.Equal(t, coin.CreatedAt, mockTime)
}