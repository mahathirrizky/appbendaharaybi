package cashflow

import "gorm.io/gorm"

type Repository interface {
	Save(cashflow CashflowTable) (CashflowTable, error)
	GetCashflow() ([]CashflowTable, error)
	UpdateCashflow(cashflow CashflowTable) (CashflowTable, error)
	FindCashflowbyID(cashflowID int)(CashflowTable, error)
	DeleteCashflow(cashflowID int)(error)
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

	err := r.db.Order("created_at DESC").Find(&cashflow).Error
	if err != nil{
		return cashflow, err
	}
	return cashflow, nil
}

func (r *repository) UpdateCashflow(cashflow CashflowTable) (CashflowTable, error) {
	err := r.db.Save(&cashflow).Error
	if err != nil {
		return cashflow, err
	}
	return cashflow, nil
}

func (r *repository)FindCashflowbyID(cashflowID int)(CashflowTable, error){
	var cashflow CashflowTable

	err := r.db.Where("ID = ?", cashflowID).Find(&cashflow).Error
	if err != nil {
		return cashflow, err
	}
	return cashflow, nil
}

func (r *repository) DeleteCashflow(cashflowID int)(error) {
	cashflow := CashflowTable{ID: cashflowID}
	err := r.db.Delete(&cashflow).Error
	if err != nil {
		return err
	}
	return nil
}