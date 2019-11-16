package test

import (
	"testing"
)

// DefaultErrorMessage is self explanatory.
const DefaultErrorMessage = "Expected: '%v'\tActual: '%v'\tInput: '%v'"

// Case is a test case.
type Case struct {
	Input, Expected interface{}
}

// Fail throws a testing.T.Error.
func (tc Case) Fail(t *testing.T, actual interface{}) {
	t.Errorf(DefaultErrorMessage, tc.Expected, actual, tc.Input)
}
