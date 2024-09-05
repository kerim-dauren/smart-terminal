package service

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/domain"
)

type (
	DeviceService interface {
		GetDeviceByImei(ctx context.Context, imei string) (*DeviceDto, error)
	}
)

type PaymentCommand interface {
	Execute(ctx context.Context, request *domain.KaspiPaymentRequest, device *DeviceDto) (*domain.KaspiPaymentResponse, error)
}
