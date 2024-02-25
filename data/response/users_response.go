package response

type UsersResponse struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	Email        string `json:"email"`
	Role_Id      int    `json:"role_id"`
}
