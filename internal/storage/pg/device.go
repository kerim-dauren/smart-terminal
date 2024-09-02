package pg

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/storage"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
)

type deviceStorage struct {
	logger loggerx.Logger
}

func NewDeviceStorage(logger loggerx.Logger) storage.DeviceStorage {
	return &deviceStorage{
		logger: logger,
	}
}

func (s *deviceStorage) GetDeviceByImei(ctx context.Context, imei string) (*storage.DeviceEntity, error) {
	return nil, nil
}
