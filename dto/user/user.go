package userdto

type CreateUser struct {
	ID      int    `json:"no_rekening"`
	Nama    string `json:"nama" form:"nama" validate:"required"`
	Nik     string `json:"nik" form:"nik" validate:"required"`
	NoHp    string `json:"no_hp" form:"no_hp" validate:"required"`
	Nominal int    `json:"nominal" form:"nominal"`
}

type UpdateUser struct {
	ID      int    `json:"no_rekening" form:"no_rekening"`
	Nama    string `json:"nama" form:"nama"`
	Nik     string `json:"nik" form:"nik"`
	NoHp    string `json:"no_hp" form:"no_hp"`
	Nominal int    `json:"nominal" form:"nominal"`
}

type UserResponse struct {
	ID      int    `json:"no_rekening"`
	Nama    string `json:"-"`
	Nik     string `json:"-"`
	NoHp    string `json:"-"`
	Nominal int    `json:"-"`
}

type SaldoResponse struct {
	Nominal int `json:"saldo_terbaru"`
}
