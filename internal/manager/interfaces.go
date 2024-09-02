package manager

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/domain"
)

type (
	KaspiPaymentManager interface {
		Process(ctx context.Context, request *domain.KaspiPaymentRequest, resultChan chan<- *domain.KaspiPaymentResponse)
	}
)
