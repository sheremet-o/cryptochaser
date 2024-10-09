package adapters

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

type Client struct {
	client http.Client
}

func NewClient(fsyms, tsyms []string) {
	c := http.Client{}
	URLRaw := "https://min-api.cryptocompare.com/data/pricemulti?fsyms=BTC,ETH&tsyms=USD,EUR"
	queryParams := url.Values{
		"fsyms": fsyms,
		"tsyms": tsyms,
	}

	u, err := url.Parse(URLRaw)
	if err != nil {
		return
	}

	u.RawQuery = queryParams.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	resp, err := c.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var responseData map[string]map[string]float64
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return
	}

	coins := []entities.Coin{}
	for fsym, values := range responseData {
		for tsym, cost := range values {
			coin := entities.Coin{
				Title:     fsym + "/" + tsym,
				Cost:      cost,
				CreatedAt: time.Now(),
			}
			coins = append(coins, coin)
		}
	}

	fmt.Println(coins)
}
