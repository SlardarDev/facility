package math

import "testing"

func TestSliceMaxInt(t *testing.T) {

	t.Logf("%d", MinFloat64(-1.0, -2.0))
	t.Logf("%d", SliceMinFloat64(-1.0, -2.0))

	t.Logf("%d", SliceMinFloat64(-1.0, -2.0, -3.0))

}
