package pg

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/storage"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
	"time"
)

type transactionStorage struct {
	logger loggerx.Logger
}

func NewTransactionStorage(logger loggerx.Logger) storage.TransactionStorage {
	return &transactionStorage{
		logger: logger,
	}
}

func (t transactionStorage) GetTransactionByKspTxnID(ctx context.Context, kspTxnID int64) (*storage.TransactionEntity, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionStorage) SaveTransaction(ctx context.Context, transaction *storage.TransactionEntity) error {
	//TODO implement me
	panic("implement me")
}

func (t transactionStorage) GetTransactionStatsByPeriod(ctx context.Context, bin string, from time.Time, to time.Time) ([]*storage.TransactionStat, error) {
	//TODO implement me
	panic("implement me")
}
