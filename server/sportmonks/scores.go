package sportmonks

import (

)

type ScoreTeam struct{
	LocalTeamScore int `json:"localteam_score"`
	VisitorTeamScore int `json:"visitorteam_score"`
	LocalTeamPenScore string `json:"localteam_pen_score"`
	VisitorTeamPenScore string `json:"visitorteam_pen_score"`
	HtScore string `json:"ht_score"`
	FtScore string `json:"ft_score"`
	EtScore int `json:"et_score"`
}