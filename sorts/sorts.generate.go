package sorts

import "sort"

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type intSliceRev []int

func (p intSliceRev) Len() int           { return len(p) }
func (p intSliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p intSliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Ints(s []int) {
	sort.Sort(intSlice(s))
}

func IntsRev(s []int) {
	sort.Sort(intSliceRev(s))
}

func SearchInts(s1 []int, s2 int) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchIntsRev(s1 []int, s2 int) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type int8Slice []int8

func (p int8Slice) Len() int           { return len(p) }
func (p int8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type int8SliceRev []int8

func (p int8SliceRev) Len() int           { return len(p) }
func (p int8SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p int8SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Int8s(s []int8) {
	sort.Sort(int8Slice(s))
}

func Int8sRev(s []int8) {
	sort.Sort(int8SliceRev(s))
}

func SearchInt8s(s1 []int8, s2 int8) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchInt8sRev(s1 []int8, s2 int8) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type int16Slice []int16

func (p int16Slice) Len() int           { return len(p) }
func (p int16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type int16SliceRev []int16

func (p int16SliceRev) Len() int           { return len(p) }
func (p int16SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p int16SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Int16s(s []int16) {
	sort.Sort(int16Slice(s))
}

func Int16sRev(s []int16) {
	sort.Sort(int16SliceRev(s))
}

func SearchInt16s(s1 []int16, s2 int16) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchInt16sRev(s1 []int16, s2 int16) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type int32Slice []int32

func (p int32Slice) Len() int           { return len(p) }
func (p int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type int32SliceRev []int32

func (p int32SliceRev) Len() int           { return len(p) }
func (p int32SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p int32SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Int32s(s []int32) {
	sort.Sort(int32Slice(s))
}

func Int32sRev(s []int32) {
	sort.Sort(int32SliceRev(s))
}

func SearchInt32s(s1 []int32, s2 int32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchInt32sRev(s1 []int32, s2 int32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type int64Slice []int64

func (p int64Slice) Len() int           { return len(p) }
func (p int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type int64SliceRev []int64

func (p int64SliceRev) Len() int           { return len(p) }
func (p int64SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p int64SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Int64s(s []int64) {
	sort.Sort(int64Slice(s))
}

func Int64sRev(s []int64) {
	sort.Sort(int64SliceRev(s))
}

func SearchInt64s(s1 []int64, s2 int64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchInt64sRev(s1 []int64, s2 int64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type uintSlice []uint

func (p uintSlice) Len() int           { return len(p) }
func (p uintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p uintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type uintSliceRev []uint

func (p uintSliceRev) Len() int           { return len(p) }
func (p uintSliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p uintSliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uints(s []uint) {
	sort.Sort(uintSlice(s))
}

func UintsRev(s []uint) {
	sort.Sort(uintSliceRev(s))
}

func SearchUints(s1 []uint, s2 uint) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchUintsRev(s1 []uint, s2 uint) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type uint8Slice []uint8

func (p uint8Slice) Len() int           { return len(p) }
func (p uint8Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p uint8Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type uint8SliceRev []uint8

func (p uint8SliceRev) Len() int           { return len(p) }
func (p uint8SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p uint8SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uint8s(s []uint8) {
	sort.Sort(uint8Slice(s))
}

func Uint8sRev(s []uint8) {
	sort.Sort(uint8SliceRev(s))
}

func SearchUint8s(s1 []uint8, s2 uint8) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchUint8sRev(s1 []uint8, s2 uint8) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type uint16Slice []uint16

func (p uint16Slice) Len() int           { return len(p) }
func (p uint16Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p uint16Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type uint16SliceRev []uint16

func (p uint16SliceRev) Len() int           { return len(p) }
func (p uint16SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p uint16SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uint16s(s []uint16) {
	sort.Sort(uint16Slice(s))
}

func Uint16sRev(s []uint16) {
	sort.Sort(uint16SliceRev(s))
}

func SearchUint16s(s1 []uint16, s2 uint16) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchUint16sRev(s1 []uint16, s2 uint16) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type uint32Slice []uint32

func (p uint32Slice) Len() int           { return len(p) }
func (p uint32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p uint32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type uint32SliceRev []uint32

func (p uint32SliceRev) Len() int           { return len(p) }
func (p uint32SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p uint32SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uint32s(s []uint32) {
	sort.Sort(uint32Slice(s))
}

func Uint32sRev(s []uint32) {
	sort.Sort(uint32SliceRev(s))
}

func SearchUint32s(s1 []uint32, s2 uint32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchUint32sRev(s1 []uint32, s2 uint32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type uint64Slice []uint64

func (p uint64Slice) Len() int           { return len(p) }
func (p uint64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p uint64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type uint64SliceRev []uint64

func (p uint64SliceRev) Len() int           { return len(p) }
func (p uint64SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p uint64SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Uint64s(s []uint64) {
	sort.Sort(uint64Slice(s))
}

func Uint64sRev(s []uint64) {
	sort.Sort(uint64SliceRev(s))
}

func SearchUint64s(s1 []uint64, s2 uint64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchUint64sRev(s1 []uint64, s2 uint64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type float32Slice []float32

func (p float32Slice) Len() int           { return len(p) }
func (p float32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p float32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type float32SliceRev []float32

func (p float32SliceRev) Len() int           { return len(p) }
func (p float32SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p float32SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Float32s(s []float32) {
	sort.Sort(float32Slice(s))
}

func Float32sRev(s []float32) {
	sort.Sort(float32SliceRev(s))
}

func SearchFloat32s(s1 []float32, s2 float32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchFloat32sRev(s1 []float32, s2 float32) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type float64Slice []float64

func (p float64Slice) Len() int           { return len(p) }
func (p float64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type float64SliceRev []float64

func (p float64SliceRev) Len() int           { return len(p) }
func (p float64SliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p float64SliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Float64s(s []float64) {
	sort.Sort(float64Slice(s))
}

func Float64sRev(s []float64) {
	sort.Sort(float64SliceRev(s))
}

func SearchFloat64s(s1 []float64, s2 float64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchFloat64sRev(s1 []float64, s2 float64) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type stringSlice []string

func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type stringSliceRev []string

func (p stringSliceRev) Len() int           { return len(p) }
func (p stringSliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p stringSliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Strings(s []string) {
	sort.Sort(stringSlice(s))
}

func StringsRev(s []string) {
	sort.Sort(stringSliceRev(s))
}

func SearchStrings(s1 []string, s2 string) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchStringsRev(s1 []string, s2 string) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}

type runeSlice []rune

func (p runeSlice) Len() int           { return len(p) }
func (p runeSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p runeSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type runeSliceRev []rune

func (p runeSliceRev) Len() int           { return len(p) }
func (p runeSliceRev) Less(i, j int) bool { return p[i] > p[j] }
func (p runeSliceRev) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Runes(s []rune) {
	sort.Sort(runeSlice(s))
}

func RunesRev(s []rune) {
	sort.Sort(runeSliceRev(s))
}

func SearchRunes(s1 []rune, s2 rune) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] >= s2
	})
}

func SearchRunesRev(s1 []rune, s2 rune) int {
	return sort.Search(len(s1), func(i int) bool {
		return s1[i] <= s2
	})
}
