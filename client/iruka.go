package iruka

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

const (
	DefaultAPIURL   = "http://localhost:3000"
	DefaultBasePath = "/api/v1-alpha"
)

type Client struct {
	// The URL of the iruka API to communicate with.
	URL string
}

func NewClient(url string) *Client {
	if os.Getenv("IRUKA_API_URL") != "" {
		url = os.Getenv("IRUKA_API_URL")
	}

	if url == "" {
		url = DefaultAPIURL
	}

	return &Client{
		URL: url,
	}
}

func (c *Client) Post(v interface{}, path string, opts interface{}) error {
	j, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL+DefaultBasePath+path, bytes.NewReader(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)
	return err
}

func (c *Client) Delete(path string) error {
	req, err := http.NewRequest("DELETE", c.URL+DefaultBasePath+path, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	return nil
}

func (c *Client) Get(v interface{}, path string) error {
	req, err := http.NewRequest("GET", c.URL+DefaultBasePath+path, nil)
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

func (c *Client) Update(v interface{}, path string, opts interface{}) error {
	j, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", c.URL+DefaultBasePath+path, bytes.NewReader(j))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&v)
	return err
}
