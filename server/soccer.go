package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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

	url := fmt.Sprintf("https://soccer.sportmonks.com/api/v2.0/teams/53?api_token=QCaGhUsahBbiKEQPGINe7G843Jr5qSsdKmPpUpvR4MJee7xGYw4t63z4YgkO")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	res, err := client.Do(req)

	defer res.Body.Close()

	var (
		response Response
	)
	body, err := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &response)
	// team := Team{
	// 	Id: response.Data["id"]
	// }
	// teamdd := Team{
	// 	Id: 53
	// }

	// fmt.Println(response)

	// team2 := teamnous {
	// 	country: teammonk.country.name
	// }
	team := response.Data

	fmt.Println(team)
	respondWithJSON(w, 200, team)
}
