package leap

import "testing"

var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{1996, true, "leap year"},
	{1997, false, "non-leap year"},
	{1998, false, "non-leap even year"},
	{1900, false, "century"},
	{2400, true, "fourth century"},
	{2000, true, "Y2K"},
}

// Define a function IsLeapYear(int) bool.
//
// Also define a testVersion with a value that matches
// the targetTestVersion here.

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
		}
	}
}
