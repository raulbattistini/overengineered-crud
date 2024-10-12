package hepers

import (
	"errors"
	"log"
	"server/enums"

	zap "go.uber.org/zap"
)

type LoggerInterface interface {
	LogError(message *string, err error)
	LogDebug(message string)
	LogInfo(message string)
	LogWarn(message string)
}

type Logger struct {
	Loggr *zap.Logger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	return &Logger{
		Loggr: logger,
	}
}

func (l *Logger) LogError(message *string, err error) {
	l.Loggr.Error(*message, zap.Error(err))
}

func (l *Logger) LogDebug(message string) {
	l.Loggr.Debug(message)
}

func (l *Logger) LogInfo(message string) {
	l.Loggr.Info(message)
}

func (l *Logger) LogWarn(message string) {
	l.Loggr.Warn(message)
}

// ou msg seria map[string]interface{}
func Log(msg interface{}, err *error, level enums.LogLevels) {
	/* contorno mal feito */
	var e error
	if err != nil {
		e = *err
	}
	/* contorno mal feito */

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	// outro contorno
	finalMsg, ok := msg.(string)
	if !ok {
		log.Printf("nvalid message to print: %v, not logged structuredly, error: %w", msg, e)
	}
	// finalMsg = strings.ReplaceAll(finalMsg, "\n", "")

	switch level {
	case enums.Error:
		sugar.Error(finalMsg, e)
	case enums.Warn:
		sugar.Warn(finalMsg)
	case enums.Info:
		sugar.Infow(finalMsg)
	case enums.Debug:
		sugar.Infof(finalMsg)
	}
}

func NewErrorFromMessage(message string) error {
	message = CleanAllInput(message)
	return errors.New(message)
}
