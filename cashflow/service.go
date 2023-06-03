package cashflow

type Service interface {
	GetCashflow() ([]CashflowTable, error)
	CreateCashflow(input CashflowInput) (CashflowTable, error)
	UpdateCashflow(input CashflowEditInput) (CashflowTable, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCashflow() ([]CashflowTable, error) {
	cashflow, err := s.repository.GetCashflow()
	if err != nil {
		return cashflow, err
	}
	return cashflow, nil
}

func (s *service) CreateCashflow(input CashflowInput) (CashflowTable, error) {
	cashflow := CashflowTable{
		Jumlah:     input.Jumlah,
		Keterangan: input.Keterangan,
		Jenis:      input.Jenis,
	}

	newCashfllow, err := s.repository.Save(cashflow)
	if err != nil {
		return newCashfllow, err
	}
	return newCashfllow, nil
}

func (s *service) UpdateCashflow(input CashflowEditInput) (CashflowTable, error) {
	cashflow, err := s.repository.FindCashflowbyID(input.ID)
	if err != nil {
		return cashflow, err
	}
	cashflow.Jenis = input.Jenis
	cashflow.Jumlah = input.Jumlah
	cashflow.Keterangan = input.Keterangan

	updatecashflow, err := s.repository.UpdateCashflow(cashflow)
	if err != nil {
		return updatecashflow, err
	}
	return updatecashflow, nil

}
