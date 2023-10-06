package cashflow

type CashflowInput struct {
	Jumlah     int    `json:"jumlah" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	Jenis      string `json:"jenis" binding:"required"`
}

type CashflowEditInput struct {
	ID         int    `json:"id" binding:"required"`
	Jumlah     int    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	ImageUrl string `json:"image_url"`
}

type CashflowDeleteInput struct {
	ID int `json:"id" binding:"required"`
	ImageUrl string `json:"image_url"`
}