package main

// Applicant contains id, name, email and phone number of prospective applicants
type Applicant struct {
	id             int
	applicantName  string
	applicantEmail string
	phoneNumber    int
	won            bool
}

func NewApplicant(name string, email string, phoneNum int) *Applicant {
	return &Applicant{
		applicantName:  name,
		applicantEmail: email,
		phoneNumber:    phoneNum,
	}
}
