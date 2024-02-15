package model

type Userdetails struct {
	UserId       int    `json:"user_id" form:"user_id" gorm:"type int; primaryKey; not null;"`
	NamaLengkap  string `json:"nama_lengkap" form:"nama_lengkap" gorm:"type varchar(100); not null;"`
	JenisKelamin string `json:"jenis_kelamin" form:"jenis_kelamin" gorm:"type enum('L', 'P'); not null;"`
	Bio          string `json:"bio" form:"bio" gorm:"type text not null;"`
}
