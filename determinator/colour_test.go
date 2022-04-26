package determinator_test

import (
	"testing"

	"github.com/amritabindukalpa/sbg-api/determinator"
)

type ColourTestCase struct {
	value    int
	expected string
}

var colourTests = []ColourTestCase{
	{32, "red"},
	{19, "red"},
	{21, "red"},
	{25, "red"},
	{34, "red"},
	{27, "red"},
	{36, "red"},
	{30, "red"},
	{23, "red"},
	{5, "red"},
	{16, "red"},
	{1, "red"},
	{14, "red"},
	{9, "red"},
	{18, "red"},
	{7, "red"},
	{12, "red"},
	{3, "red"},
	{15, "black"},
	{4, "black"},
	{2, "black"},
	{17, "black"},
	{6, "black"},
	{13, "black"},
	{11, "black"},
	{8, "black"},
	{10, "black"},
	{24, "black"},
	{33, "black"},
	{20, "black"},
	{31, "black"},
	{22, "black"},
	{29, "black"},
	{28, "black"},
	{35, "black"},
	{26, "black"},
}

func TestColourOutput(t *testing.T) {
	for _, tc := range colourTests {
		actual := determinator.Colour(tc.value)
		if actual != tc.expected {
			t.Fatalf("failed for test value %d", tc.value)
		}
	}
}
