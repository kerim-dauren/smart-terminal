package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kerim-dauren/smart-terminal/internal/manager"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PaymentController struct {
	kaspiPaymentManager manager.KaspiPaymentManager
	kaspiPayChan        chan manager.KaspiPaymentResponse
	kaspiErrorChan      chan error
}

func NewPaymentController(kaspiPaymentManager manager.KaspiPaymentManager) *PaymentController {
	return &PaymentController{
		kaspiPaymentManager: kaspiPaymentManager,
		kaspiPayChan:        make(chan manager.KaspiPaymentResponse, 1),
		kaspiErrorChan:      make(chan error, 1),
	}
}

func (c *PaymentController) ProcessKaspiRequest(ctx *gin.Context) {
	command := ctx.Query("command")

	transactionID, err := strconv.ParseInt(ctx.Query("txn_id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param 'txn_id'"})
		return
	}

	imei := ctx.Query("account")

	transactionDateStr := ctx.Query("txn_date")
	var transactionDate *time.Time
	if transactionDateStr != "" {
		parsedTime, err := time.Parse("20060102150405", transactionDateStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param 'txn_date'"})
			return
		}
		transactionDate = &parsedTime
	}

	sumStr := ctx.Query("sum")
	var sum *float64 = nil
	if sumStr != "" {
		if !validateSumFormat(sumStr) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Param 'sum' invalid format"})
			return
		}

		parsedSum, err := strconv.ParseFloat(sumStr, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Param 'sum' invalid value"})
			return
		}
		sum = &parsedSum
	}

	request := &manager.KaspiPaymentRequest{
		KaspiTransactionID:   transactionID,
		Command:              command,
		IMEI:                 imei,
		KaspiTransactionDate: transactionDate,
		Sum:                  sum,
	}

	go c.kaspiPaymentManager.Process(ctx.Request.Context(), request, c.kaspiPayChan, c.kaspiErrorChan)

	select {
	case result := <-c.kaspiPayChan:
		ctx.JSON(http.StatusOK, result)
	case err := <-c.kaspiErrorChan:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	case <-ctx.Request.Context().Done():
		ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
	}
}

func validateSumFormat(sumStr string) bool {
	parts := strings.Split(sumStr, ".")
	return len(parts) == 2 && len(parts[1]) == 2
}
