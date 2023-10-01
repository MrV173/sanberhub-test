package transactionsdto

type CreateTransaction struct {
	UserID        int    `json:"user_no_rekening"`
	Nominal       int    `json:"nominal"`
	Waktu         string `json:"waktu"`
	KodeTransaksi string `json:"kode_transaksi"`
}

type TransactionResponse struct {
	UserID        int    `json:"user_no_rekening"`
	Nominal       int    `json:"-"`
	Waktu         string `json:"-"`
	KodeTransaksi string `json:"-"`
}
