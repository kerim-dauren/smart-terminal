package pg

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/domain"
	"github.com/kerim-dauren/smart-terminal/internal/storage"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
)

type tariffStorage struct {
	logger loggerx.Logger
}

func NewTariffStorage(logger loggerx.Logger) storage.TariffStorage {
	return &tariffStorage{
		logger: logger,
	}
}

func (t tariffStorage) GetTariffByID(ctx context.Context, id int64) (*domain.TariffEntity, error) {
	//TODO implement me
	panic("implement me")
}
