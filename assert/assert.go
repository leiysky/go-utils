package assert

import (
	"reflect"
	"testing"
)

type Asserter interface {
	Equal(actual, expect interface{})
	NEqual(actual, expect interface{})
	True(v bool)
	False(v bool)
}

type asserter struct {
	t *testing.T
}

// New create an Asserter
func New(t *testing.T) Asserter {
	return &asserter{
		t: t,
	}
}

func (a *asserter) Equal(actual, expect interface{}) {
	a.t.Helper()
	if !reflect.DeepEqual(actual, expect) {
		a.t.Errorf("\nActual value:\n\t%T: %v\nExpect to be:\n\t%T: %v\n", actual, actual, expect, expect)
	}
}

func (a *asserter) NEqual(actual, expect interface{}) {
	a.t.Helper()
	if reflect.DeepEqual(actual, expect) {
		a.t.Errorf("\nActual value:\n\t%T: %v\nExpect not to be:\n\t%T: %v\n", actual, actual, expect, expect)
	}
}

func (a *asserter) True(v bool) {
	a.t.Helper()
	if !v {
		a.t.Errorf("\nActual value:\n\t%T: %v\nExpect to be:\n\t%T: %v\n", v, v, true, true)
	}
}

func (a *asserter) False(v bool) {
	a.t.Helper()
	if v {
		a.t.Errorf("\nActual value:\n\t%T: %v\nExpect to be:\n\t%T: %v\n", v, v, false, false)
	}
}
