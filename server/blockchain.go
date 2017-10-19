package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tsauvajon/ws-blockchain/server/blockchain"
)

// connect
func (app *App) connect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	if username == "" {
		respondWithJSON(w, 400, "Invalid empty Username")
		return
	}

	password := vars["password"]
	if password == "" {
		respondWithJSON(w, 400, "Invalid empty Password")
		return
	}

	user, err := blockchain.ConnectUser(username, password)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 401, err)
		return
	}

	respondWithJSON(w, 200, user)
}

// getUsers
func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	user, err := blockchain.GetUsers()

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, user)
}

func (app *App) getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithJSON(w, 400, "Invalid empty ID")
		return
	}

	standing, err := blockchain.GetUser(id)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, standing)
}

func (app *App) getBetsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var idUser = vars["idUser"]
	if idUser == "" {
		respondWithJSON(w, 400, "Invalid empty idUser")
		return
	}

	user, err := blockchain.GetBetsByUserID(idUser)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, user)
}

func (app *App) getBets(w http.ResponseWriter, r *http.Request) {
	user, err := blockchain.GetBets()

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, user)
}

func (app *App) getBetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithJSON(w, 400, "Invalid fixture ID")
		return
	}

	standing, err := blockchain.GetBet(id)

	if err != nil {
		fmt.Println(err)
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, 200, standing)
}
