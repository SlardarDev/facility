package mapdefault

import "testing"

func TestMapDefault(t *testing.T) {

	m := map[int64]float32{
		10: 11.0,
	}
	Int64Float32(m, 11, 0) // will return 0
	Int64Float32(m, 10, 0) // will return 11.0
}
