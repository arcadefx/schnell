package models

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`(?:[a-z0-9!#$%&'*+\/=?^_\x60{|}~-]+(?:\.[a-z0-9!#$%&'*+\/=?^_\x60{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)
var passwordRegex = regexp.MustCompile(`^[a-zA-Z0-9.!@?#"$%&:';()*\+,\/;\-=[\\\]\^_{|}<>~\x60 ]+$`)
var oneTimeUseTokenRegex = regexp.MustCompile(`^(0[0-9]|1[0-9]|2[0-3])[0-5][0-9]$`)

// IsEmailValid returns true if a valid email is passed.
func IsEmailValid(input string) bool {
	if len(input) > 254 {
		return false
	}
	return emailRegex.MatchString(input)
}

// IsPasswordValid returns true if a valid password is passed
func IsPasswordValid(input string) bool {
	if len(input) < 8 || len(input) > 254 {
		return false
	}
	return passwordRegex.MatchString(input)
}

// IsOneTimeUseTokenValid returns true if a valid token is passed.
func IsOneTimeUseTokenValid(input string) bool {
	return oneTimeUseTokenRegex.MatchString(input)
}

// HasValidCredentials check to see if credentials match
func HasValidCredentials(email string, pass string) bool {
	// pretend to have gotten credentials from NoSQL or a data store
	storedData := UserData{
		Email:    "c137@onecause.com",
		Password: "#th@nH@rm#y#r!$100%D0p#",
	}
	if email == storedData.Email && pass == storedData.Password {
		return true
	}
	return false
}
