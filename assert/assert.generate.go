package assert

import (
	"github.com/lixi520/facility/set"
	"github.com/lixi520/facility/slice"
)

func EqualIntSliceAsSet(a, b []int) {
	if !slice.EqualIntSliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualIntSlice(a, b []int) {
	if !slice.EqualIntSlice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualIntSliceAsSet(a, b []int) {
	if slice.EqualIntSliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualIntSlice(a, b []int) {
	if slice.EqualIntSlice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualIntSet(s1, s2 set.IntSet) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualIntSet(s1, s2 set.IntSet) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualIntSliceAsSetWithMsg(a, b []int, msg string) {
	if !slice.EqualIntSliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualIntSliceWithMsg(a, b []int, msg string) {
	if !slice.EqualIntSlice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualIntSliceAsSetWithMsg(a, b []int, msg string) {
	if slice.EqualIntSliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualIntSliceWithMsg(a, b []int, msg string) {
	if slice.EqualIntSlice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualIntSetWithMsg(s1, s2 set.IntSet, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualIntSetWithMsg(s1, s2 set.IntSet, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt(a, b int) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt(a, b int) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualIntWithMsg(a, b int, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualIntWithMsg(a, b int, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualInt8SliceAsSet(a, b []int8) {
	if !slice.EqualInt8SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualInt8Slice(a, b []int8) {
	if !slice.EqualInt8Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt8SliceAsSet(a, b []int8) {
	if slice.EqualInt8SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualInt8Slice(a, b []int8) {
	if slice.EqualInt8Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt8Set(s1, s2 set.Int8Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualInt8Set(s1, s2 set.Int8Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualInt8SliceAsSetWithMsg(a, b []int8, msg string) {
	if !slice.EqualInt8SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt8SliceWithMsg(a, b []int8, msg string) {
	if !slice.EqualInt8Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt8SliceAsSetWithMsg(a, b []int8, msg string) {
	if slice.EqualInt8SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt8SliceWithMsg(a, b []int8, msg string) {
	if slice.EqualInt8Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt8SetWithMsg(s1, s2 set.Int8Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt8SetWithMsg(s1, s2 set.Int8Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt8(a, b int8) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt8(a, b int8) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt8WithMsg(a, b int8, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt8WithMsg(a, b int8, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualInt16SliceAsSet(a, b []int16) {
	if !slice.EqualInt16SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualInt16Slice(a, b []int16) {
	if !slice.EqualInt16Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt16SliceAsSet(a, b []int16) {
	if slice.EqualInt16SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualInt16Slice(a, b []int16) {
	if slice.EqualInt16Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt16Set(s1, s2 set.Int16Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualInt16Set(s1, s2 set.Int16Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualInt16SliceAsSetWithMsg(a, b []int16, msg string) {
	if !slice.EqualInt16SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt16SliceWithMsg(a, b []int16, msg string) {
	if !slice.EqualInt16Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt16SliceAsSetWithMsg(a, b []int16, msg string) {
	if slice.EqualInt16SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt16SliceWithMsg(a, b []int16, msg string) {
	if slice.EqualInt16Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt16SetWithMsg(s1, s2 set.Int16Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt16SetWithMsg(s1, s2 set.Int16Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt16(a, b int16) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt16(a, b int16) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt16WithMsg(a, b int16, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt16WithMsg(a, b int16, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualInt32SliceAsSet(a, b []int32) {
	if !slice.EqualInt32SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualInt32Slice(a, b []int32) {
	if !slice.EqualInt32Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt32SliceAsSet(a, b []int32) {
	if slice.EqualInt32SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualInt32Slice(a, b []int32) {
	if slice.EqualInt32Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt32Set(s1, s2 set.Int32Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualInt32Set(s1, s2 set.Int32Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualInt32SliceAsSetWithMsg(a, b []int32, msg string) {
	if !slice.EqualInt32SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt32SliceWithMsg(a, b []int32, msg string) {
	if !slice.EqualInt32Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt32SliceAsSetWithMsg(a, b []int32, msg string) {
	if slice.EqualInt32SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt32SliceWithMsg(a, b []int32, msg string) {
	if slice.EqualInt32Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt32SetWithMsg(s1, s2 set.Int32Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt32SetWithMsg(s1, s2 set.Int32Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt32(a, b int32) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt32(a, b int32) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt32WithMsg(a, b int32, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt32WithMsg(a, b int32, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualInt64SliceAsSet(a, b []int64) {
	if !slice.EqualInt64SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualInt64Slice(a, b []int64) {
	if !slice.EqualInt64Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt64SliceAsSet(a, b []int64) {
	if slice.EqualInt64SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualInt64Slice(a, b []int64) {
	if slice.EqualInt64Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt64Set(s1, s2 set.Int64Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualInt64Set(s1, s2 set.Int64Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualInt64SliceAsSetWithMsg(a, b []int64, msg string) {
	if !slice.EqualInt64SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt64SliceWithMsg(a, b []int64, msg string) {
	if !slice.EqualInt64Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt64SliceAsSetWithMsg(a, b []int64, msg string) {
	if slice.EqualInt64SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt64SliceWithMsg(a, b []int64, msg string) {
	if slice.EqualInt64Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt64SetWithMsg(s1, s2 set.Int64Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt64SetWithMsg(s1, s2 set.Int64Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualInt64(a, b int64) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualInt64(a, b int64) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualInt64WithMsg(a, b int64, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualInt64WithMsg(a, b int64, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualUint8SliceAsSet(a, b []uint8) {
	if !slice.EqualUint8SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualUint8Slice(a, b []uint8) {
	if !slice.EqualUint8Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint8SliceAsSet(a, b []uint8) {
	if slice.EqualUint8SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualUint8Slice(a, b []uint8) {
	if slice.EqualUint8Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint8Set(s1, s2 set.Uint8Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualUint8Set(s1, s2 set.Uint8Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualUint8SliceAsSetWithMsg(a, b []uint8, msg string) {
	if !slice.EqualUint8SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint8SliceWithMsg(a, b []uint8, msg string) {
	if !slice.EqualUint8Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint8SliceAsSetWithMsg(a, b []uint8, msg string) {
	if slice.EqualUint8SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint8SliceWithMsg(a, b []uint8, msg string) {
	if slice.EqualUint8Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint8SetWithMsg(s1, s2 set.Uint8Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint8SetWithMsg(s1, s2 set.Uint8Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint8(a, b uint8) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint8(a, b uint8) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint8WithMsg(a, b uint8, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint8WithMsg(a, b uint8, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualUint16SliceAsSet(a, b []uint16) {
	if !slice.EqualUint16SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualUint16Slice(a, b []uint16) {
	if !slice.EqualUint16Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint16SliceAsSet(a, b []uint16) {
	if slice.EqualUint16SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualUint16Slice(a, b []uint16) {
	if slice.EqualUint16Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint16Set(s1, s2 set.Uint16Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualUint16Set(s1, s2 set.Uint16Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualUint16SliceAsSetWithMsg(a, b []uint16, msg string) {
	if !slice.EqualUint16SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint16SliceWithMsg(a, b []uint16, msg string) {
	if !slice.EqualUint16Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint16SliceAsSetWithMsg(a, b []uint16, msg string) {
	if slice.EqualUint16SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint16SliceWithMsg(a, b []uint16, msg string) {
	if slice.EqualUint16Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint16SetWithMsg(s1, s2 set.Uint16Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint16SetWithMsg(s1, s2 set.Uint16Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint16(a, b uint16) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint16(a, b uint16) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint16WithMsg(a, b uint16, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint16WithMsg(a, b uint16, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualUint32SliceAsSet(a, b []uint32) {
	if !slice.EqualUint32SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualUint32Slice(a, b []uint32) {
	if !slice.EqualUint32Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint32SliceAsSet(a, b []uint32) {
	if slice.EqualUint32SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualUint32Slice(a, b []uint32) {
	if slice.EqualUint32Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint32Set(s1, s2 set.Uint32Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualUint32Set(s1, s2 set.Uint32Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualUint32SliceAsSetWithMsg(a, b []uint32, msg string) {
	if !slice.EqualUint32SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint32SliceWithMsg(a, b []uint32, msg string) {
	if !slice.EqualUint32Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint32SliceAsSetWithMsg(a, b []uint32, msg string) {
	if slice.EqualUint32SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint32SliceWithMsg(a, b []uint32, msg string) {
	if slice.EqualUint32Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint32SetWithMsg(s1, s2 set.Uint32Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint32SetWithMsg(s1, s2 set.Uint32Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint32(a, b uint32) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint32(a, b uint32) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint32WithMsg(a, b uint32, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint32WithMsg(a, b uint32, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualUint64SliceAsSet(a, b []uint64) {
	if !slice.EqualUint64SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualUint64Slice(a, b []uint64) {
	if !slice.EqualUint64Slice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint64SliceAsSet(a, b []uint64) {
	if slice.EqualUint64SliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualUint64Slice(a, b []uint64) {
	if slice.EqualUint64Slice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint64Set(s1, s2 set.Uint64Set) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualUint64Set(s1, s2 set.Uint64Set) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualUint64SliceAsSetWithMsg(a, b []uint64, msg string) {
	if !slice.EqualUint64SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint64SliceWithMsg(a, b []uint64, msg string) {
	if !slice.EqualUint64Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint64SliceAsSetWithMsg(a, b []uint64, msg string) {
	if slice.EqualUint64SliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint64SliceWithMsg(a, b []uint64, msg string) {
	if slice.EqualUint64Slice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint64SetWithMsg(s1, s2 set.Uint64Set, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint64SetWithMsg(s1, s2 set.Uint64Set, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualUint64(a, b uint64) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualUint64(a, b uint64) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualUint64WithMsg(a, b uint64, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualUint64WithMsg(a, b uint64, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualFloat32(a, b float32) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualFloat32(a, b float32) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualFloat32WithMsg(a, b float32, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualFloat32WithMsg(a, b float32, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualFloat64(a, b float64) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualFloat64(a, b float64) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualFloat64WithMsg(a, b float64, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualFloat64WithMsg(a, b float64, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func EqualStringSliceAsSet(a, b []string) {
	if !slice.EqualStringSliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func EqualStringSlice(a, b []string) {
	if !slice.EqualStringSlice(a, b) {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualStringSliceAsSet(a, b []string) {
	if slice.EqualStringSliceAsSet(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func NotEqualStringSlice(a, b []string) {
	if slice.EqualStringSlice(a, b) {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualStringSet(s1, s2 set.StringSet) {
	if !s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is equal", s1, s2))
	}
}

func NotEqualStringSet(s1, s2 set.StringSet) {
	if s1.Equal(s2) {
		panic(newAssertPanic("%v and %v is not equal", s1, s2))
	}
}

func EqualStringSliceAsSetWithMsg(a, b []string, msg string) {
	if !slice.EqualStringSliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualStringSliceWithMsg(a, b []string, msg string) {
	if !slice.EqualStringSlice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualStringSliceAsSetWithMsg(a, b []string, msg string) {
	if slice.EqualStringSliceAsSet(a, b) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualStringSliceWithMsg(a, b []string, msg string) {
	if slice.EqualStringSlice(a, b) {
		panic(newAssertPanic(msg))
	}
}

func EqualStringSetWithMsg(s1, s2 set.StringSet, msg string) {
	if !s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func NotEqualStringSetWithMsg(s1, s2 set.StringSet, msg string) {
	if s1.Equal(s2) {
		panic(newAssertPanic(msg))
	}
}

func EqualString(a, b string) {
	if a != b {
		panic(newAssertPanic("%v and %v is not equal", a, b))
	}
}

func NotEqualString(a, b string) {
	if a == b {
		panic(newAssertPanic("%v and %v is equal", a, b))
	}
}

func EqualStringWithMsg(a, b string, msg string) {
	if a != b {
		panic(newAssertPanic(msg))
	}
}

func NotEqualStringWithMsg(a, b string, msg string) {
	if a == b {
		panic(newAssertPanic(msg))
	}
}

func GreatInt(a, b int) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleInt(a, b int) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatIntWithMsg(a, b int, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleIntWithMsg(a, b int, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatInt8(a, b int8) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleInt8(a, b int8) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatInt8WithMsg(a, b int8, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleInt8WithMsg(a, b int8, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatInt16(a, b int16) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleInt16(a, b int16) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatInt16WithMsg(a, b int16, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleInt16WithMsg(a, b int16, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatInt32(a, b int32) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleInt32(a, b int32) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatInt32WithMsg(a, b int32, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleInt32WithMsg(a, b int32, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatInt64(a, b int64) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleInt64(a, b int64) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatInt64WithMsg(a, b int64, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleInt64WithMsg(a, b int64, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatUint8(a, b uint8) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleUint8(a, b uint8) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatUint8WithMsg(a, b uint8, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleUint8WithMsg(a, b uint8, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatUint16(a, b uint16) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleUint16(a, b uint16) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatUint16WithMsg(a, b uint16, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleUint16WithMsg(a, b uint16, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatUint32(a, b uint32) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleUint32(a, b uint32) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatUint32WithMsg(a, b uint32, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleUint32WithMsg(a, b uint32, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatUint64(a, b uint64) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleUint64(a, b uint64) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatUint64WithMsg(a, b uint64, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleUint64WithMsg(a, b uint64, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatFloat32(a, b float32) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleFloat32(a, b float32) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatFloat32WithMsg(a, b float32, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleFloat32WithMsg(a, b float32, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}

func GreatFloat64(a, b float64) {
	if a <= b {
		panic(newAssertPanic("%v is less than %v", a, b))
	}
}

func LittleFloat64(a, b float64) {
	if a >= b {
		panic(newAssertPanic("%v is great than %v", a, b))
	}
}

func GreatFloat64WithMsg(a, b float64, msg string) {
	if a <= b {
		panic(newAssertPanic(msg))
	}
}

func LittleFloat64WithMsg(a, b float64, msg string) {
	if a >= b {
		panic(newAssertPanic(msg))
	}
}
