package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/kerim-dauren/smart-terminal/api/http_api"
	"github.com/kerim-dauren/smart-terminal/config"
	"github.com/kerim-dauren/smart-terminal/pkg/loggerx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type App struct {
}

func (a *App) Start(ctx context.Context, cfg config.Config) error {
	logger := loggerx.NewLogger()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HttpPort),
		Handler: http_api.NewHttpRouter(logger, nil),
	}

	errch := make(chan error, 1)

	go func() {
		logger.Info("server started", zap.Int("port", cfg.HttpPort))
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(errch)
	}()

	select {
	case err := <-errch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
