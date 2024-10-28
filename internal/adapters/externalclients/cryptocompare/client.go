package cryptocompare

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/sheremet-o/cryptochaser/internal/entities"
)

type Client struct {
	Client http.Client
}

func NewClient() (*Client, error) {
	c := Client{
		Client: http.Client{},
	}

	return &c, nil
}

func (c *Client) GetActualCoin(ctx context.Context, titles []string) ([]*entities.Coin, error) {
	URLRaw := "https://min-api.cryptocompare.com/data/pricemulti?tsyms=USD"
	queryParams := url.Values{
		"fsyms": titles,
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

	return []*entities.Coin{}, nil
}
