package main

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

var entries []*Applicant

var numWinners int

func listEntrants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "[")
	for i, v := range entries {
		if i != 0 {
			fmt.Fprint(w, ",")
		}
		fmt.Fprintf(w, "%d", v.id)
	}
	fmt.Fprint(w, "]")
}

func addEntrant(w http.ResponseWriter, r *http.Request) {
	var newEntrant Applicant
	var m Message
	d := json.NewDecoder(r.Body)
	err := d.Decode(&m)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusBadRequest)
		return
	}

	// check if the name or email already exists
	if checkEntrantExists(m.Name, m.Email) {
		http.Error(w, `{"error":"Entrant by that name or email already exists","success":false}`, http.StatusUnprocessableEntity)
		return
	}

	// copy m to newEntrant and store its address in entries
	// because its address is still in use newEntrant won't be garbage collected
	newEntrant.applicantEmail, newEntrant.applicantName, newEntrant.phoneNumber = m.Email, m.Name, m.Phone
	newID, err := newRandomID()
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusBadRequest)
		return
	}
	newEntrant.id = newID
	entries = append(entries, &newEntrant)
	fmt.Fprintf(w, `{"applicant_id":%d,"success":true}`, newEntrant.id)
}

func updateEntrant(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusNotFound)
		return
	}

	var m Message
	d := json.NewDecoder(r.Body)
	err = d.Decode(&m)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusBadRequest)
		return
	}
	if m.Name != "" {
		entrant.applicantName = m.Name
	}
	if m.Email != "" {
		entrant.applicantEmail = m.Email
	}
	if m.Phone != 0 {
		entrant.phoneNumber = m.Phone
	}
	fmt.Fprintf(
		w,
		`{"applicant_name":"%s","applicant_email":"%s","phone_number":%d,"success":true}`,
		entrant.applicantName,
		entrant.applicantEmail,
		entrant.phoneNumber,
	)
}

func deleteEntrant(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	idx, err := findEntrantIndexByID(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusNotFound)
		return
	}
	entries = append(entries[:idx], entries[idx+1:]...)
	fmt.Fprint(w, `{"success":true}`)
}

func listEntrant(w http.ResponseWriter, r *http.Request) {
	id := vestigo.Param(r, "id")
	entrant, err := findEntrantByID(id)
	if err != nil {
		http.Error(w, `{"error":"`+err.Error()+`","success":false}`, http.StatusNotFound)
		return
	}
	fmt.Fprintf(
		w,
		`{"applicant_name":"%s","applicant_email":"%s","phone_number":%d,"success":true}`,
		entrant.applicantName,
		entrant.applicantEmail,
		entrant.phoneNumber,
	)
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
	if bRand, _ := rand.Int(rand.Reader, big.NewInt(100)); bRand.Int64() == 1 && numWinners < 5 {
		numWinners++
		return true
	}
	return false
}

// generates a random 8 digit number
func newRandomID() (int, error) {
	randBig, err := rand.Int(rand.Reader, big.NewInt(99999999))
	if err != nil {
		return 0, err
	}
	newIDValue := 10000000 + int(randBig.Int64())
	for a, _ := findEntrantByID(string(newIDValue)); a == nil; {
		randBig, err = rand.Int(rand.Reader, big.NewInt(99999999))
		if err != nil {
			return 0, err
		}
		newIDValue = 10000000 + int(randBig.Int64())
	}
	return newIDValue, nil
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

func findEntrantIndexByID(idString string) (int, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, errors.New("ID must be an integer")
	}
	for i, v := range entries {
		if v.id == id {
			return i, nil
		}
	}
	return 0, errors.New("could not find entrant with that ID")
}

func checkEntrantExists(name, email string) bool {
	for _, v := range entries {
		if v.applicantName == name || v.applicantEmail == email {
			return true
		}
	}
	return false
}
