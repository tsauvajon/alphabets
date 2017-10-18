package sportmonks

import (
	"encoding/json"
	"fmt"
	"sync"
)

// Fixture struct fixture
type Fixture struct {
	ID            int           `json:"id"`
	LeagueID      int           `json:"league_id"`
	SeasonID      int           `json:"season_id"`
	StageID       int           `json:"stage_id"`
	RoundID       int           `json:"round_id"`
	AggregateID   int           `json:"aggregate_id"`
	VenueID       int           `json:"venue_id"`
	RefereeID     int           `json:"referee_id"`
	LocalTeamID   int           `json:"localteam_id"`
	VisitorTeamID int           `json:"visitorteam_id"`
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

// GetFixtures get Fixture with date deb and date fin
func GetFixtures(deb, fin string) ([]Fixture, error) {
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

	// fixturesChan := make(chan Fixture, len(fixtures))

	updatedFixtures := getTeams(fixtures)

	// for _, fixture := range fixtures {
	// hometeam, _ := GetTeam(fixture.LocalTeamID)
	// visitorteam, _ := GetTeam(fixture.VisitorTeamID)
	// fixtures[i].LocalTeam = hometeam
	// fixtures[i].VisitorTeam = visitorteam
	// }

	// updatedFixtures := make([]Fixture, 0)

	// for i := 0; i < len(fixturesChan); i++ {
	// 	updatedFixtures = append(updatedFixtures, i)
	// }

	// for fixture := range fixturesChan {
	// 	updatedFixtures = append(updatedFixtures, fixture)
	// 	fmt.Println("-")
	// }

	fmt.Print("outside the for: ", updatedFixtures)

	return updatedFixtures, nil
}

func getTeams(fixtures []Fixture) []Fixture {
	var wg sync.WaitGroup
	wg.Add(len(fixtures))

	fixturesChan := make(chan Fixture, len(fixtures))
	for i, fixture := range fixtures {
		go func(fixture Fixture) {
			defer wg.Done()
			// local := make(chan Team, 1)
			// visitor := make(chan Team, 1)

			local, err := GetTeam(fixture.LocalTeamID)

			if err != nil {
				fmt.Println("Impossible to get local team for fixute ", fixture.ID)
				return
			}

			visitor, err := GetTeam(fixture.VisitorTeamID)

			if err != nil {
				fmt.Println("Impossible to get visitor team for fixute ", fixture.ID)
				return
			}

			fixture.LocalTeam = local
			fixture.VisitorTeam = visitor

			// fixtures[i] = fixture
			fixturesChan <- fixture
		}(fixture)
	}

	updatedFixtures := make([]Fixture, 0)

	for fixture := range fixturesChan {
		updatedFixtures = append(updatedFixtures, fixture)
		fmt.Println("-")
	}

	wg.Wait()
	return updatedFixtures
}
