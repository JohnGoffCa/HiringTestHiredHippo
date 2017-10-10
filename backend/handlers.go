package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

var entries []*Applicant

var numWinners int

func listEntrants(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	for _, v := range entries {
		e.Encode(v)
	}
}

func addEntrant(w http.ResponseWriter, r *http.Request) {
	var newEntrant Applicant
	d := json.NewDecoder(r.Body)
	err := d.Decode(&newEntrant)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, `{"applicant_id":"%s","success":true}`, newEntrant.id)
}

func updateEntrant(w http.ResponseWriter, r *http.Request) {
}

func deleteEntrant(w http.ResponseWriter, r *http.Request) {
}

func listEntrant(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusNotFound)
		return
	}
	e := json.NewEncoder(w)
	e.Encode(entrant)
}

func entrantHasWon(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusNotFound)
		return
	}
	entrant.won = didWin()
	status := "Lost"
	if entrant.won {
		status = "Won"
	}
	fmt.Fprintf(w, `{"status":"`+status+`","success":true}`)
}

func didWin() bool {
	if rand.Intn(100) == 1 && numWinners < 5 {
		numWinners++
		return true
	}
	return false
}

func findEntrantByID(idString string) (*Applicant, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return &Applicant{}, errors.New("ID must be an integer")
	}
	for _, v := range entries {
		if v.id == id {
			return v, nil
		}
	}
	return &Applicant{}, errors.New("could not find entrant with that ID")
}
