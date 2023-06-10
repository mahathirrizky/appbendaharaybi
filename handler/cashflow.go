package handler

import (
	"encoding/json"
	"time"

	"net/http"
	"net/url"

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
	var isUploaded bool
var path string
	
	err := c.Request.ParseMultipartForm(32 << 20) 
	if err != nil {
		response := helper.APIResponse("Failed to create cashflow", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	jsonData := c.Request.PostFormValue("data")
	var input cashflow.CashflowInput

	err = json.Unmarshal([]byte(jsonData), &input)
	if err != nil {
		errors:= helper.FormatValidationError(err)
		errorMessage:= gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create cashflow", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	file, err := c.FormFile("image")
if err != nil {
    // No image uploaded
    isUploaded = false
} else {
	// Image uploaded
	isUploaded = true

	// Generate the safe file name by escaping special characters
	escapedFileName := url.PathEscape(file.Filename)
	currentDate := time.Now().Format("2006-01-02")
	path = "images/" + currentDate + escapedFileName
	err = c.SaveUploadedFile(file, path)
	if err != nil {
			response := helper.APIResponse("Failed to save image", http.StatusInternalServerError, "error", nil)
			c.JSON(http.StatusInternalServerError, response)
			return
	}
}

newCashflow, err := h.cashflowService.CreateCashflow(input, path)
if err != nil {
    response := helper.APIResponse("Failed to create cashflow", http.StatusBadRequest, "error", nil)
    c.JSON(http.StatusBadRequest, response)
    return
}

data := gin.H{
	"is_uploaded": isUploaded,
	"cashflow":    cashflow.FormatCashflow(newCashflow),
}
statusCode := http.StatusOK
if isUploaded {
	statusCode = http.StatusCreated
}
response := helper.APIResponse("Cashflow created", statusCode, "success", data)
c.JSON(statusCode, response)
}

func(h *cashflowHandler) UpdateCashflow(c *gin.Context){
	var input cashflow.CashflowEditInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to edit cashflow", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatecashflow, err:= h.cashflowService.UpdateCashflow(input)
	if err != nil {
		response := helper.APIResponse("Failed to update materi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Success to update materi", http.StatusOK, "success", cashflow.FormatCashflow(updatecashflow))
	c.JSON(http.StatusOK, response)
}

func (h *cashflowHandler) DeleteCashflow(c *gin.Context) {
	var input cashflow.CashflowDeleteInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to delete cashflow", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = h.cashflowService.DeleteCashflow(input.ID)
	if err != nil {
		response := helper.APIResponse("Failed to delete cashflow", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to delete cashflow", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
