package mqtt

type MessageService interface {
	// PublishWithCallback публикует сообщение с заданным топиком, сообщением и IMEI.
	// callback вызывается, когда сообщение получено.
	PublishWithCallback(topic string, serverMsg Message, imei string, callback func(deviceMsg Message))

	OnMessageReceived(topic string, payload string)
}
