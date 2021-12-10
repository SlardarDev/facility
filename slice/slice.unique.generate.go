package slice

import "github.com/SlardarDev/facility/set"

func UniqueIntSlice(eles []int) []int {
	s := set.IntSet{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualIntSliceAsSet(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.IntSet{}
	sb := set.IntSet{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueInt8Slice(eles []int8) []int8 {
	s := set.Int8Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualInt8Slice(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualInt8SliceAsSet(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Int8Set{}
	sb := set.Int8Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueInt16Slice(eles []int16) []int16 {
	s := set.Int16Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualInt16Slice(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualInt16SliceAsSet(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Int16Set{}
	sb := set.Int16Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueInt32Slice(eles []int32) []int32 {
	s := set.Int32Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualInt32Slice(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualInt32SliceAsSet(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Int32Set{}
	sb := set.Int32Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueInt64Slice(eles []int64) []int64 {
	s := set.Int64Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualInt64Slice(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualInt64SliceAsSet(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Int64Set{}
	sb := set.Int64Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueUintSlice(eles []uint) []uint {
	s := set.UintSet{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualUintSlice(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualUintSliceAsSet(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.UintSet{}
	sb := set.UintSet{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueUint8Slice(eles []uint8) []uint8 {
	s := set.Uint8Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualUint8Slice(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualUint8SliceAsSet(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Uint8Set{}
	sb := set.Uint8Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueUint16Slice(eles []uint16) []uint16 {
	s := set.Uint16Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualUint16Slice(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualUint16SliceAsSet(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Uint16Set{}
	sb := set.Uint16Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueUint32Slice(eles []uint32) []uint32 {
	s := set.Uint32Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualUint32Slice(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualUint32SliceAsSet(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Uint32Set{}
	sb := set.Uint32Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueUint64Slice(eles []uint64) []uint64 {
	s := set.Uint64Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualUint64Slice(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualUint64SliceAsSet(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.Uint64Set{}
	sb := set.Uint64Set{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueStringSlice(eles []string) []string {
	s := set.StringSet{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualStringSliceAsSet(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.StringSet{}
	sb := set.StringSet{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}

func UniqueFloat64Slice(eles []float64) []float64 {
	s := set.Float64Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func UniqueFloat32Slice(eles []float32) []float32 {
	s := set.Float32Set{}
	s.AddAll(eles...)
	return s.ToList()
}

func UniqueBoolSlice(eles []bool) []bool {
	s := set.BoolSet{}
	s.AddAll(eles...)
	return s.ToList()
}

func EqualBoolSlice(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, v := range a {
		if b[idx] != v {
			return false
		}
	}
	return true
}

func EqualBoolSliceAsSet(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 {
		return true
	}

	sa := set.BoolSet{}
	sb := set.BoolSet{}

	sa.AddAll(a...)
	sb.AddAll(b...)
	return sa.Equal(sb)

}
