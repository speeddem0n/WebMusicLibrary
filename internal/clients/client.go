//go:generate mockgen -source=client.go -destination=mock/client_mock.go

package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Интерфейс для связи обработчиков с REST клиетом
type RestClient interface {
	GetSongDetails(group, song string) (*SongDetail, error)
}

type httpClient struct {
	client *resty.Client
}

// Конструктор для REST клиента
func NewRestClient() RestClient {

	// Создаем новый REST клиент по адресу baseURL
	client := resty.New().SetBaseURL(config.Conf.ExternalClientUrl).
		SetRetryCount(3).
		SetHeader("Content-type", "application/json").
		OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
			logrus.Debugf("Request: %s %s", req.Method, req.URL)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
			logrus.Debugf("Response: %d %s", resp.StatusCode(), resp.String())
			return nil
		})

	return &httpClient{client: client}
}

// Структура данных для записи деталей песни
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// Функция для получения деталей песни со внешнего API
func (rc *httpClient) GetSongDetails(group, song string) (*SongDetail, error) {
	var songDetail SongDetail

	// Делаем GET запрос по адресу /info
	resp, err := rc.client.R().
		SetQueryParams(map[string]string{
			"group": group,
			"song":  song,
		}).SetResult(&songDetail).
		Get("/info")

	if err != nil {
		return nil, fmt.Errorf("error fetching song details: %v", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("HTTP error: %s", resp.Status())
	}

	return &songDetail, nil
}
