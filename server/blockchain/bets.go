package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Bet : Bet
type Bet struct {
	// assetId of bet
	AssetID string `json:"assetId"`
	// userId
	UserID string `json:"userId"`
	// eventId
	EventID int `json:"eventId"`
	// choice
	Choice int `json:"choice"`
	// bet
	Bet int `json:"bet"`
	// dateBet
	DateBet time.Time `json:"dateBet"`
	// paid
	Paid bool `json:"paid"`
}

// GetBet : Retrieves a bet from the sportmonks API
func GetBet(betID int) (Bet, error) {
	uri := "world.alphabets.Bet/" + strconv.Itoa(betID)

	response, err := getBcAnything(uri)

	if err != nil {
		return Bet{}, fmt.Errorf("Error getting the data: %v", err)
	}

	var bet Bet

	// Marshal the data part in order to decode it from JSON later
	jsonEncodedBet, err := json.Marshal(response.Data)

	if err != nil {
		return Bet{}, fmt.Errorf("Error marshalling the bet: %v", err)
	}

	if err = json.Unmarshal(jsonEncodedBet, &bet); err != nil {
		return Bet{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return bet, nil
}

// GetBets : Get all the bets
func GetBets() ([]Bet, error) {
	client := &http.Client{}
	uri := "world.alphabets.Bet"

	url := apiURI + uri

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return []Bet{}, fmt.Errorf("Error creating the request: %v", err)
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		return []Bet{}, fmt.Errorf("Error executing the request: %v", err)
	}

	defer res.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Bet{}, fmt.Errorf("Error reading response: %v", err)
	}

	var bets []Bet

	if err = json.Unmarshal(body, &bets); err != nil {
		return []Bet{}, fmt.Errorf("Error unmarshalling the response data: %v", err)
	}

	return bets, nil
}

// GetBetsByUserID : idUser => "userId:9230"
func GetBetsByUserID(idUser string) ([]Bet, error) {
	betsFiltered := make([]Bet, 0)

	var bets []Bet
	bets, err := GetBets()
	if err != nil {
		return []Bet{}, fmt.Errorf("Error getting Bets: %v", err)
	}

	for _, bt := range bets {
		if bt.UserID == idUser {
			betsFiltered = append(betsFiltered, bt)
		}
	}

	return bets, nil
}
