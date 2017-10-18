package sportmonks

type TimeMatch struct {
	Status      string                 `json:"status"`
	StartingAt  map[string]interface{} `json:"starting_at"`
	Minute      int                    `json:"minute"`
	ExtraMinute int                    `json:"extra_minute"`
	InjuryTime  int                    `json:"injury_time"`
}
