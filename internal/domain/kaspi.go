package domain

import (
	"fmt"
	"time"
)

type KaspiPaymentRequest struct {
	TransactionID   int64
	Command         string
	IMEI            string
	TransactionDate time.Time
	Sum             float64
}

type KaspiPaymentResponse struct {
	TransactionID         int64             `json:"txn_id,omitempty"`
	Result                ApiResult         `json:"result"`
	ProviderTransactionId string            `json:"prv_txn_id,omitempty"`
	Sum                   float64           `json:"sum,omitempty"`
	Comment               string            `json:"comment,omitempty"`
	Fields                []*NameValueModel `json:"fields,omitempty"`
}

type NameValueModel struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ApiResult int

const (
	Success             ApiResult = 0 // абонент/счёт/заказ найден и доступен для пополнения/оплаты
	NotFound            ApiResult = 1 // абонент/счёт не найден или заказ не найден, если запрос check был на проверку состояния заказа
	Canceled            ApiResult = 2 // заказ отменен
	AlreadyPaid         ApiResult = 3 // заказ уже оплачен
	Processing          ApiResult = 4 // платеж в обработке
	InternalServerError ApiResult = 5 // Другая ошибка провайдера
)

func (r ApiResult) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", r)), nil
}
