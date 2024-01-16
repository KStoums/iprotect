package service

import (
	"IProtect/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ api.LoggerServiceFactory = (*ZapLoggerFactory)(nil)

type ZapLoggerFactory struct {
	api.LoggerServiceFactory
}

type ZapLoggerService struct {
	logger *zap.SugaredLogger
}

func (l *ZapLoggerFactory) NewLogger() api.LoggerService {
	config := zap.NewDevelopmentConfig()
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.StacktraceKey = ""
	encoderConfig.EncodeCaller = nil
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("02/01/2006 - 15:04:05")
	config.EncoderConfig = encoderConfig

	logger, _ := config.Build()
	defer logger.Sync() // flushes buffer, if any

	return &ZapLoggerService{logger: logger.Sugar()}
}

func (l *ZapLoggerService) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *ZapLoggerService) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args...)
}

func (l *ZapLoggerService) Debugw(msg string, keysAndValues ...interface{}) {
	l.logger.Debugw(msg, keysAndValues...)
}

func (l *ZapLoggerService) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *ZapLoggerService) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args...)
}

func (l *ZapLoggerService) Errorw(msg string, keysAndValues ...interface{}) {
	l.logger.Errorw(msg, keysAndValues...)
}

func (l *ZapLoggerService) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *ZapLoggerService) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args...)
}

func (l *ZapLoggerService) Fatalw(msg string, keysAndValues ...interface{}) {
	l.logger.Fatalw(msg, keysAndValues...)
}

func (l *ZapLoggerService) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *ZapLoggerService) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args...)
}

func (l *ZapLoggerService) Infow(msg string, keysAndValues ...interface{}) {
	l.logger.Infow(msg, keysAndValues...)
}

func (l *ZapLoggerService) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *ZapLoggerService) Panicf(template string, args ...interface{}) {
	l.logger.Panicf(template, args...)
}

func (l *ZapLoggerService) Panicw(msg string, keysAndValues ...interface{}) {
	l.logger.Panicw(msg, keysAndValues...)
}

func (l *ZapLoggerService) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *ZapLoggerService) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args...)
}

func (l *ZapLoggerService) Warnw(msg string, keysAndValues ...interface{}) {
	l.logger.Warnw(msg, keysAndValues...)
}

func (l *ZapLoggerService) With(args ...interface{}) *zap.SugaredLogger {
	return l.logger.With(args...)
}
