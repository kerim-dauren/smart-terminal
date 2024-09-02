package manager

import "github.com/kerim-dauren/smart-terminal/pkg/loggerx"

type Deps struct {
}

type Managers struct {
	KaspiPaymentManager KaspiPaymentManager
}

func NewManagers(logger loggerx.Logger, deps *Deps) *Managers {
	return &Managers{
		KaspiPaymentManager: NewKaspiPaymentManager(logger, deps.),
	}
}
