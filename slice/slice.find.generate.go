package slice

func IntContains(s []int, e int) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Int8Contains(s []int8, e int8) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Int16Contains(s []int16, e int16) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Int32Contains(s []int32, e int32) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Int64Contains(s []int64, e int64) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func UintContains(s []uint, e uint) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Uint8Contains(s []uint8, e uint8) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Uint16Contains(s []uint16, e uint16) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Uint32Contains(s []uint32, e uint32) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Uint64Contains(s []uint64, e uint64) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func StringContains(s []string, e string) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Float64Contains(s []float64, e float64) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func Float32Contains(s []float32, e float32) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func BoolContains(s []bool, e bool) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func InterfaceContains(s []interface{}, e interface{}) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}

func StructContains(s []struct{}, e struct{}) bool {
	for _, d := range s {
		if d == e {
			return true
		}
	}
	return false
}
