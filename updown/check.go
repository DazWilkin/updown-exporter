package updown

// Check is a type that represents an updown Check
type Check struct {
	Token             string   `json:"token"`
	URL               string   `json:"url"`
	Alias             string   `json:"alias"`
	LastStatus        uint16   `json:"last_status"`
	Uptime            float32  `json:"uptime"`
	Down              bool     `json:"down"`
	DownSince         string   `json:"down_since"`
	Error             string   `json:"error"`
	Period            uint     `json:"period"`
	ApdexThreshold    float32  `json:"apdex_t"`
	Enabled           bool     `json:"enabled"`
	Published         bool     `json:"published"`
	DisabledLocations []string `json:"disabled_locations"`
	Recipients        []string `json:"recipients"`
	LastCheckAt       string   `json:"last_check_at"`
	NextCheckAt       string   `json:"next_check_at"`
	SSL               SSL      `json:"ssl"`
}

// SSL is a type that represents the SSL subtype of an updown Check
type SSL struct {
	TestedAt  string `json:"tested_at"`
	ExpiresAt string `json:"expires_at"`
	Valid     bool   `json:"valid"`
	Error     string `json:"error"`
}
