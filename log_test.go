package check

import (
	"bytes"
	"testing"
)

func TestFatalIfNilSucceedsOnNonNilParameter(t *testing.T) {
	// given
	value := "test"
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	FatalIfNil(&message, value)

	// then
	if buffer.String() != "" {
		t.Error("<Fatal> should not have been called, because the given parameter is not nil.")
	}
}

func TestFatalIfNilFailsOnIfParameterIsNil(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	FatalIfNil(&message, nil)

	// then
	if buffer.String() != "my message here" {
		t.Error("Expected the invocation of <Fatal>, because the given parameter is <nil>, but was not.")
	}
}

func TestPanicIfNilSucceedsOnNonNilParameter(t *testing.T) {
	// given
	value := "test"
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	PanicIfNil(&message, value)

	// then
	if buffer.String() != "" {
		t.Error("<Panic> should not have been called, because the given parameter is not nil.")
	}
}

func TestPanicIfNilFailsOnIfParameterIsNil(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	PanicIfNil(&message, nil)

	// then
	if buffer.String() != "my message here" {
		t.Error("Expected the invocation of <Panic>, because the given parameter is <nil>, but was not.")
	}
}
