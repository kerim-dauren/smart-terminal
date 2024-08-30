package manager

type Deps struct {
}

type Managers struct {
	KaspiPaymentManager KaspiPaymentManager
}

func NewManagers() *Managers {
	return &Managers{}
}
