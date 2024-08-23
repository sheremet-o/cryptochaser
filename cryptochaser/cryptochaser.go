package cryptochaser

import (
	"errors"
	"time"
)

type Coin struct {
	Title     string
	Cost      float64
	CreatedAt time.Time
}

func NewCoin(title string, cost float64) (*Coin, error) {
	if title == "" {
		return nil, errors.New("название валюты не может быть пустым")
	}
	if cost <= 0 {
		return nil, errors.New("курс валюты должен быть положительным числом")
	}

	return &Coin{
		Title:     title,
		Cost:      cost,
		CreatedAt: time.Now(),
	}, nil
}
