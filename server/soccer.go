package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tsauvajon/ws-blockchain/server/sportmonks"
)

// Odds struct
type Odds struct {
	Local float64 `json:"local"`
	Draw  float64 `json:"draw"`
	Away  float64 `json:"away"`
}

//getTeam method return Team
func (app *App) getTeamByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithJSON(w, 400, "Invalid team ID")
		return
	}

	team, err := sportmonks.GetTeam(id)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	if team.ID == 0 {
		fmt.Println("Not found")
		respondWithJSON(w, 404, "Team not found")
		return
	}

	respondWithJSON(w, 200, team)
}

//getFixtureByDate method return Fixture by date
func (app *App) getFixtureByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	datedeb := vars["deb"]
	datefin := vars["fin"]

	if datedeb == "" {
		respondWithJSON(w, 400, "Invalid  date debut")
		return
	}
	if datefin == "" {
		respondWithJSON(w, 400, "Invalid  date fin")
		return
	}

	fixture, err := sportmonks.GetFixture(datedeb, datefin)
	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, fixture)
}

//getBetByFixture method return odds by fixture id
func (app *App) getBetByFixture(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithJSON(w, 400, "Invalid team ID")
		return
	}

	bet, err := sportmonks.GetBet(id)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, bet)
}

//calculteOdds return moy of all odds
// func calculateOdds(bet sportmonks.Bet) Odds {

// }
