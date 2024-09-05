package server

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/kerim-dauren/smart-terminal/pkg/mqttx"
)

// MqttServerMessage - базовая структура для сообщений MQTT сервера.
type MqttServerMessage interface {
	ToJson() (string, error)
	GetAction() mqttx.MqttAction
	GetMessageID() string
}

// базовая структура для сообщения, реализующая общие методы
type baseMqttServerMessage struct {
	Action    mqttx.MqttAction
	MessageID string
}

// создание нового baseMqttServerMessage
func newBaseMqttServerMessage(action mqttx.MqttAction) baseMqttServerMessage {
	return baseMqttServerMessage{
		Action:    action,
		MessageID: uuid.New().String(),
	}
}

func (b baseMqttServerMessage) GetAction() mqttx.MqttAction {
	return b.Action
}

func (b baseMqttServerMessage) GetMessageID() string {
	return b.MessageID
}

// PingMessage - ответ на запрос PING.
type PingMessage struct {
	baseMqttServerMessage
}

func NewPingMessage(messageID string) *PingMessage {
	return &PingMessage{baseMqttServerMessage{mqttx.PING, messageID}}
}

func (p PingMessage) ToJson() (string, error) {
	data, err := json.Marshal(p)
	return string(data), err
}

// CheckMessage - проверка модуля на готовность принимать оплату.
type CheckMessage struct {
	baseMqttServerMessage
}

func NewCheckMessage() *CheckMessage {
	return &CheckMessage{newBaseMqttServerMessage(mqttx.CHECK)}
}

func (c CheckMessage) ToJson() (string, error) {
	data, err := json.Marshal(c)
	return string(data), err
}

// PayData - данные для сообщения Pay.
type PayData struct {
	TransactionID int64 `json:"transactionId"`
	PulseCount    int   `json:"pulseCount"`
}

// PayMessage - сообщение на оплату.
type PayMessage struct {
	baseMqttServerMessage
	Data PayData `json:"data"`
}

func NewPayMessage(data PayData) *PayMessage {
	return &PayMessage{baseMqttServerMessage: newBaseMqttServerMessage(mqttx.PAY), Data: data}
}

func (p PayMessage) ToJson() (string, error) {
	data, err := json.Marshal(p)
	return string(data), err
}

// ParseMqttServerMessage - парсинг строки в MqttServerMessage.
func ParseMqttServerMessage(str string) (MqttServerMessage, error) {
	var messageModel struct {
		Action    mqttx.MqttAction `json:"action"`
		Data      json.RawMessage  `json:"data"`
		MessageID string           `json:"messageId"`
	}

	if err := json.Unmarshal([]byte(str), &messageModel); err != nil {
		return nil, err
	}

	switch messageModel.Action {
	case mqttx.PING:
		return NewPingMessage(messageModel.MessageID), nil
	case mqttx.CHECK:
		return NewCheckMessage(), nil
	case mqttx.PAY:
		var payData PayData
		if err := json.Unmarshal(messageModel.Data, &payData); err != nil {
			return nil, err
		}
		return NewPayMessage(payData), nil
	default:
		return nil, fmt.Errorf("unknown action")
	}
}
