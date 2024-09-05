package mqttx

// MqttAction определяет возможные действия для MQTT сообщений.
type MqttAction string

const (
	PING   MqttAction = "PING"
	PONG   MqttAction = "PONG"
	CHECK  MqttAction = "CHECK"  // Проверка модуля на готовность принимать оплату
	PAY    MqttAction = "PAY"    // Передача сообщения на оплату
	REBOOT MqttAction = "REBOOT" // Принудительная перезагрузка модуля
)
