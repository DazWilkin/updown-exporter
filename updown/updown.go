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

// Client is a type that represents an updown (REST) client
type Client struct {
	APIKey string
	Client *http.Client
	Log    logr.Logger
}

// NewClient is a function that returns a new Client
func NewClient(apiKey string, log logr.Logger) *Client {
	return &Client{
		APIKey: apiKey,
		Client: &http.Client{},
		Log:    log,
	}
}

// GetChecks is a method that returns a list of updown checks
// The method implements updown's /api/checks method
// See: https://updown.io/api#GET-/api/checks
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
