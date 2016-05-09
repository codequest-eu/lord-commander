package lordcommander

import (
	"bytes"

	"github.com/stretchr/testify/mock"
)

// MockCommander is a mock implementation of lordcommander.Commander.
type MockCommander struct {
	mock.Mock
}

// Run is a mock implementation of Commander's Run method.
func (mc *MockCommander) Run(command string, args ...string) error {
	arguments := []interface{}{command}
	for _, arg := range args {
		arguments = append(arguments, arg)
	}
	return mc.Called(arguments...).Error(0)
}

// Stdout is a mock implementation of Commander's Stdout method.
func (mc *MockCommander) Stdout() *bytes.Buffer {
	return mc.Called().Get(0).(*bytes.Buffer)
}

// Stderr is a mock implementation of Commander's Stderr method.
func (mc *MockCommander) Stderr() *bytes.Buffer {
	return mc.Called().Get(0).(*bytes.Buffer)
}
