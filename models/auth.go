package models

// RequestData for authentication request.
type RequestData struct {
	Email    string
	Password string
	Token    string
}

// PayloadResponse for authentication response.
type PayloadResponse struct {
	Status        string
	Message       string
	SessionCookie string
}

// ValidationResponse for authentication validation.
type ValidationResponse struct {
	IsValid bool
	Message string
}

// UserData for authentication request.
type UserData struct {
	Email    string
	Password string
}
