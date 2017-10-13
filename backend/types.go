package main

import (
	"errors"
	"strconv"
	"sync"
)

// Applicant contains id, name, email and phone number of prospective applicants
type Applicant struct {
	id             int
	applicantName  string
	applicantEmail string
	phoneNumber    int
	won            bool
	mu             sync.Mutex
}

// Message is for the raw json to come into, will get saved into an Applicant
type Message struct {
	Name  string `json:"applicant_name"`
	Email string `json:"applicant_email"`
	Phone string `json:"phone_number"`
}

// NewApplicant creates a new applicant and returns a pointer to it
func NewApplicant(name string, email string, phoneNum int) *Applicant {
	return &Applicant{
		applicantName:  name,
		applicantEmail: email,
		phoneNumber:    phoneNum,
	}
}

// CopyMessageToEntrant copies the fields of the Message struct to entrant
func (a *Applicant) CopyMessageToEntrant(m Message) error {
	if m.Phone != "" {
		phoneNumber, err := strconv.Atoi(m.Phone)
		if err != nil {
			return err
		}
		a.SetPhone(phoneNumber)
	}
	if m.Name != "" {
		a.SetName(m.Name)
	}
	if m.Email != "" {
		a.SetEmail(m.Email)
	}
	return nil
}

// ID returns the ID
func (a *Applicant) ID() (int, error) {
	if a.id == 0 {
		return 0, errors.New("No ID")
	}
	return a.id, nil
}

// SetID sets the ID
func (a *Applicant) SetID(newID int) error {
	if a.id != 0 {
		return errors.New("Already has ID")
	}
	a.mu.Lock()
	a.id = newID
	a.mu.Unlock()
	return nil
}

// Name returns the applicantName
func (a *Applicant) Name() (string, error) {
	if a.applicantName == "" {
		return "", errors.New("No name")
	}
	return a.applicantName, nil
}

// SetName sets the applicantName
func (a *Applicant) SetName(newName string) {
	a.mu.Lock()
	a.applicantName = newName
	a.mu.Unlock()
}

// Email returns the applicantEmail
func (a *Applicant) Email() (string, error) {
	if a.applicantEmail == "" {
		return "", errors.New("No email")
	}
	return a.applicantEmail, nil
}

// SetEmail sets the Email
func (a *Applicant) SetEmail(newEmail string) {
	a.mu.Lock()
	a.applicantEmail = newEmail
	a.mu.Unlock()
}

// Phone returns the phoneNumber
func (a *Applicant) Phone() (int, error) {
	if a.phoneNumber == 0 {
		return 0, errors.New("No Phone Number")
	}
	return a.phoneNumber, nil
}

// SetPhone sets the Phone
func (a *Applicant) SetPhone(newPhone int) {
	a.mu.Lock()
	a.phoneNumber = newPhone
	a.mu.Unlock()
}

// Won returns the won status
func (a *Applicant) Won() bool {
	return a.won
}

// SetWon sets the won status
func (a *Applicant) SetWon(newWon bool) {
	a.mu.Lock()
	a.won = newWon
	a.mu.Unlock()
}
