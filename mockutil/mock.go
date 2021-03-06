package mockutil

import "testing"

// Call is an array of type interface{} which describes a single mocked
// function call.
type Call []interface{}

// Registry collects a series of mocked calls.
type Registry struct {
	T     *testing.T
	calls []Call
}

// Register adds a MockCall to the array of registered calls.
func Register(registry *Registry, call Call) {
	registry.calls = append(registry.calls, call)
}

// Verify checks if 'expectedCall' is registered at the first position of
// the call array.
func Verify(registry *Registry, expectedCall Call) {
	registeredCall := registry.calls[0]
	registry.calls = registry.calls[1:]

	if !verifyCall(registeredCall, expectedCall) {
		registry.T.Errorf(
			"\nexpected call:\n> %s\nregistered call:\n> %s",
			expectedCall,
			registeredCall)
	}
}

func verifyCall(registeredCall Call, expectedCall Call) bool {

	if len(registeredCall) != len(expectedCall) {
		return false
	}

	for index, expected := range expectedCall {
		if expected != registeredCall[index] {
			return false
		}
	}

	return true
}
