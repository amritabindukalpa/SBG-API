package determinator_test

import (
	"testing"

	"github.com/amritabindukalpa/sbg-api/determinator"
)

type ParityTestCase struct {
	value    int
	expected bool
}

var parityTests = []ParityTestCase{
	{1, true},
	{2, false},
	{3, true},
	{4, false},
	{5, true},
	{11, true},
	{15, true},
	{20, false},
	{21, true},
}

func TestParityOutput(t *testing.T) {
	for _, tc := range parityTests {
		actual := determinator.IsOdd(tc.value)
		if actual != tc.expected {
			t.Fail()
		}
	}
}
