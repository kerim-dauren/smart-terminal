package service

import (
	"context"
	"github.com/kerim-dauren/smart-terminal/internal/domain"
)

type checkCommand struct {
}

func NewCheckCommand() PaymentCommand {
	return &checkCommand{}
}

func (c checkCommand) Execute(ctx context.Context, request *domain.KaspiPaymentRequest, device *DeviceDto) (*domain.KaspiPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}

type payCommand struct {
}

func NewPayCommand() PaymentCommand {
	return &payCommand{}
}

func (p payCommand) Execute(ctx context.Context, request *domain.KaspiPaymentRequest, device *DeviceDto) (*domain.KaspiPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}
