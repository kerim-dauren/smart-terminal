package service

import "context"

type checkCommand struct {
}

func NewCheckCommand() KaspiCommand {
	return &checkCommand{}
}

func (c checkCommand) Execute(ctx context.Context, request *KaspiPaymentRequest, device *DeviceDto) (*KaspiPaymentResponse, error) {
	panic("implement me")
}

type payCommand struct {
}

func NewPayCommand() KaspiCommand {
	return &payCommand{}
}

func (p payCommand) Execute(ctx context.Context, request *KaspiPaymentRequest, device *DeviceDto) (*KaspiPaymentResponse, error) {
	//TODO implement me
	panic("implement me")
}
