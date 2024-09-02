package storage

import (
	"github.com/kerim-dauren/smart-terminal/internal/storage/pg"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
)

type Deps struct {
}

type Storages struct {
	DeviceStorage      DeviceStorage
	PartnerStorage     PartnerStorage
	TariffStorage      TariffStorage
	TransactionStorage TransactionStorage
}

func NewStorages(logger loggerx.Logger, deps *Deps) *Storages {
	return &Storages{
		DeviceStorage:      pg.NewDeviceStorage(logger),
		PartnerStorage:     pg.NewPartnerStorage(logger),
		TariffStorage:      pg.NewTariffStorage(logger),
		TransactionStorage: pg.NewTransactionStorage(logger),
	}
}
