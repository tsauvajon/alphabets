package sportmonks

import (
	
)

type GoalsTeam struct{
	Id int `json: "id"`
	TeamId int `json: "team_id"`
	Type string `json: "type"`
	FixtureId int`json: "fixture_id"`
	PlayerId int `json: "player_id"`
	PlayerName string `json: "player_name"`
	PlayerAssistId int `json: "player_assist_id"`
	PlayerAssistName string `json: "player_assist_name"`
	Minute int `json: "minute"`
	ExtraMinute int `json: "extra_minute"`
}