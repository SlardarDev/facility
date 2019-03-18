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
