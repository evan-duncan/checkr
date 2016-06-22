package checkr

import "testing"

func TestCandidateCreate(t *testing.T) {
	SetAPIKey(testKey)
	c := Candidate{
		FirstName:           "John",
		MiddleName:          "Alfred",
		LastName:            "Smith",
		Email:               "john.smith@gmail.com",
		Phone:               "5555555555",
		Zipcode:             "90401",
		DOB:                 "1970-01-22",
		SSN:                 "543-43-4645",
		DriverLicenseNumber: "F211165",
		DriverLicenseState:  "CA",
	}
	if err := c.Create(); err != nil {
		t.Error(err)
	}
}

func TestCandidatesIndex(t *testing.T) {
	SetAPIKey(testKey)
	c := Candidates{}
	if err := c.Index(); err != nil {
		t.Error(err)
	}
}

func TestCandidateShow(t *testing.T) {
	SetAPIKey(testKey)
	c := Candidates{}
	if err := c.Index(); err != nil {
		t.Error(err)
	}
	if len(c.Data) > 0 {
		if err := c.Data[0].Show(); err != nil {
			t.Error(err)
		}
	}
}
