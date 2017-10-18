package sportmonks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Team : équipe
type Team struct {
	// ID of team
	ID int `json:"id"`
	// LegacyID : unused
	LegacyID int `json:"legacy_id"`
	// Twitter : twitter account
	Twitter string `json:"twitter"`
	// Name : team's name
	Name string `json:"name"`
	// CountryID :  ID of the team's country
	CountryID int `json:"country_id"`
	// NationalTeam : is it a national team
	NationalTeam bool `json:"national_team"`
	// Founded : date of team creation
	Founded int `json:"founded"`
	// LogoPath : url of the team's logo
	LogoPath string `json:"logo_path"`
	// VenueID : match ID
	VenueID int `json:"venue_id"`
}

// GetTeam : retrieves a team from the sportmonks API
func GetTeam(teamID int) (Team, error) {
	uri := "teams/" + strconv.Itoa(teamID)

	response, err := getAnything(uri)

	if err != nil {
		return Team{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var team Team

	// Marshal the data part in order to decode it from JSON later
	jsonEncodedTeam, err := json.Marshal(response.Data)

	if err != nil {
		return Team{}, fmt.Errorf("Error marshalling the team: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedTeam, &team); err != nil {
		return Team{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return team, nil
}

// GetTeamRoutine : Routine for concurrently getting teams
func GetTeamRoutine(teamID int, teamChannel chan Team) {
	uri := "teams/" + strconv.Itoa(teamID)

	response, err := getAnything(uri)

	if err != nil {
		fmt.Println("Error getting the data: ", err)
		return
	}

	var team Team

	// Marshal the data part in order to decode it from JSON later
	jsonEncodedTeam, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println("Error marshalling the team: ", err)
		return
	}

	if err = json.Unmarshal(jsonEncodedTeam, &team); err != nil {
		fmt.Println("Error unmarshalling the response data: ", err)
		return
	}

	teamChannel <- team
}
