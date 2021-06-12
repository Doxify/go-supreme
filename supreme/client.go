package supreme

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	StockURL = "https://www.supremenewyork.com/mobile_stock.json"
)

type Client struct {
	BaseURL    string
	Name       string
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(name string) *Client {
	return &Client{
		BaseURL: StockURL,
		Logger:  log.New(os.Stdout, name+" ", log.LstdFlags),
		HTTPClient: &http.Client{
			// TODO: Add more config to the HTTP client.
			Timeout: time.Minute,
		},
	}
}

func (c *Client) makeRequest(req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	// TODO: Add more headers here...

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request was unsuccessful, recieved status: %d", res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
