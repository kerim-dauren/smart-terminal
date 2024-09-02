package service

import (
	"math/big"
	"time"
)

type DeviceDto struct {
	ID                int64        `json:"id"`
	IMEI              string       `json:"imei"`
	Address           string       `json:"address"`
	Status            DeviceStatus `json:"status"`
	StatusDescription *string      `json:"statusDescription,omitempty"`
	Partner           PartnerDto   `json:"partner"`
	Tariff            TariffDto    `json:"tariff"`
}

type DeviceStatus string

const (
	Active      DeviceStatus = "ACTIVE"
	Preparation DeviceStatus = "PREPARATION"
	Inactive    DeviceStatus = "INACTIVE"
)

type PartnerDto struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	BIN   string   `json:"bin"`
	Phone string   `json:"phone"`
	Email []string `json:"email"`
}

type TariffDto struct {
	ID      int64  `json:"id"`
	Tariff  string `json:"tariff"`
	Divider int    `json:"divider"`
}

type KaspiPaymentRequest struct {
	TransactionID   int64
	Command         string
	IMEI            string
	TransactionDate time.Time
	Sum             float64
}

type KaspiPaymentResponse struct {
	TransactionID         int64
	Result                ApiResult
	ProviderTransactionId string
	Sum                   big.Float
	Comment               string
	Fields                []NameValueModel
}
