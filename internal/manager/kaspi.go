package manager

import (
	"context"
	"errors"
	"github.com/kerim-dauren/smart-terminal/internal/domain"
	"github.com/kerim-dauren/smart-terminal/internal/service"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
	"go.uber.org/zap"
	"time"
)

type kaspiPaymentManager struct {
	logger        loggerx.Logger
	deviceService service.DeviceService
	kaspiCommands map[string]service.KaspiCommand
}

func NewKaspiPaymentManager(
	logger loggerx.Logger,
	deviceService service.DeviceService,
	kaspiCommands map[string]service.KaspiCommand,
) KaspiPaymentManager {
	return &kaspiPaymentManager{
		logger:        logger,
		deviceService: deviceService,
		kaspiCommands: kaspiCommands,
	}
}

func (m *kaspiPaymentManager) Process(ctx context.Context, request *domain.KaspiPaymentRequest, resultChan chan<- *domain.KaspiPaymentResponse) {
	device, err := m.deviceService.GetDeviceByImei(ctx, request.IMEI)
	if err != nil {
		m.logger.Error("failed to get device by imei", zap.Error(err))
		resultChan <- errorResponse(request.TransactionID, domain.InternalServerError, err)
	}

	if device == nil {
		m.logger.Error("device not found", zap.String("imei", request.IMEI))
		resultChan <- errorResponse(request.TransactionID, domain.NotFound, nil)
		return
	}

	if command, ok := m.kaspiCommands[request.Command]; !ok {
		m.logger.Error("command not found", zap.String("command", request.Command))
		resultChan <- errorResponse(request.TransactionID, domain.NotFound, nil)
	} else {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
		defer cancel()

		commandResult, err := command.Execute(ctxWithTimeout, &service.KaspiPaymentRequest{
			TransactionID:   request.TransactionID,
			IMEI:            request.IMEI,
			Command:         request.Command,
			Sum:             request.Sum,
			TransactionDate: request.TransactionDate,
		}, device)

		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				m.logger.Error("command execution timeout", zap.Error(err))
				resultChan <- errorResponse(request.TransactionID, domain.InternalServerError, err)
				return
			} else {
				m.logger.Error("failed to execute command", zap.Error(err))
				resultChan <- errorResponse(request.TransactionID, domain.InternalServerError, err)
			}
		}

		resultChan <- &domain.KaspiPaymentResponse{
			TransactionID: commandResult.TransactionID,
			Result:        commandResult.Result,
			Comment:       commandResult.Comment,
		}
	}
}

func errorResponse(transactionId int64, apiResult domain.ApiResult, err error) *domain.KaspiPaymentResponse {
	return &domain.KaspiPaymentResponse{
		TransactionID: transactionId,
		Result:        apiResult,
		Comment:       err.Error(),
	}
}
