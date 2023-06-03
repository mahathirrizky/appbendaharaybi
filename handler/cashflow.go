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

func NewCashflowrHandler(cashflowService cashflow.Service) *cashflowHandler {
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