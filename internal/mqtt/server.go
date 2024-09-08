package mqtt

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type ServerMessage interface {
	Message
}

type baseServerMessage struct {
	MessageID string `json:"messageId"`
	Action    Action `json:"action"`
}

func newBaseServerMessage(action Action) baseServerMessage {
	return baseServerMessage{
		MessageID: uuid.New().String(),
		Action:    action,
	}
}

func (b baseServerMessage) GetMessageID() string {
	return b.MessageID
}

func (b baseServerMessage) GetAction() Action {
	return b.Action
}

type ServerDataPay struct {
	TransactionID int64 `json:"transactionId"`
	PulseCount    int   `json:"pulseCount"`
}

type ServerMessagePay struct {
	baseServerMessage
	Data ServerDataPay
}

func (s ServerMessagePay) GetMessageID() string {
	return s.MessageID
}

func (s ServerMessagePay) GetAction() Action {
	return s.Action
}

type ServerMessagePing struct {
	baseServerMessage
}

func (s ServerMessagePing) GetMessageID() string {
	return s.MessageID
}

func (s ServerMessagePing) GetAction() Action {
	return s.Action
}

type ServerMessageCheck struct {
	baseServerMessage
}

func (s ServerMessageCheck) GetMessageID() string {
	return s.MessageID
}

func (s ServerMessageCheck) GetAction() Action {
	return s.Action
}

func ParseServerMessage(str string) (ServerMessage, error) {
	var messageModel struct {
		Action    Action          `json:"action"`
		Data      json.RawMessage `json:"data"`
		MessageID string          `json:"messageId"`
	}

	if err := json.Unmarshal([]byte(str), &messageModel); err != nil {
		return nil, err
	}

	switch messageModel.Action {
	case PING:
		return &ServerMessagePing{newBaseServerMessage(PING)}, nil
	case CHECK:
		return &ServerMessageCheck{newBaseServerMessage(CHECK)}, nil
	case PAY:
		var payData ServerDataPay
		if err := json.Unmarshal(messageModel.Data, &payData); err != nil {
			return nil, err
		}
		return &ServerMessagePay{newBaseServerMessage(PAY), payData}, nil
	default:
		return nil, fmt.Errorf("unknown action")
	}
}
