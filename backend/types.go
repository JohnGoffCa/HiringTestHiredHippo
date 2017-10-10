package main

// Applicant contains id, name, email and phone number of prospective applicants
type Applicant struct {
	id             int
	applicantName  string
	applicantEmail string
	phoneNumber    int
	won            bool
}

// Message is for the raw json to come into, will get saved into an Applicant
type Message struct {
	Name  string `json:"applicant_name"`
	Email string `json:"applicant_email"`
	Phone int    `json:"phone_number"`
}

func NewApplicant(name string, email string, phoneNum int) *Applicant {
	return &Applicant{
		applicantName:  name,
		applicantEmail: email,
		phoneNumber:    phoneNum,
	}
}
