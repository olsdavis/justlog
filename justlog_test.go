package justlog

import (
	"testing"
)

var (
	logger = NewWithHandlers(NewConsoleHandler(DebugLevel), NewFileHandler("test/test.txt", ErrorLevel, DebugLevel)).SetFormatters(NewFormatter("%{LEVEL}:%{TIME} -> %{MESSAGE}"))
	lineLogger = NewWithHandlers(NewConsoleHandler()).SetFormatters(NewFormatter("%{SHORT_CALLER} %{LONG_CALLER} %{PID} EH OUI %{MESSAGE}. Swag "))
	// Known bug: add a trailing space at the end of the format, the last char is removed if you have some text after an argument
)

// It is quite annoying to test a logging library with unit tests.
// So I just tried out a lot of different possible things with the library.
// Do not worry, the library will work: I am going to use it, so if I notice a bug,
// I will fix it.
func TestLogger(t *testing.T) {
	logger.Debug("yeah", "wow")
	logger.Error("Output to the file.")
}

func TestCallerLogger(t *testing.T) {
	lineLogger.Info("youpidiyeah")
}

func BenchmarkFileWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger.Debug("test message quite short")
	}
}
