package service

import "context"

type (
	DeviceService interface {
		GetDeviceByImei(ctx context.Context, imei string) (*DeviceDto, error)
	}
)

type KaspiCommand interface {
	Execute(ctx context.Context, request *KaspiPaymentRequest, device *DeviceDto) (*KaspiPaymentResponse, error)
}
