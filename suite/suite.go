package suite

import (
	"testing"

	"github.com/leiysky/go-utils/assert"
)

type Suite interface {
	SetUp()
	TearDown()
}

type TestFunc func(assert.Asserter)

func Run(t *testing.T, suite Suite, tests ...TestFunc) {
	for _, test := range tests {
		func() {
			suite.SetUp()
			defer suite.TearDown()
			a := assert.New(t)
			test(a)
		}()
	}
}
