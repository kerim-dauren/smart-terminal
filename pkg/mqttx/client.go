package mqttx

import (
	"github.com/kerim-dauren/smart-terminal/pkg/mqttx/device"
	"github.com/kerim-dauren/smart-terminal/pkg/mqttx/server"
)

type MqttMessageService interface {
	// Publish публикует сообщение с заданным топиком и нагрузкой.
	Publish(topic string, payload string)

	// PublishWithCallback публикует сообщение с заданным топиком, сообщением и IMEI.
	// callback вызывается, когда сообщение получено.
	PublishWithCallback(topic string, message server.MqttServerMessage, imei string, callback func(device.MqttDeviceMessage))

	// OnMessageReceived вызывается при получении сообщения с заданным топиком и нагрузкой.
	OnMessageReceived(topic string, payload string)
}
