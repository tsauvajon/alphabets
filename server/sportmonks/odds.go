package sportmonks

type Odds struct {
	Label      string                 `json:"label"`
	Value      string                 `json:"value"`
	Winning    string                 `json:"winning"`
	Handicap   string                 `json:"handicap"`
	Total      string                 `json:"total"`
	LastUpdate map[string]interface{} `json:"last_update"`
}
