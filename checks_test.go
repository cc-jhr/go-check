package check

import (
	"bytes"
	"testing"
)

func TestCheckFails(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	Check(&message, func() bool { return false })

	// then
	if buffer.String() != message {
		t.Error("Expected <Check> to fail, because condition returns <false>, but it didn't.")
	}
}

func TestCheckSucceeds(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	Check(&message, func() bool { return true })

	// then
	if buffer.String() != "" {
		t.Error("Expected <Check> to succeed, because condition returns <true>, but it didn't.")
	}
}

func TestCheckNotNilFails(t *testing.T) {
	// given
	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	CheckNotNil(nil)

	// then
	if buffer.String() != "Reference must not be nil" {
		t.Error("Expected <CheckNotNil> to fail, because given parameter is <nil>, but it didn't.")
	}
}

func TestCheckNotNilSucceeds(t *testing.T) {
	// given
	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	CheckNotNil("test")

	// then
	if buffer.String() != "" {
		t.Error("Expected <CheckNotNil> to succeed, because given parameter is not <nil>, but it didn't.")
	}
}

func TestRequireFails(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	Require(&message, func() bool { return false })

	// then
	if buffer.String() != message {
		t.Error("Expected <Require> to fail, because condition returns <false>, but it didn't.")
	}
}

func TestRequireSucceeds(t *testing.T) {
	// given
	message := "my message here"

	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	Require(&message, func() bool { return true })

	// then
	if buffer.String() != "" {
		t.Error("Expected <Require> to succeed, because condition returns <true>, but it didn't.")
	}
}

func TestRequireNotNilFails(t *testing.T) {
	// given
	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	RequireNotNil(nil)

	// then
	if buffer.String() != "Reference must not be nil" {
		t.Error("Expected <RequireNotNil> to fail, because given parameter is <nil>, but it didn't.")
	}
}

func TestRequireNotNilSucceeds(t *testing.T) {
	// given
	buffer := bytes.NewBufferString("")
	SetBuffer(buffer)

	// when
	RequireNotNil("test")

	// then
	if buffer.String() != "" {
		t.Error("Expected <RequireNotNil> to succeed, because given parameter is not <nil>, but it didn't.")
	}
}
