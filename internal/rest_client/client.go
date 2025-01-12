package client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type RestClient struct {
	client *resty.Client
}

func NewRestClient(baseURL string) *RestClient {
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

type SongDetail struct { // Структура данных возвращаемая внешним API
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func (rc *RestClient) GetSongDetails(group, song string) (*SongDetail, error) {
	var songDetail SongDetail

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
