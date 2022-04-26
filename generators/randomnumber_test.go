package generators_test

import (
	"testing"

	"github.com/amritabindukalpa/sbg-api/generators"
)

func TestColourOutput(t *testing.T) {
	var numLimit = []int{3, 4, 5, 6, 7, 12, 34, 51, 78, 89, 100}
	for _, tc := range numLimit {
		actual := generators.RandomNumber(tc)
		if actual > tc {
			t.Fatalf("Failed to generate random number within the range for %d. Number generated %d", tc, actual)
		}
	}
}
