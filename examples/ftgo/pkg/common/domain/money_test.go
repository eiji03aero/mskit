package domain

import (
	"testing"
)

func TestMoney_Equals(t *testing.T) {
	t.Run("with proper args", func(t *testing.T) {
		tests := []struct {
			arg      int
			target   interface{}
			expected bool
		}{
			{10, 10, true},
			{10, 20, false},
		}

		for _, test := range tests {
			m := Money(test.arg)
			if m.Equals(test.target) != test.expected {
				t.Fatal(test)
			}
		}
	})
}
