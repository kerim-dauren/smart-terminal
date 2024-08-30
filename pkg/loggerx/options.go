package loggerx

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"strings"
)

type Option func(logger *DefaultLogger)

func WithServiceName(serviceName string) Option {
	return func(c *DefaultLogger) {
		c.serviceName = serviceName
	}
}

func WithEnv(env Env) Option {
	return func(c *DefaultLogger) {
		c.env = env
	}
}

func WithLevel(level Level) Option {
	return func(c *DefaultLogger) {
		c.level = level
	}
}

func WithHandler(handler Handler) Option {
	return func(c *DefaultLogger) {
		c.handler = handler
	}
}

func WithFileConfig(cfg *FileConfig) Option {
	return func(c *DefaultLogger) {
		c.fileConfig = cfg
	}
}

// WithFields adds the zap.Field's to an existing logger.
func WithFields(logger Logger, fields ...zap.Field) *DefaultLogger {
	return &DefaultLogger{
		zapLogger: logger.(*DefaultLogger).zapLogger.With(fields...),
	}
}

// WithContextStandardFields adds the zap.Field's with httpx.StandardHeaderKeys stored in the context to an existing logger.
func WithContextStandardFields(ctx context.Context, logger Logger) *DefaultLogger {
	return &DefaultLogger{
		zapLogger: logger.(*DefaultLogger).zapLogger.With(GetStandardFields(ctx)...),
	}
}

// WithContext gets the zap.Field's from a context and adds them to an existing logger.
func WithContext(ctx context.Context, logger Logger) *DefaultLogger {
	return &DefaultLogger{
		zapLogger: logger.(*DefaultLogger).zapLogger.With(GetFields(ctx)...),
	}
}

// GetFields returns the zap.Field's stored in the context or nil if none are found.
func GetFields(ctx context.Context, keys ...string) []zap.Field {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	if len(keys) > 0 {
		var fields []zap.Field
		for _, key := range keys {
			if val, ok := md[strings.ToLower(key)]; ok {
				fields = append(fields, zap.String(key, val[0]))
			}
		}

		return fields
	}

	var fields []zap.Field
	for name, value := range md {
		fields = append(fields, zap.String(name, value[0]))
	}

	return fields
}

// GetStandardFields returns the zap.Field's with httpx.StandardHeaderKeys stored in the context or nil if none are found.
func GetStandardFields(ctx context.Context) []zap.Field {
	var fields []zap.Field

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	if len(md) != 0 {
		for _, key := range standardHeaderKeys {
			keyLower := strings.ToLower(key)
			if val, ok := md[keyLower]; ok {
				fields = append(fields, zap.String(keyLower, val[0]))
			}
		}
	}

	return fields
}

// GetField returns the zap.Field stored in the context or nil if none are found.
func GetField(ctx context.Context, key string) *zap.Field {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	keyLower := strings.ToLower(key)

	if val, ok := md[keyLower]; ok {
		field := zap.String(keyLower, val[0])
		return &field
	}

	return nil
}
