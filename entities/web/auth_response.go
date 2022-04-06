package web

type AuthResponse struct {
	Token string
	User UserResponse
}