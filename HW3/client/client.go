package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"
)

//Обьявляем клиента для получения данных с сторонних сайтов
type Client struct {
	client *http.Client
}

// Функция создания клиента связи с сторонними сайтами
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

func (c Client) GetNames() ([]byte, error) {

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

	var dataName []AssetName

	for i := range data.Data {
		dataName = append(dataName, AssetName{Name: data.Data[i].Name})
	}

	result, errJson := json.Marshal(dataName)
	if errJson != nil{
		return nil, errJson
	}

	return result, nil

}

// Получаем голые данные из сайта
func (c Client) GetData() ([]byte, error) {
	
	respons, err := c.client.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return nil, err
	}

	defer respons.Body.Close()

	body, errB := io.ReadAll(respons.Body)
	if errB != nil {
		return nil, errB
	}

	return body, nil

}

