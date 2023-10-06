package cashflow


type Service interface {
	GetCashflow() ([]CashflowTable, error)
	CreateCashflow(input CashflowInput, path string) (CashflowTable, error)
	UpdateCashflow(input CashflowEditInput, path string) (CashflowTable, error)
	DeleteCashflow(input int) error
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

func (s *service) CreateCashflow(input CashflowInput, path string) (CashflowTable, error) {
		
	cashflow := CashflowTable{
		Jumlah:     input.Jumlah,
		Keterangan: input.Keterangan,
		Jenis:      input.Jenis,
		ImageUrl: path,
	}

	newCashfllow, err := s.repository.Save(cashflow)
	if err != nil {
		return newCashfllow, err
	}
	return newCashfllow, nil
}

func (s *service) UpdateCashflow(input CashflowEditInput, path string) (CashflowTable, error) {
	cashflow, err := s.repository.FindCashflowbyID(input.ID)
	if err != nil {
		return cashflow, err
	}
	cashflow.Jumlah = input.Jumlah
	cashflow.Keterangan = input.Keterangan
	cashflow.ImageUrl = path

	updatecashflow, err := s.repository.UpdateCashflow(cashflow)
	if err != nil {
		return updatecashflow, err
	}
	return updatecashflow, nil

}

func (s *service) DeleteCashflow(input int) error {
	err := s.repository.DeleteCashflow(input)
	if err != nil {
		return err
	}
	return nil
}
