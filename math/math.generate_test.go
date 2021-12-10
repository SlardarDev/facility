package math

import "testing"

func TestSliceMaxInt(t *testing.T) {

	t.Logf("%f", MinFloat64(-1.0, -2.0))
	t.Logf("%f", SliceMinFloat64(-1.0, -2.0))

	t.Logf("%f", SliceMinFloat64(-1.0, -2.0, -3.0))
	t.Logf("%f", SliceAvgFloat64(-1.0, 2.0, 3.0))
	t.Logf("%f", SliceSumFloat64(-1.0, 2.0, 3.0))
	t.Logf("%f", SliceAvgInt8(-1.0, 2.0, 3.0))
	t.Logf("%d", SliceSumInt(2, 3, 4))
}
