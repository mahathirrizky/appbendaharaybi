package cashflow

type Service interface {
	GetCashflow() ([]CashflowTable, error)
	// CreateCashflow(input CashflowInput) ([]CashflowTable, error)
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