package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	ID 			string `json:"id"`
	Username 	string `json:"username"`
	Email 		string `json:"email"`
	Role		string `json:"role"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

type RegisterResponse struct {
	ID 			string `json:"id"`
	Username 	string `json:"username"`
	Email 		string `json:"email"`
	Role		string `json:"role"`
}

type RefreshResponse struct {
	AccessTokenExpiresIn  int64 `json:"access_token_expires_in"`
	RefreshTokenExpiresIn int64 `json:"refresh_token_expires_in"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type MeResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}