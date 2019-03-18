package assert

import (
	"fmt"
	"runtime"
)

func getFileAndLine(level int) (string, int) {
	_, file, line, _ := runtime.Caller(level)

	return file, line
}

type assertPanic struct {
	file   string
	lineNo int
	msg    string
}

func (s assertPanic) Msg() string {
	return s.msg
}

func (s assertPanic) File() string {
	return s.file
}

func (s assertPanic) Line() int {
	return s.lineNo
}

func newAssertPanic(format string, args ...interface{}) assertPanic {

	msg := fmt.Sprintf(format, args...)
	file, line := getFileAndLine(3)
	return assertPanic{
		file,
		line,
		msg,
	}

}
