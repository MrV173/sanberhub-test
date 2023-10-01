package models

type User struct {
	ID           int           `json:"-" gorm:"primaryKey;autoIncrement:false"`
	Nama         string        `json:"-" gorm:"type: varchar(255)"`
	Nik          string        `json:"-" gorm:"type: varchar(255)"`
	NoHp         string        `json:"-" gorm:"type: varchar(255)"`
	Nominal      int           `json:"-"`
	Transactions []Transaction `json:"mutasi" gorm:"foreignKey:UserID"`
}

type UserResponse struct {
	ID           int                 `json:"no_rekening"`
	Nama         string              `json:"-"`
	Nik          string              `json:"-"`
	NoHp         string              `json:"-"`
	Nominal      int                 `json:"nominal"`
	Transactions TransactionResponse `json:"transaction"`
}

func (UserResponse) TableName() string {
	return "users"
}
