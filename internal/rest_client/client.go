package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type RestClient struct {
	client *resty.Client
}

// Конструктор для REST клиента
func NewRestClient(addres, port string) *RestClient {
	// Строим адрес для REST клиента
	baseURL := "http://" + addres + ":" + port

	// Создаем новый REST клиент по адресу baseURL
	client := resty.New().SetBaseURL(baseURL).
		SetRetryCount(3).
		SetHeader("Content-type", "application/json").
		OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
			logrus.Printf("Request: %s %s", req.Method, req.URL)
			return nil
		}).
		OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
			logrus.Printf("Response: %d %s", resp.StatusCode(), resp.String())
			return nil
		})

	return &RestClient{client: client}
}

// Структура данных для записи деталей песни
type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// Функция для получения деталей песни со внешнего API
func (rc *RestClient) GetSongDetails(group, song string) (*SongDetail, error) {
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
