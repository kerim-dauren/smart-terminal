package storage

import (
	"context"
	"time"
)

type DeviceStorage interface {
	GetDeviceByImei(ctx context.Context, imei string) (*DeviceEntity, error)
}

type PartnerStorage interface {
	GetPartnerByID(ctx context.Context, id int64) (*PartnerEntity, error)
	GetPartners(ctx context.Context) ([]*PartnerEntity, error)
	GetPartnerByBIN(ctx context.Context, bin string) (*PartnerEntity, error)
}

type TariffStorage interface {
	GetTariffByID(ctx context.Context, id int64) (*TariffEntity, error)
}

type TransactionStorage interface {
	GetTransactionByKspTxnID(ctx context.Context, kspTxnID int64) (*TransactionEntity, error)
	SaveTransaction(ctx context.Context, transaction *TransactionEntity) error
	GetTransactionStatsByPeriod(ctx context.Context, bin string, from time.Time, to time.Time) ([]*TransactionStat, error)
}
