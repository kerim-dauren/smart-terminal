package storage

import (
	"time"
)

type DeviceEntity struct {
	ID                int64        `json:"id"`
	PartnerID         int64        `json:"partner_id"`
	TariffID          int64        `json:"tariff_id"`
	IMEI              string       `json:"imei"`
	Address           string       `json:"address"`
	Status            DeviceStatus `json:"status"`
	StatusDescription *string      `json:"status_description,omitempty"`
}

type DeviceStatus string

const (
	Active      DeviceStatus = "ACTIVE"
	Preparation DeviceStatus = "PREPARATION"
	Inactive    DeviceStatus = "INACTIVE"
)

type PartnerEntity struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	BIN   string   `json:"bin"`
	Phone string   `json:"phone"`
	Email []string `json:"email"`
}

type TariffEntity struct {
	ID      int64  `json:"id"`
	Tariff  string `json:"tariff"`
	Divider int    `json:"divider"`
}

type TransactionEntity struct {
	TxnID      string    `json:"txn_id"`
	TxnDate    time.Time `json:"txn_date"`
	KspTxnID   int64     `json:"ksp_txn_id"`
	KspTxnDate time.Time `json:"ksp_txn_date"`
	IMEI       string    `json:"imei"`
	BIN        string    `json:"bin"`
	Sum        *float64  `json:"sum"`
	Address    string    `json:"address"`
}

type TransactionStat struct {
	IMEI    string   `json:"imei"`
	Address string   `json:"address"`
	Sum     *float64 `json:"sum"`
}
