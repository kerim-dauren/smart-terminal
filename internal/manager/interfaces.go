package manager

import "context"

type (
	KaspiPaymentManager interface {
		Process(ctx context.Context, request *KaspiPaymentRequest, resultChan chan<- KaspiPaymentResponse, errorChan chan<- error)
	}
)
