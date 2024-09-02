package service

type Services struct {
	DeviceService DeviceService
	KaspiCommands map[string]KaspiCommand
}

func NewServices() *Services {
	kaspiCommands := map[string]KaspiCommand{
		"check": NewCheckCommand(),
		"pay":   NewPayCommand(),
	}
	return &Services{
		DeviceService: NewDeviceService(),
		KaspiCommands: kaspiCommands,
	}
}
