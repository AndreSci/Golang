package coincap

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can't be zero")
	}

	// Нечитабельный стиль кода....
	// Сложно увидеть где второй параметр возвращается....
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			// Определяем наш middleware для logger
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {
	// respons, err := client.Get("https://api.bybit.com/v5/market/tickers")
	respons, err := c.client.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return nil, err
	}

	defer respons.Body.Close()

	body, errB := io.ReadAll(respons.Body)
	if errB != nil {
		return nil, errB
	}

	var data assetResponse

	errJson := json.Unmarshal(body, &data.Data)

	if errJson != nil {
		return nil, errJson
	}

	return data.Data, nil

}
