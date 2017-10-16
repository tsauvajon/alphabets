package sportmonks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	client := &http.Client{}

	url := fmt.Sprintf("https://soccer.sportmonks.com/api/v2.0/teams/" + strconv.Itoa(teamID) + "?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO")

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Team{}, fmt.Errorf("Error creating the team: %v", err)
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return Team{}, fmt.Errorf("Error executing the GetTeam request: %v", err)
	}

	defer res.Body.Close()

	var response Response

	// Read the response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return Team{}, fmt.Errorf("Error reading the response body: %v", err)
	}

	// Unmarshal the response : { data: { ... }, meta: { ... }}
	if err = json.Unmarshal(body, &response); err != nil {
		return Team{}, fmt.Errorf("Error unmarshalling the response: %v", err)
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
