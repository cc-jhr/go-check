package check

var errorMessage = "Reference must not be nil"

func Check(message *string, condition condition) {
	PanicIf(message, func() bool { return !condition() })
}

func CheckNotNil(ref interface{}) {
	Check(&errorMessage, func() bool { return ref != nil })
}

func Require(message *string, condition condition) {
	FatalIf(message, func() bool { return !condition() })
}

func RequireNotNil(ref interface{}) {
	Require(&errorMessage, func() bool { return ref != nil })
}
