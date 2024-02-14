package model

type Userdetails struct {
	UserId       int    `json:"user_id" form:"user_id" gorm:"type int; primaryKey;"`
	NamaLengkap  string `json:"nama_lengkap" form:"nama_lengkap" gorm:"type varchar(100);"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" gorm:"type enum('L', 'P');"`
	Bio          string `json:"bio" form:"bio" gorm:"type text"`
}
