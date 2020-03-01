package order

import (
	"testing"
)

func TestAddress_Merge(t *testing.T) {
	tests := []struct {
		arg      Address
		expected Address
	}{
		{
			Address{
				ZipCode: "300-0000",
			},
			Address{
				ZipCode: "300-0000",
			},
		},
	}

	for _, test := range tests {
		a := Address{}
		a.Merge(&test.arg)
		testAddress(t, a, test.expected)
	}
}

func testAddress(t *testing.T, subject Address, expected Address) bool {
	if subject.ZipCode != expected.ZipCode {
		t.Errorf("Address not equal: got=%v want=%v", subject, expected)
		return false
	}

	return true
}
