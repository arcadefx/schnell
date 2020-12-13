package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arcadefx/schnell/helpers"
	"github.com/arcadefx/schnell/models"
)

// Authenticate handles logic of validation of request
func Authenticate(w http.ResponseWriter, r *http.Request) {
	helpers.SetHeaders(&w)
	helpers.EnableJSON(&w)

	requestData := getAuthRequest(r)
	responseData := models.PayloadResponse{}
	responseValidation := validateRequest(requestData)
	responseCookie := helpers.GetCookie(requestData.Email, requestData.Token, "browser identification")

	if responseValidation.IsValid {
		responseData = models.PayloadResponse{
			Status:        "success",
			Message:       responseValidation.Message,
			SessionCookie: responseCookie,
		}
	} else {
		responseData = models.PayloadResponse{
			Status:        "failed",
			Message:       responseValidation.Message,
			SessionCookie: "",
		}
	}

	json.NewEncoder(w).Encode(responseData)
}

func getAuthRequest(r *http.Request) models.RequestData {
	decoder := json.NewDecoder(r.Body)
	var data models.RequestData
	err := decoder.Decode(&data)
	if err != nil {
		data = models.RequestData{}
	}
	log.Println(data.Email)
	return data
}

func validateRequest(data models.RequestData) models.ValidationResponse {
	isValid := true
	message := "username or password is incorrect"

	if !models.IsEmailValid(data.Email) {
		isValid = false
	}

	if !models.IsPasswordValid(data.Password) {
		isValid = false
	}

	if !models.IsOneTimeUseTokenValid(data.Token) {
		isValid = false
	}

	if !models.HasValidCredentials(data.Email, data.Password) {
		isValid = false
	}

	if isValid {
		message = "user authorized"
	}

	return models.ValidationResponse{
		IsValid: isValid,
		Message: message,
	}
}
