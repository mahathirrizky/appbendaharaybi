package cashflow

type CashflowFormatter struct {
	ID         int    `json:"id"`
	Jumlah     int    `json:"jumlah"`
	Keterangan string `json:"keterangan"`
	Jenis      string `json:"jenis"`
}

func FormatCashflow(cashflow CashflowTable) CashflowFormatter {
	return CashflowFormatter{
		ID:    cashflow.ID,
		Jumlah:cashflow.Jumlah,
		Keterangan: cashflow.Keterangan,
		Jenis: cashflow.Jenis,
	}
}

func FormatCashflows(cashflows []CashflowTable) []CashflowFormatter{
	cashflowsFormatter := []CashflowFormatter{}
	for _, cashflow := range cashflows{
		cashflowFormatter:= FormatCashflow(cashflow)
		cashflowsFormatter = append(cashflowsFormatter, cashflowFormatter)
	}
	return cashflowsFormatter
}