package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ybi.com/appbendaharaybi/cashflow"
	"ybi.com/appbendaharaybi/helper"
)

type cashflowHandler struct {
	cashflowService cashflow.Service
}

func NewCashflowHandler(cashflowService cashflow.Service) *cashflowHandler {
	return &cashflowHandler{cashflowService}
}

func (h *cashflowHandler) GetCashflow(c *gin.Context) {
 cashflows, err:= h.cashflowService.GetCashflow()
 if err != nil {
	response := helper.APIResponse("Error Get Cash Flow", http.StatusBadRequest,"error", nil)
	c.JSON(http.StatusBadRequest,response)
	return
 }
 response := helper.APIResponse("List Cashflow",http.StatusOK,"success", cashflow.FormatCashflows(cashflows))
 c.JSON(http.StatusOK, response)

}

func(h *cashflowHandler) CreateCashflow(c *gin.Context){
	var input cashflow.CashflowInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors:= helper.FormatValidationError(err)
		errorMessage:= gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create cashflow", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newCashflow, err :=  h.cashflowService.CreateCashflow(input)
	if err != nil {
		response := helper.APIResponse("Failed to create cashflow", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("subscribe created", http.StatusCreated, "success", cashflow.FormatCashflow(newCashflow))
	c.JSON(http.StatusCreated, response)
}