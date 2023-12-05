package utils

// TestCase is a helper struct to define test cases
type TestCase[InputType any, ExpectedType any] struct {
	Input    InputType
	Expected ExpectedType
}
