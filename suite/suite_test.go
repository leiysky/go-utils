package suite

import (
	"fmt"
	"testing"

	"github.com/leiysky/go-utils/assert"
)

type suite struct {
}

func (s *suite) SetUp() {
	fmt.Println("set up")
}

func (s *suite) TearDown() {
	fmt.Println("tear down")
}

func TestSuite(t *testing.T) {
	tests := []TestFunc{
		func(a assert.Asserter) {
			fmt.Println("test 1")
		},
		func(a assert.Asserter) {
			fmt.Println("test 2")
		},
		func(a assert.Asserter) {
			fmt.Println("test 3")
		},
	}
	Run(t, &suite{}, tests...)
}
