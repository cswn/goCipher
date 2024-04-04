package cmd

import "testing"

func TestShiftText(t *testing.T) {
	msg := ShiftText("test", false, 2)
	if msg != "vguv" {
		t.Fatalf(`ShiftText("test") = %q, should have been "vguv"`, msg)
	}
}
