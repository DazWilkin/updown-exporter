package updown

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	body, err := io.ReadAll(resp.Body)
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

	// log.Info("Result",
	// 	"checks", checks,
	// )

	return checks, nil
}

// GetCheckMetrics is a method that returns a list of updown metrics for a check
// The method implemnents updown's /api/checks/:token/metrics
// See: https://updown.io/api#GET-/api/checks/:token/metrics
func (c *Client) GetCheckMetrics(token string) (Metrics, error) {
	log := c.Log.WithName("GetCheckMetrics")

	if token == "" {
		msg := "method requires a valid Check token"
		log.Info(msg)
		return Metrics{}, errors.New(msg)
	}

	url := fmt.Sprintf("%s/%s/%s/metrics?api-key=%s", root, "api/checks", token, c.APIKey)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Info("Unable to create new request")
		return Metrics{}, err
	}

	rqst.Header.Add("Content-Type", "application/json")

	resp, err := c.Client.Do(rqst)
	if err != nil {
		log.Info("Unable to do request")
		return Metrics{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Info("Unable to read response body")
		return Metrics{}, err
	}

	// log.Info("Body",
	// 	"body", string(body),
	// )

	metrics := Metrics{}
	if err := json.Unmarshal(body, &metrics); err != nil {
		return Metrics{}, err
	}

	// log.Info("Result",
	// 	"metrics", metrics,
	// )

	return metrics, nil
}
