package models

type Transaction struct {
	ID            int    `json:"-"`
	UserID        int    `json:"-"`
	Nominal       int    `json:"nominal"`
	Waktu         string `json:"waktu"`
	KodeTransaksi string `json:"kode_transaksi"`
}

type TransactionResponse struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_no_rekening"`
	Nominal       int    `json:"nominal"`
	Waktu         string `json:"waktu"`
	KodeTransaksi string `json:"kode_transaksi"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
