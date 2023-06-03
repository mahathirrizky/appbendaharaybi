package cashflow

type CashflowInput struct {
	Jumlah    int `json:"jumlah" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	Jenis string `json:"jenis" binding:"required"`
}