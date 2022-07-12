package updown

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-logr/logr"
)

const (
	root string = "https://updown.io"
)

type Client struct {
	APIKey string
	Client *http.Client
	Log    logr.Logger
}

func NewClient(apiKey string, log logr.Logger) *Client {
	return &Client{
		APIKey: apiKey,
		Client: &http.Client{},
		Log:    log,
	}
}
func (c *Client) GetChecks() ([]Check, error) {
	log := c.Log.WithName("GetChecks")

	url := fmt.Sprintf("%s/%s?api-key=%s", root, "api/checks", c.APIKey)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Info("Unable to create new request")
		return []Check{}, err
	}

	rqst.Header.Add("Content-Type", "application/json")

	resp, err := c.Client.Do(rqst)
	if err != nil {
		log.Info("Unable to do request")
		return []Check{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Info("Unable to read response body")
		return []Check{}, err
	}

	// log.Info("Body",
	// 	"body", string(body),
	// )

	checks := []Check{}
	if err := json.Unmarshal(body, &checks); err != nil {
		return []Check{}, err
	}

	return checks, nil
}
