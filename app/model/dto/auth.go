package dto

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	RefreshToken string
	AccessToken  string
}
