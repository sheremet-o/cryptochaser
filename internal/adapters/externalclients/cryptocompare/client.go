package adapters

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

type Client struct {
	Client http.Client
}

func NewClient(*Client, error) {
	c := Client{
		Client: http.Client{},
	}

	return &c(), nil
}

func GetActualCoin(fsyms, tsyms []string) {
	c := http.Client{}
	URLRaw := "https://min-api.cryptocompare.com/data/pricemulti?"
	queryParams := url.Values{
		"fsyms": fsyms,
		"USD":   tsyms,
	}

	u, err := url.Parse(URLRaw)
	if err != nil {
		return nil, errors.Wrap(err, "Parse")
	}

	u.RawQuery = queryParams.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "Encode")
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Do")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "ReadAll")
	}

	var responseData map[string]map[string]float64
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal")
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
}
