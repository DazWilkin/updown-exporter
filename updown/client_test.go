package updown

import (
	stdlog "log"
	"os"
	"testing"

	"github.com/go-logr/stdr"
)

// TestGetChecks tests the Client's GetChecks method
func TestGetChecks(t *testing.T) {
	log := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
	log = log.WithName("TestGetChecks")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal("expected environment to contain API_KEY")
	}

	client := NewClient(apiKey, log)
	_, err := client.GetChecks()
	if err != nil {
		t.Fatal("expected to be able to get Checks")
	}
}

// TestGetCheckMetrics tests the Client's GetCheckMetrics method
func TestGetCheckMetrics(t *testing.T) {
	log := stdr.NewWithOptions(stdlog.New(os.Stderr, "", stdlog.LstdFlags), stdr.Options{LogCaller: stdr.All})
	log = log.WithName("TestGetMetrics")

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal("expected environment to contain API_KEY")
	}

	client := NewClient(apiKey, log)

	// Need some Check token to make Metrics requests
	checks, err := client.GetChecks()
	if err != nil {
		t.Fatal("expected to be able to get Checks")
	}

	if len(checks) == 0 {
		t.Log("unable to test Metrics as there are no Checks")
		return
	}

	for _, check := range checks {
		_, err := client.GetCheckMetrics(check.Token)
		if err != nil {
			t.Fatalf("expected to be able to get Check's (%s) metrics", check.URL)
		}
	}
}
