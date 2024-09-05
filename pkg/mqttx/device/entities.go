package device

import (
	"encoding/json"
	"fmt"
	"github.com/kerim-dauren/smart-terminal/pkg/mqttx"
)

// MqttDeviceMessage - интерфейс для сообщений устройства MQTT.
type MqttDeviceMessage interface{}

// PingMessage - структура для сообщения типа PING.
type PingMessage struct {
	MessageID string   `json:"messageId"`
	Data      PingData `json:"data"`
}

// PingData - данные для сообщения DPingMessage.
type PingData struct {
	Status TerminalStatus `json:"status"`
}

// CheckMessage - структура для сообщения типа CHECK.
type CheckMessage struct {
	MessageID string `json:"messageId"`
	Status    Status `json:"status"`
}

// PayMessage - структура для сообщения типа PAY.
type PayMessage struct {
	MessageID string  `json:"messageId"`
	Data      PayData `json:"data"`
}

// PayData - данные для сообщения PayMessage.
type PayData struct {
	TransactionID int64  `json:"transactionId"`
	Status        Status `json:"status"`
}

// MqttMessageModel - модель, которая используется для парсинга общего сообщения.
type MqttMessageModel struct {
	Action    mqttx.MqttAction `json:"action"`
	Data      json.RawMessage  `json:"data"`
	MessageID string           `json:"messageId"`
}

// ParseMqttDeviceMessage - функция для парсинга строки в MqttDeviceMessage.
func ParseMqttDeviceMessage(str string) (MqttDeviceMessage, error) {
	var messageModel MqttMessageModel

	if err := json.Unmarshal([]byte(str), &messageModel); err != nil {
		return nil, err
	}

	switch messageModel.Action {
	case mqttx.PING:
		var data PingData
		if err := json.Unmarshal(messageModel.Data, &data); err != nil {
			return nil, err
		}
		return PingMessage{
			MessageID: messageModel.MessageID,
			Data:      data,
		}, nil
	case mqttx.CHECK:
		return CheckMessage{
			MessageID: messageModel.MessageID,
			Status:    SUCCESS, // Замените на соответствующий статус
		}, nil
	case mqttx.PAY:
		var data PayData
		if err := json.Unmarshal(messageModel.Data, &data); err != nil {
			return nil, err
		}
		return PayMessage{
			MessageID: messageModel.MessageID,
			Data:      data,
		}, nil
	default:
		return nil, fmt.Errorf("unknown action")
	}
}

// Status - пример перечисления для статуса сообщения.
type Status string

const (
	SUCCESS Status = "SUCCESS"
	FAILURE Status = "FAILURE"
)

// TerminalStatus - пример перечисления для статуса терминала.
type TerminalStatus string

const (
	READY TerminalStatus = "READY"
	BUSY  TerminalStatus = "BUSY"
)
