package updown

type Check struct {
	Token             string   `json:"token"`
	URL               string   `json:"url"`
	Alias             string   `json:"alias"`
	LastStatus        uint16   `json:"last_status"`
	Uptime            float32  `json:"uptime"`
	Down              bool     `json:"down"`
	DownSince         uint     `json:"down_since"`
	Error             string   `json:"error"`
	Period            uint     `json:"period"`
	Enabled           bool     `json:"enabled"`
	Published         bool     `json:"published"`
	DisabledLocations []string `json:"disabled_locations"`
	Recipients        []string `json:"recipients"`
	LastCheckAt       string   `json:"last_check_at"`
	NextCheckAt       string   `json:"next_check_at"`
	SSL               SSL      `json:"ssl"`
}
type SSL struct {
	TestedAt  string `json:"tested_at"`
	ExpiresAt string `json:"expires_at"`
	Valid     bool   `json:"valid"`
	Error     string `json:"error"`
}
