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

// {
// "token": "ngg8",
// "url": "https://updown.io",
// "alias": "",
// "last_status": 200,
// "uptime": 100,
// "down": false,
// "down_since": null,
// "error": null,
// "period": 15,
// "apdex_t": 0.5,
// "string_match": "",
// "enabled": true,
// "published": true,
// "disabled_locations": [],
// "recipients": ["email:1246848337", "sms:231178295"],
// "last_check_at": "2021-12-17T05:00:01Z",
// "next_check_at": "2021-12-17T05:00:16Z",
// "mute_until": null,
// "favicon_url": "https://updown.io/favicon.png",
// "custom_headers": {},
// "http_verb": "GET/HEAD",
// "http_body": "",
// "ssl": {
// 	"tested_at": "2021-12-17T04:58:04Z",
// 	"expires_at": "2022-02-21T15:57:36Z",
// 	"valid": true,
// 	"error": null
// }
// }
