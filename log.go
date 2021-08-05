package check

import (
	"bytes"
	"log"
)

type condition func() bool

var delegate *log.Logger = log.Default()
var buffer *bytes.Buffer = nil

func SetBuffer(_buffer *bytes.Buffer) {
	buffer = _buffer
}

func FatalIf(message *string, condition condition) {
	if condition() {
		fatal(message)
	}
}

func FatalIfNil(message *string, obj interface{}) {
	FatalIf(message, func() bool { return obj == nil })
}

func fatal(message *string) {
	if buffer == nil {
		delegate.Fatal(message)
	} else {
		buffer.WriteString(*message)
	}
}

func PanicIf(message *string, condition condition) {
	if condition() {
		panic(message)
	}
}

func PanicIfNil(message *string, obj interface{}) {
	PanicIf(message, func() bool { return obj == nil })
}

func panic(message *string) {
	if buffer == nil {
		delegate.Panic(message)
	} else {
		buffer.WriteString(*message)
	}
}
