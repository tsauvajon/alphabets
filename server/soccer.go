package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tsauvajon/ws-blockchain/server/sportmonks"
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
		respondWithJSON(w, 500, err)
		return
	}

	if team.ID == 0 {
		respondWithJSON(w, 404, "Team not found")
		return
	}

	respondWithJSON(w, 200, team)
}
