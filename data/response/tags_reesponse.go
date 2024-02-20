package response

type UsersResponse struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	Role_Id      uint   `json:"role_id"`
}
