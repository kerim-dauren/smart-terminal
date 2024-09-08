package mqtt

import (
	"encoding/json"
	"fmt"
)

type DeviceMessage interface {
	GetMessageID() string
	GetAction() Action
}

type DeviceStatus string

const (
	SUCCESS DeviceStatus = "SUCCESS"
	FAILURE DeviceStatus = "FAILURE"
)

type TerminalStatus string

const (
	READY TerminalStatus = "READY"
	BUSY  TerminalStatus = "BUSY"
)

type DeviceMsgPing struct {
	MessageID string            `json:"messageId"`
	Data      DeviceMsgPingData `json:"data"`
}

type DeviceMsgPingData struct {
	Status TerminalStatus `json:"status"`
}

func (d DeviceMsgPing) GetMessageID() string {
	return d.MessageID
}

func (d DeviceMsgPing) GetAction() Action {
	return PING
}

type DeviceMsgCheck struct {
	MessageID string       `json:"messageId"`
	Status    DeviceStatus `json:"status"`
}

func (d DeviceMsgCheck) GetMessageID() string {
	return d.MessageID
}

func (d DeviceMsgCheck) GetAction() Action {
	return CHECK
}

type DeviceMsgPay struct {
	MessageID string           `json:"messageId"`
	Data      DeviceMsgPayData `json:"data"`
}

type DeviceMsgPayData struct {
	TransactionID int64        `json:"transactionId"`
	Status        DeviceStatus `json:"status"`
}

func (d DeviceMsgPay) GetMessageID() string {
	return d.MessageID
}

func (d DeviceMsgPay) GetAction() Action {
	return PAY
}

func ParseDeviceMessage(str string) (DeviceMessage, error) {
	var deviceMsg struct {
		MessageID string          `json:"messageId"`
		Action    Action          `json:"action"`
		Data      json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal([]byte(str), &deviceMsg); err != nil {
		return nil, err
	}

	switch deviceMsg.Action {
	case PING:
		var data DeviceMsgPingData
		if err := json.Unmarshal(deviceMsg.Data, &data); err != nil {
			return nil, err
		}
		return DeviceMsgPing{
			MessageID: deviceMsg.MessageID,
			Data:      data,
		}, nil
	case CHECK:
		return DeviceMsgCheck{
			MessageID: deviceMsg.MessageID,
			Status:    SUCCESS,
		}, nil
	case PAY:
		var data DeviceMsgPayData
		if err := json.Unmarshal(deviceMsg.Data, &data); err != nil {
			return nil, err
		}
		return DeviceMsgPay{
			MessageID: deviceMsg.MessageID,
			Data:      data,
		}, nil
	default:
		return nil, fmt.Errorf("unknown action")
	}
}
