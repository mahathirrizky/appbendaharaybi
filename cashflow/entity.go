package cashflow

import "time"

type CashflowTable struct {
	ID             int
	Jumlah          int
	Keterangan          string
	Jenis   				string
	ImageUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
