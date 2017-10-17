package sportmonks

type Weather struct {
	Code        string                 `json:"code"`
	Type        string                 `json:"type"`
	Icon        string                 `json:"icon"`
	Temperature map[string]interface{} `json:"temperature"`
	Clouds      string                 `json:"clouds"`
	Humidity    string                 `json:"humidity"`
	Wind        map[string]interface{} `json:"wind"`
}
