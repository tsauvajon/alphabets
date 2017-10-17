package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/geoffreydalfin/ws-blockchain/server/sportmonks"
	"github.com/gorilla/mux"
)

func (app *App) getTeam(w http.ResponseWriter, r *http.Request) {
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

	// if bet.BookMaker.ID == 0 {
	// 	fmt.Println("Not found")
	// 	respondWithJSON(w, 404, "Bookmaker not found")
	// 	return
	// }

	respondWithJSON(w, 200, bet)
}
