package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tsauvajon/ws-blockchain/server/sportmonks"
)

func (app *App) getTeam(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	url := fmt.Sprintf("https://soccer.sportmonks.com/api/v2.0/teams/53?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO")

	// Prepare the request
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating the team: ", err)
		return
	}

	// Execute the request
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error executing the GetTeam request: ", err)
		return
	}

	defer res.Body.Close()

	var response sportmonks.Response

	// Read the response
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading the response body", err)
	}

	// Unmarshal the response : { data: { ... }, meta: { ... }}
	if err = json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error unmarshalling the response", err)
	}

	var team sportmonks.Team

	jsonEncodedTeam, err := json.Marshal(response.Data)

	if err != nil {
		fmt.Println(err)
	}

	if err = json.Unmarshal(jsonEncodedTeam, &team); err != nil {
		fmt.Println("Error unmarshalling the response data", err)
	}

	fmt.Println(team)

	respondWithJSON(w, 200, team)
}
