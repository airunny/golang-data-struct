package stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMultiBracket(t *testing.T) {
	tests := []struct {
		In       string
		Expected bool
	}{
		{
			In:       "",
			Expected: false,
		},
		{
			In:       "(",
			Expected: false,
		},
		{
			In:       "()",
			Expected: true,
		},
		{
			In:       "(())",
			Expected: true,
		},
		{
			In:       "(([[]]))",
			Expected: true,
		},
		{
			In:       "(([[))]]",
			Expected: false,
		},
		{
			In:       "(([[{{}}]]))",
			Expected: true,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, CheckMultiBracket(test.In))
	}
}

func Test001(t *testing.T) {
	fmt.Println(3 + 8 - 5*10/2)
}

func TestNewCalculator(t *testing.T) {
	tests := []struct {
		Express  string
		Err      bool
		Expected float64
	}{
		{
			Express:  "3+8*5+10",
			Err:      false,
			Expected: 53,
		},
		{
			Express:  "3 + 8 - 5*10/2",
			Err:      false,
			Expected: -14,
		},
		{
			Express:  "2+2",
			Err:      false,
			Expected: 4,
		},
		{
			Express:  "2-2",
			Err:      false,
			Expected: 0,
		},

		{
			Express:  "2*2",
			Err:      false,
			Expected: 4,
		},
		{
			Express:  "2/2",
			Err:      false,
			Expected: 1,
		},
	}

	cal := NewCalculator()
	for _, test := range tests {
		res, err := cal.Calculate(test.Express)
		assert.Equal(t, test.Err, err != nil)
		assert.Equal(t, test.Expected, res)
	}
}
