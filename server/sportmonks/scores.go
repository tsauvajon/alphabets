package sportmonks

type ScoreTeam struct {
	LocalTeamScore   int `json:"localteam_score"`
	VisitorTeamScore int `json:"visitorteam_score"`

	HtScore string `json:"ht_score"`
	FtScore string `json:"ft_score"`
	EtScore int    `json:"et_score"`
}
