package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tsauvajon/ws-blockchain/server/sportmonks"
)

// Team : soccer team
type Team struct {
	// ID of team
	ID int `json:"id"`
	// Twitter : twitter account
	Twitter string `json:"twitter"`
	// Name : team's name
	Name string `json:"name"`
	// Founded : date of team creation
	Founded int `json:"founded"`
	// LogoPath : url of the team's logo
	LogoPath string `json:"logo_path"`
}

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
