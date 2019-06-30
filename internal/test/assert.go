package test

import (
	"testing"
)

// Assert is a statement that a predicate (Boolean-valued function, i.e. a true–false expression) is always true at that point in code execution.
type Assert struct {
	t *testing.T
}

// NewAssert returns a new Assert.
func NewAssert(t *testing.T) Assert {
	return Assert{t}
}

// Equals throws an error if the expected value doesn't equal the actual value.
func (a Assert) Equals(tc Case, actual interface{}) {
	if tc.Expected != actual {
		tc.Fail(a.t, actual)
	}
}
