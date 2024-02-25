package request

type UpdateUsersRequest struct {
	ID           int    `validate:"min=1,max=200" json:"id"`
	Username     string `validate:"min=8,max=50" json:"username"`
	PasswordHash string `validate:"varchar(255)" json:"password_hash"`
	Email        string `validate:"varchar(255)" json:"email"`
	Role_Id      int    `validate:"int" json:"role_id"`
}
