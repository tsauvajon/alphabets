package sportmonks

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Bet struct
type Bet struct {
	Name      string        `json:"name"`
	BookMaker ArrayResponse `json:"bookmaker"`
}

// GetBet Odds with fixtureId
func GetBet(fixturesID int) ([]Bet, error) {
	uri := "odds/fixture/" + strconv.Itoa(fixturesID)

	response, err := getAnythingArray(uri)

	if err != nil {
		return []Bet{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var bet []Bet
	jsonEncodedBet, err := json.Marshal(response.Data)

	if err != nil {
		return []Bet{}, fmt.Errorf("Error marshalling the bet: %v", err)
	}
	if err = json.Unmarshal(jsonEncodedBet, &bet); err != nil {
		return []Bet{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return bet, nil
}
