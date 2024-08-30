package loggerx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

type DefaultLogger struct {
	zapLogger   *zap.Logger
	serviceName string
	env         Env
	level       Level
	handler     Handler
	fileConfig  *FileConfig
}

type FileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

// NewLogger
// Логгер для всех сервисов.
// Принимает параметры serviceName - название сервиса и devStage - флаг дев и прод среды.
// Возвращает структуру Logger
func NewLogger(opts ...Option) *DefaultLogger {
	l := &DefaultLogger{
		serviceName: "",
		env:         Dev,
		level:       Debug,
		handler:     JSON,
		fileConfig:  nil,
	}
	for _, opt := range opts {
		opt(l)
	}

	return &DefaultLogger{
		zapLogger: l.initZapLogger(),
	}
}

func (l *DefaultLogger) initZapLogger() *zap.Logger {
	var zapConfig zapcore.EncoderConfig
	// log level enabler
	logLevel, err := zapcore.ParseLevel(string(l.level))
	if err != nil {
		logLevel = zapcore.DebugLevel
	}

	// error and fatal level enabler
	errorFatalLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.ErrorLevel || level == zapcore.FatalLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level == zapcore.WarnLevel
	})
	stdoutSyncer := zapcore.Lock(os.Stdout)
	stderrSyncer := zapcore.Lock(os.Stderr)

	if l.env == Prod {
		zapConfig = zap.NewProductionEncoderConfig()
		zapConfig.CallerKey = zapcore.OmitKey
		zapConfig.EncodeCaller = zapcore.ShortCallerEncoder
	} else {
		zapConfig = zap.NewDevelopmentEncoderConfig()
		zapConfig.CallerKey = "caller_key"
	}
	zapConfig.TimeKey = "timestamp"
	//zapConfig.FunctionKey = "method"
	zapConfig.LevelKey = "level"
	zapConfig.MessageKey = "message"
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	zapConfig.EncodeCaller = zapcore.ShortCallerEncoder

	var fileWriter zapcore.WriteSyncer
	if l.fileConfig != nil {
		fileWriter = zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.fileConfig.Filename,
			MaxSize:    l.fileConfig.MaxSize, // megabytes
			MaxBackups: l.fileConfig.MaxBackups,
			MaxAge:     l.fileConfig.MaxAge, // days
		})
	}

	var encoder zapcore.Encoder
	if l.handler == JSON {
		encoder = zapcore.NewJSONEncoder(zapConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(zapConfig)
	}

	consoleCore := []zapcore.Core{
		zapcore.NewCore(
			encoder,
			stdoutSyncer,
			logLevel,
		),
		zapcore.NewCore(
			encoder,
			stderrSyncer,
			errorFatalLevel,
		),
		zapcore.NewCore(
			encoder,
			stdoutSyncer,
			warnLevel,
		),
	}

	var core zapcore.Core
	if fileWriter != nil {
		fileCore := []zapcore.Core{
			zapcore.NewCore(
				encoder,
				fileWriter,
				logLevel,
			),
			zapcore.NewCore(
				encoder,
				fileWriter,
				warnLevel,
			),
			zapcore.NewCore(
				encoder,
				fileWriter,
				errorFatalLevel,
			),
		}
		core = zapcore.NewTee(append(consoleCore, fileCore...)...)
	} else {
		core = zapcore.NewTee(consoleCore...)
	}

	envFields := zap.Fields(zap.String("env", string(l.env)))
	serviceFields := zap.Fields()
	if l.serviceName != "" {
		serviceFields = zap.Fields(zap.String("service", l.serviceName))
	}

	return zap.New(
		core,
		envFields,
		serviceFields,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
}

func (l *DefaultLogger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

// Errorf использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Errorf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Errorf(msg, args...)
}

func (l *DefaultLogger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Warnf использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Warnf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Warnf(msg, args...)
}

func (l *DefaultLogger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Fatalf использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Fatalf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Fatalf(msg, args...)
}

func (l *DefaultLogger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Infof использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Infof(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Infof(msg, args...)
}

func (l *DefaultLogger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Debugf использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Debugf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Debugf(msg, args...)
}

func (l *DefaultLogger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

// Panicf использует fmt.Sprintf для логирование шаблонного сообщения.
func (l *DefaultLogger) Panicf(msg string, args ...interface{}) {
	l.zapLogger.Sugar().Panicf(msg, args...)
}

// Zap возвращает zap.Logger.
func (l *DefaultLogger) Zap() *zap.Logger {
	return l.zapLogger
}
