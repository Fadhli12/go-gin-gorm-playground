package auth

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         UserResponse
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
