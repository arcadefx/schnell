package helpers

import (
	"testing"
)

func Test_GetCookie(t *testing.T) {
	result := GetCookie("car@motor.com", "1230", "192.168.1.111")
	expectedResult := "BgvwzQkFlR4LRlGRJOKlSQrw8ixI2YDnRZsyfsVDFtk="
	if result != expectedResult {
		t.Errorf("Test_GetCookie() expected result to be \"%v\" but it was \"%v\"", expectedResult, result)
	}
}
