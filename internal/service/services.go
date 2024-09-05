package service

type Services struct {
	DeviceService DeviceService
	KaspiCommands map[string]PaymentCommand
}

func NewServices() *Services {
	kaspiCommands := map[string]PaymentCommand{
		"check": NewCheckCommand(),
		"pay":   NewPayCommand(),
	}
	return &Services{
		DeviceService: NewDeviceService(),
		KaspiCommands: kaspiCommands,
	}
}
