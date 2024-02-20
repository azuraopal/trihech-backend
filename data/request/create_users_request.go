package request

type CreateUsersRequest struct {
	ID           uint   `validate:"required,min=1,max=200" json:"id"`
	Username     string `validate:"required,min=8,max=50" json:"username"`
	PasswordHash string `validate:"required,varchar(255)" json:"password_hash"`
	Email        string `validate:"required,varchar(255)" json:"email"`
	Role_Id      uint `validate:"required" json:"role_id"`
}
