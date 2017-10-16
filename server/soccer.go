package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {}

type FormationTeam struct{}

type ScoreTeam struct{}

type TimeMatch struct{}

type CoachesTeam struct{}

type Ranking struct{}

type GoalsTeam struct{}

type StatsTeam struct{}

type TvStations struct{}

type Fixture struct {
	Id int `json:"id"`
	LeagueId int `json:"league_id"`
	SeasonId int `json:"season_id"`
	StageId int `json:"stage_id"`
	RoundId int `json:"round_id"`
	AggregateId int `json:"aggregate_id"`
	VenueId int `json:"venue_id"`
	RefereeId int `json:"referee_id"`
	LocalteamId int `json:"localteam_id"`
	VisitorteamId int `json:"visitorteam_id"`
	WeatherReport Weather
	Formations FormationTeam
	Scores ScoreTeam
	Time TimeMatch
	Coaches CoachesTeam
	Standings Ranking
	Deleted bool `json:"deleted"`
	LocalTeam Team
	VisitorTeam Team
	Goals GoalsTeam
	Stats StatsTeam
	TV TvStations

}

 type Team struct {
 	// Id : id of team
 	Id int `json:"id"`
 	// Legacy_id : id of legacy
 	LegacyId int `json:"legacy_id"`
 	// twitter
 	Twitter string `json:"twitter"`
 	// Name : name of team
 	Name string `json:"name"`
 	// country_id : id of team country
 	CountryId int `json:"country_id"`
 	// national_team : yes or no it's an national team
 	NationalTeam bool `json:"national_team"`
 	// founded : date of team
 	Founded int `json:"founded"`
 	// logo_path : logo of team
 	LogoPath string `json:"logo_path"`
 	//
 	VenueId int `json:"venue_id"`
 }

func (app *App) getTeam(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	url := fmt.Sprintf("https://soccer.sportmonks.com/api/v2.0/teams/284?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	res, err := client.Do(req)
	if err != nil{
		log.Fatal("Result: ", err)
		return
	}

	defer res.Body.Close()

	var (
		response Response
	)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal("Body: ", err)
		return
	}

	json.Unmarshal(body, &response)
	team := response.Data
	fmt.Println(team)
	respondWithJSON(w, 200, team)
}

func (app *App) getMatch(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	url := fmt.Sprintf("https://soccer.sportmonks.com/api/v2.0/fixtures/between/2017-10-14/2017-10-15?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO&include=localTeam,visitorTeam,goals,league=Premiership,stats,tvstations")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	res, err := client.Do(req)
	if err != nil{
		log.Fatal("Result: ", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal("Body: ", err)
		return
	}
	fmt.Println(body)
	var (
		response Response
	)
	json.Unmarshal(body, &response)

}