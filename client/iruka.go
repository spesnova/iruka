package iruka

import (
	"encoding/json"
	"net/http"
)

type Client struct {
	// The URL of the iruka API to communicate with.
	URL string
}

func NewClient() *Client {
	url := "http://localhost:3000/api/v1"

	return &Client{
		URL: url,
	}
}

func (c *Client) Get(v interface{}, path string) error {
	req, err := http.NewRequest("GET", c.URL+path, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)
	return err
}
