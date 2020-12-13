package models

import (
	"testing"
)

func Test_IsEmailValid_True(t *testing.T) {
	result := IsEmailValid("car@motor.com")
	if result == false {
		t.Errorf("IsEmailValid(\"car@motor.com\") expected result to be true but it was %v", result)
	}
}

func Test_IsEmailValid_False(t *testing.T) {
	result := IsEmailValid("car")
	if result == true {
		t.Errorf("IsEmailValid(\"car\") expected result to be false but it was %v", result)
	}
}

func Test_IsPasswordValid_True(t *testing.T) {
	result := IsPasswordValid("abc1237p")

	if result == false {
		t.Errorf("IsPasswordValid(\"abc1237p\") expected result to be true but it was %v", result)
	}
}

func Test_IsPasswordValid_False(t *testing.T) {
	result := IsPasswordValid("abc")

	if result == true {
		t.Errorf("IsPasswordValid(\"abc\") expected result to be false but it was %v", result)
	}
}

func Test_IsOneTimeUseTokenValid_True(t *testing.T) {
	result := IsOneTimeUseTokenValid("1230")

	if result == false {
		t.Errorf("IsOneTimeUseTokenValid(\"1230\") expected result to be true but it was %v", result)
	}
}

func Test_IsOneTimeUseTokenValid_False(t *testing.T) {
	result := IsOneTimeUseTokenValid("5030")

	if result == true {
		t.Errorf("IsOneTimeUseTokenValid(\"5030\") expected result to be true but it was %v", result)
	}
}

func Test_HasValidCredentials_True(t *testing.T) {
	result := HasValidCredentials("c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#")

	if result == false {
		pass := "#th@nH@rm#y#r!$100%D0p#" // %D interference, oops!
		t.Errorf("HasValidCredentials(\"c137@onecause.com\", \"%v\") expected result to be true but it was %v", pass, result)
	}
}

func Test_HasValidCredentials_False(t *testing.T) {
	result := HasValidCredentials("someone@onecause.com", "frogger123")

	if result == true {
		t.Errorf("HasValidCredentials(\"someone@onecause.com\", \"frogger123\") expected result to be true but it was %v", result)
	}
}
