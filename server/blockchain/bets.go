package blockchain

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"io/ioutil"
	"net/http"
)

// Bet
type Bet struct {
	// assetId of bet
	assetId string `json:"assetId"`
	// userId 
	userId string `json:"userId"`
	// eventId
	eventId int `json:"eventId"`
	// choice
	choice int `json:"choice"`
	// bet 
	bet int `json:"bet"`
	// dateBet 
	dateBet time.Time `json:"dateBet"`
	// paid
	paid bool `json:"paid"`
}


// GetBet : retrieves a bet from the sportmonks API
func GetBet(betId int) (Bet, error) {
	uri := "world.alphabets.Bet/" + strconv.Itoa(betId)

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

// GetBets
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

// GetBets idUser => "userId:9230"
func GetBetsByIdUser(idUser string) ([]Bet, error) {
	betsFiltered := make([]Bet, 0) 

	var bets []Bet
	bets, err := GetBets()
	if err != nil{
		return []Bet{}, fmt.Errorf("Error getting Bets:", err)
	}

	for _, bt := range bets {
		if(bt.userId == idUser){
			betsFiltered = append(betsFiltered, bt)
		}
	}

	return bets, nil
}