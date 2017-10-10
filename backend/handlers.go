package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/husobee/vestigo"
)

var entries []Applicant

func listEntrants(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	for _, v := range entries {
		e.Encode(v)
	}
}

func addEntrant(w http.ResponseWriter, r *http.Request) {
}

func updateEntrant(w http.ResponseWriter, r *http.Request) {
}

func deleteEntrant(w http.ResponseWriter, r *http.Request) {
}

func listEntrant(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, err)
	}
	e := json.NewEncoder(w)
	e.Encode(entrant)
}

func entrantHasWon(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		fmt.Fprintf(w, `{"error":"%s"}`, err)
	}
}

func findEntrantByID(id string) (Applicant, error) {
	for _, v := range entries {
		if v.id == id {
			return v, nil
		}
	}
	return Applicant{}, errors.New("could not find entrant with that ID")
}
