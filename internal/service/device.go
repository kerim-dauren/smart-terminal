package service

import "context"

type deviceService struct {
}

func NewDeviceService() DeviceService {
	return &deviceService{}
}

func (d deviceService) GetDeviceByImei(ctx context.Context, imei string) (*DeviceDto, error) {
	//TODO implement me
	panic("implement me")
}
