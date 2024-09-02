package pg

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/storage"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
)

type partnerStorage struct {
	logger loggerx.Logger
}

func NewPartnerStorage(logger loggerx.Logger) storage.PartnerStorage {
	return &partnerStorage{
		logger: logger,
	}
}

func (p partnerStorage) GetPartnerByID(ctx context.Context, id int64) (*storage.PartnerEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (p partnerStorage) GetPartners(ctx context.Context) ([]*storage.PartnerEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (p partnerStorage) GetPartnerByBIN(ctx context.Context, bin string) (*storage.PartnerEntity, error) {
	//TODO implement me
	panic("implement me")
}
