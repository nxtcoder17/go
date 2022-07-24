package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type KV struct {
	Key   string
	Value any
}

type Logger interface {
	Debugf(msg string, args ...any)
	Infof(msg string, args ...any)
	Errorf(err error, msg string, args ...any)
	Warnf(msg string, args ...any)
	WithKV(kv KV) Logger
	WithName(name string) Logger
}

type customLogger struct {
	zapLogger *zap.SugaredLogger
}

func (c customLogger) WithKV(kv KV) Logger {
	return &customLogger{zapLogger: c.zapLogger.With(kv.Key, kv.Value)}
}

func (c customLogger) Debugf(msg string, args ...any) {
	c.zapLogger.Debugf(msg, args...)
}

func (c customLogger) Infof(msg string, args ...any) {
	c.zapLogger.Infof(msg, args...)
}

func (c customLogger) Errorf(err error, msg string, args ...any) {
	c.zapLogger.Errorf("%s AS %+v happened", fmt.Sprintf(msg, args...), err)
}

func (c customLogger) Warnf(msg string, args ...any) {
	c.zapLogger.Warnf(msg, args...)
}

func (c customLogger) WithName(name string) Logger {
	return &customLogger{zapLogger: c.zapLogger.Named(name)}
}

type Options struct {
	Name      string
	Dev       bool
	ZapOption zap.Option
}

func New(options *Options) (Logger, error) {
	opts := Options{}
	if options != nil {
		opts = *options
	}
	if opts.Dev {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		cfg.EncoderConfig.LineEnding = "\n\n"
		cfg.EncoderConfig.TimeKey = ""
		logger, err := func () (*zap.Logger, error){
			if opts.ZapOption != nil {
				return cfg.Build(zap.AddCallerSkip(1), opts.ZapOption)
			}
			return cfg.Build(zap.AddCallerSkip(1))
		}()
		if err != nil {
			return nil, err
		}
		if opts.Name != "" {
			return &customLogger{zapLogger: logger.Sugar().Named(opts.Name)}, nil
		}
		return &customLogger{zapLogger: logger.Sugar()}, nil
	}
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	if opts.Name != "" {
		return &customLogger{zapLogger: logger.Sugar().Named(opts.Name)}, nil
	}
	return &customLogger{zapLogger: logger.Sugar()}, nil
}

func NewOrDie(options *Options) Logger {
	logger, err := New(options)
	if err != nil {
		panic(err)
	}
	return logger
}
