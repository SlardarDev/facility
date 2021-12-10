package math

func MinInt(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxInt(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxInt(v ...int) int {
	max := LimitsMinInt
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinInt(v ...int) int {
	min := LimitsMaxInt
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgInt(v ...int) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumInt(v ...int) int64 {
	var sum int64
	for _, i := range v {
		sum += int64(i)
	}
	return sum
}

func MinInt8(v1, v2 int8) int8 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxInt8(v1, v2 int8) int8 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxInt8(v ...int8) int8 {
	max := LimitsMinInt8
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinInt8(v ...int8) int8 {
	min := LimitsMaxInt8
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgInt8(v ...int8) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumInt8(v ...int8) int64 {
	var sum int64
	for _, i := range v {
		sum += int64(i)
	}
	return sum
}

func MinInt16(v1, v2 int16) int16 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxInt16(v1, v2 int16) int16 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxInt16(v ...int16) int16 {
	max := LimitsMinInt16
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinInt16(v ...int16) int16 {
	min := LimitsMaxInt16
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgInt16(v ...int16) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumInt16(v ...int16) int64 {
	var sum int64
	for _, i := range v {
		sum += int64(i)
	}
	return sum
}

func MinInt32(v1, v2 int32) int32 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxInt32(v1, v2 int32) int32 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxInt32(v ...int32) int32 {
	max := LimitsMinInt32
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinInt32(v ...int32) int32 {
	min := LimitsMaxInt32
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgInt32(v ...int32) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumInt32(v ...int32) int64 {
	var sum int64
	for _, i := range v {
		sum += int64(i)
	}
	return sum
}

func MinInt64(v1, v2 int64) int64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxInt64(v1, v2 int64) int64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxInt64(v ...int64) int64 {
	max := LimitsMinInt64
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinInt64(v ...int64) int64 {
	min := LimitsMaxInt64
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgInt64(v ...int64) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumInt64(v ...int64) int64 {
	var sum int64
	for _, i := range v {
		sum += i
	}
	return sum
}

func MinUint(v1, v2 uint) uint {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxUint(v1, v2 uint) uint {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxUint(v ...uint) uint {
	max := LimitsMinUint
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinUint(v ...uint) uint {
	min := LimitsMaxUint
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgUint(v ...uint) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumUint(v ...uint) uint64 {
	var sum uint64
	for _, i := range v {
		sum += uint64(i)
	}
	return sum
}

func MinUint8(v1, v2 uint8) uint8 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxUint8(v1, v2 uint8) uint8 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxUint8(v ...uint8) uint8 {
	max := LimitsMinUint8
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinUint8(v ...uint8) uint8 {
	min := LimitsMaxUint8
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgUint8(v ...uint8) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumUint8(v ...uint8) uint64 {
	var sum uint64
	for _, i := range v {
		sum += uint64(i)
	}
	return sum
}

func MinUint16(v1, v2 uint16) uint16 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxUint16(v1, v2 uint16) uint16 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxUint16(v ...uint16) uint16 {
	max := LimitsMinUint16
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinUint16(v ...uint16) uint16 {
	min := LimitsMaxUint16
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgUint16(v ...uint16) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumUint16(v ...uint16) uint64 {
	var sum uint64
	for _, i := range v {
		sum += uint64(i)
	}
	return sum
}

func MinUint32(v1, v2 uint32) uint32 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxUint32(v1, v2 uint32) uint32 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxUint32(v ...uint32) uint32 {
	max := LimitsMinUint32
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinUint32(v ...uint32) uint32 {
	min := LimitsMaxUint32
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgUint32(v ...uint32) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumUint32(v ...uint32) uint64 {
	var sum uint64
	for _, i := range v {
		sum += uint64(i)
	}
	return sum
}

func MinUint64(v1, v2 uint64) uint64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxUint64(v1, v2 uint64) uint64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxUint64(v ...uint64) uint64 {
	max := LimitsMinUint64
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinUint64(v ...uint64) uint64 {
	min := LimitsMaxUint64
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgUint64(v ...uint64) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumUint64(v ...uint64) uint64 {
	var sum uint64
	for _, i := range v {
		sum += i
	}
	return sum
}

func MinFloat32(v1, v2 float32) float32 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxFloat32(v1, v2 float32) float32 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxFloat32(v ...float32) float32 {
	max := LimitsMinFloat32
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinFloat32(v ...float32) float32 {
	min := LimitsMaxFloat32
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgFloat32(v ...float32) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += float64(i)
	}
	return avg / float64(len(v))
}

func SliceSumFloat32(v ...float32) float32 {
	var sum float32
	for _, i := range v {
		sum += i
	}
	return sum
}

func MinFloat64(v1, v2 float64) float64 {
	if v1 < v2 {
		return v1
	}
	return v2
}

func MaxFloat64(v1, v2 float64) float64 {
	if v1 > v2 {
		return v1
	}
	return v2
}

func SliceMaxFloat64(v ...float64) float64 {
	max := LimitsMinFloat64
	for _, i := range v {
		if max < i {
			max = i
		}
	}
	return max
}

func SliceMinFloat64(v ...float64) float64 {
	min := LimitsMaxFloat64
	for _, i := range v {
		if min > i {
			min = i
		}
	}
	return min
}

func SliceAvgFloat64(v ...float64) float64 {
	var avg float64
	if len(v) == 0 {
		return avg
	}
	for _, i := range v {
		avg += i
	}
	return avg / float64(len(v))
}

func SliceSumFloat64(v ...float64) float64 {
	var sum float64
	for _, i := range v {
		sum += i
	}
	return sum
}
