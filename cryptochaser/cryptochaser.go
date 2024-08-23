package cryptochaser

import (
	"time"
)

type Coin struct {
	Title     string
	Cost      float64
	CreatedAt time.Time
}
