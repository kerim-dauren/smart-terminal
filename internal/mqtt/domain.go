package mqtt

type Action string

const (
	PING   Action = "PING"
	PONG   Action = "PONG"
	CHECK  Action = "CHECK"  // Проверка модуля на готовность принимать оплату
	PAY    Action = "PAY"    // Передача сообщения на оплату
	REBOOT Action = "REBOOT" // Принудительная перезагрузка модуля
)

type Message interface {
	GetMessageID() string
	GetAction() Action
}
