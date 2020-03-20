package watermillzlog

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/rs/zerolog/log"
)

// WatermillZeroLogger - implements watermill.LoggerAdapter
type WatermillZeroLogger struct {
	lf watermill.LogFields
}

func NewWatermillZeroLogger() *WatermillZeroLogger {
	return &WatermillZeroLogger{}
}

func (l *WatermillZeroLogger) Error(msg string, err error, f watermill.LogFields) {
	fields := l.lf.Add(f)
	log.Error().Err(err).Interface("fields", fields).Msg(msg)
}

func (l *WatermillZeroLogger) Info(msg string, f watermill.LogFields) {
	fields := l.lf.Add(f)
	log.Info().Interface("fields", fields).Msg(msg)
}

func (l *WatermillZeroLogger) Debug(msg string, f watermill.LogFields) {
	fields := l.lf.Add(f)
	log.Debug().Interface("fields", fields).Msg(msg)
}

func (l *WatermillZeroLogger) Trace(msg string, f watermill.LogFields) {
	fields := l.lf.Add(f)
	log.Trace().Interface("fields", fields).Msg(msg)
}

func (l *WatermillZeroLogger) With(f watermill.LogFields) watermill.LoggerAdapter {
	return &WatermillZeroLogger{
		lf: f,
	}
}
