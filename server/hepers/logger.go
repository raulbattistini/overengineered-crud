package hepers

import (
	"errors"
	"log"
	"server/enums"

	zap "go.uber.org/zap"
)

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
