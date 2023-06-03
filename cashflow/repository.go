package cashflow

import "gorm.io/gorm"

type Repository interface {
	Save(cashflow CashflowTable) (CashflowTable, error)
	GetCashflow() ([]CashflowTable, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(cashflow CashflowTable) (CashflowTable, error) {
	err := r.db.Create(&cashflow).Error
	if err != nil {
		return cashflow, err
	}
	return cashflow, nil
}

func (r *repository) GetCashflow() ([]CashflowTable, error){
	var cashflow []CashflowTable

	err := r.db.Find(&cashflow).Error
	if err != nil{
		return cashflow, err
	}
	return cashflow, nil
}