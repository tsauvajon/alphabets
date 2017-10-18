package sportmonks

import (
	"encoding/json"
	"fmt"
)

// Fixture struct fixture
type Fixture struct {
	Id            int           `json:"id"`
	LeagueId      int           `json:"league_id"`
	SeasonId      int           `json:"season_id"`
	StageId       int           `json:"stage_id"`
	RoundId       int           `json:"round_id"`
	AggregateId   int           `json:"aggregate_id"`
	VenueId       int           `json:"venue_id"`
	RefereeId     int           `json:"referee_id"`
	LocalteamId   int           `json:"localteam_id"`
	VisitorteamId int           `json:"visitorteam_id"`
	WeatherReport Weather       `json:"weather_report"`
	Formations    FormationTeam `json:"formations"`
	Scores        ScoreTeam     `json:"scores"`
	Time          TimeMatch     `json:"time"`
	Coaches       CoachesTeam   `json:"coaches"`
	Standings     Ranking       `json:"standings"`
	Deleted       bool          `json:"deleted"`
	LocalTeam     Team
	VisitorTeam   Team
	// Goals         Response      `json:"goals"`
	// Stats         Response      `json:"stats"`
	// TV            Response      `json:"tvstations"`
}

// GetFixture get Fixture with date deb and date fin
func GetFixture(deb, fin string) ([]Fixture, error) {
	uri := "fixtures/between/" + deb + "/" + fin
	response, err := getAnythingArray(uri)
	if err != nil {
		return []Fixture{}, fmt.Errorf("Error getting the data: %v", err)
	}
	var fixtures []Fixture
	// Marshal the data part in order to decode it from JSON later
	jsonEncodedFixture, err := json.Marshal(response.Data)

	if err != nil {
		return []Fixture{}, fmt.Errorf("Error marshalling the team: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedFixture, &fixtures); err != nil {
		return []Fixture{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	for i, fixture := range fixtures {
		hometeam, _ := GetTeam(fixture.LocalteamId)
		visitorteam, _ := GetTeam(fixture.VisitorteamId)
		// GetBet(fixture.Id)
		fixtures[i].LocalTeam = hometeam
		fixtures[i].VisitorTeam = visitorteam
	}
	return fixtures, nil
}
