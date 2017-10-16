package main

import (
	"net/http"

	"github.com/tsauvajon/ws-blockchain/server/sportmonks"
)

func (app *App) getTeam(w http.ResponseWriter, r *http.Request) {
	team, err := sportmonks.GetTeam(53)

	if err != nil {
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, team)
}
