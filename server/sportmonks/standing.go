package sportmonks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Standing struct {
	Position int    `json:"position"`
	Played   int    `json:"played"`
	TeamId   int    `json:"team_id"`
	TeamName string `json:"team_name"`
	TeamLogo string `json:"team_logo"`
	Goals    string `json:"goals"`
	GoalDiff int    `json:"goal_diff"`
	Wins     int    `json:"wins"`
	Lost     int    `json:"lost"`
	Draws    int    `json:"draws"`
	Points   int    `json:"points"`
}

// GetStanding func for get rank
func GetStanding(saisonID int) ([]Standing, error) {
	uri := "standings/season/live/" + strconv.Itoa(saisonID)
	response, err := getAnythingArray(uri)

	if err != nil {
		return []Standing{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var standing []Standing
	jsonEncodedStanding, err := json.Marshal(response.Data)

	if err != nil {
		return []Standing{}, fmt.Errorf("Error marshalling the standing: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedStanding, &standing); err != nil {
		return []Standing{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return standing, nil

}
