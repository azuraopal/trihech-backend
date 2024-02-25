package request

type CreateUsersRequest struct {
	ID           int    `validate:"required,min=1,max=200" json:"id"`
	Username     string `validate:"required,min=8,max=50" json:"username"`
	PasswordHash string `gorm:"size:255" validate:"required" json:"password_hash"`
	Email        string `gorm:"size:255,unique" validate:"required,email" json:"email"`
	Role_Id      int    `validate:"required" json:"role_id"`
}
