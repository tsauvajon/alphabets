package sportmonks

// Response : SportMonks soccer API response
type Response struct {
	// Data : string
	Data map[string]interface{} `json:"data"`
	// Number
	Meta map[string]interface{} `json:"meta"`
}

var (
	api   = "https://soccer.sportmonks.com/api/v2.0"
	teams = "/teams"
)
