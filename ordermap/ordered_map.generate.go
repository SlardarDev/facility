package ordermap

import "github.com/lixi520/facility/sorts"

func IntInt(a map[int]int) func() (int, int, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, int, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntIntRev(a map[int]int) func() (int, int, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, int, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntInt8(a map[int]int8) func() (int, int8, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, int8, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntInt8Rev(a map[int]int8) func() (int, int8, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, int8, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntInt16(a map[int]int16) func() (int, int16, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, int16, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntInt16Rev(a map[int]int16) func() (int, int16, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, int16, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntInt32(a map[int]int32) func() (int, int32, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, int32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntInt32Rev(a map[int]int32) func() (int, int32, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, int32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntInt64(a map[int]int64) func() (int, int64, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, int64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntInt64Rev(a map[int]int64) func() (int, int64, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, int64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntUint8(a map[int]uint8) func() (int, uint8, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, uint8, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntUint8Rev(a map[int]uint8) func() (int, uint8, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, uint8, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntUint16(a map[int]uint16) func() (int, uint16, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, uint16, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntUint16Rev(a map[int]uint16) func() (int, uint16, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, uint16, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntUint32(a map[int]uint32) func() (int, uint32, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, uint32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntUint32Rev(a map[int]uint32) func() (int, uint32, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, uint32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntUint64(a map[int]uint64) func() (int, uint64, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, uint64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntUint64Rev(a map[int]uint64) func() (int, uint64, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, uint64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntString(a map[int]string) func() (int, string, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, string, bool) {
		if current >= len(keys) {
			return defaultInt, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntStringRev(a map[int]string) func() (int, string, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, string, bool) {
		if current >= len(keys) {
			return defaultInt, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntFloat32(a map[int]float32) func() (int, float32, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, float32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntFloat32Rev(a map[int]float32) func() (int, float32, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, float32, bool) {
		if current >= len(keys) {
			return defaultInt, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntFloat64(a map[int]float64) func() (int, float64, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, float64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntFloat64Rev(a map[int]float64) func() (int, float64, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, float64, bool) {
		if current >= len(keys) {
			return defaultInt, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntInterface(a map[int]interface{}) func() (int, interface{}, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntInterfaceRev(a map[int]interface{}) func() (int, interface{}, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntBool(a map[int]bool) func() (int, bool, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, bool, bool) {
		if current >= len(keys) {
			return defaultInt, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntBoolRev(a map[int]bool) func() (int, bool, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, bool, bool) {
		if current >= len(keys) {
			return defaultInt, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func IntStruct(a map[int]struct{}) func() (int, struct{}, bool) {
	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Ints(keys)
	var current = 0
	return func() (int, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func IntStructRev(a map[int]struct{}) func() (int, struct{}, bool) {

	var keys []int
	for k := range a {
		keys = append(keys, k)
	}

	sorts.IntsRev(keys)
	var current = 0
	return func() (int, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Int(a map[int8]int) func() (int8, int, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, int, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8IntRev(a map[int8]int) func() (int8, int, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, int, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Int8(a map[int8]int8) func() (int8, int8, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, int8, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Int8Rev(a map[int8]int8) func() (int8, int8, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, int8, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Int16(a map[int8]int16) func() (int8, int16, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, int16, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Int16Rev(a map[int8]int16) func() (int8, int16, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, int16, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Int32(a map[int8]int32) func() (int8, int32, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, int32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Int32Rev(a map[int8]int32) func() (int8, int32, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, int32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Int64(a map[int8]int64) func() (int8, int64, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, int64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Int64Rev(a map[int8]int64) func() (int8, int64, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, int64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Uint8(a map[int8]uint8) func() (int8, uint8, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, uint8, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Uint8Rev(a map[int8]uint8) func() (int8, uint8, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, uint8, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Uint16(a map[int8]uint16) func() (int8, uint16, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, uint16, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Uint16Rev(a map[int8]uint16) func() (int8, uint16, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, uint16, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Uint32(a map[int8]uint32) func() (int8, uint32, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, uint32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Uint32Rev(a map[int8]uint32) func() (int8, uint32, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, uint32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Uint64(a map[int8]uint64) func() (int8, uint64, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, uint64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Uint64Rev(a map[int8]uint64) func() (int8, uint64, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, uint64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8String(a map[int8]string) func() (int8, string, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, string, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8StringRev(a map[int8]string) func() (int8, string, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, string, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Float32(a map[int8]float32) func() (int8, float32, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, float32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Float32Rev(a map[int8]float32) func() (int8, float32, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, float32, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Float64(a map[int8]float64) func() (int8, float64, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, float64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8Float64Rev(a map[int8]float64) func() (int8, float64, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, float64, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Interface(a map[int8]interface{}) func() (int8, interface{}, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8InterfaceRev(a map[int8]interface{}) func() (int8, interface{}, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Bool(a map[int8]bool) func() (int8, bool, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, bool, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8BoolRev(a map[int8]bool) func() (int8, bool, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, bool, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int8Struct(a map[int8]struct{}) func() (int8, struct{}, bool) {
	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8s(keys)
	var current = 0
	return func() (int8, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int8StructRev(a map[int8]struct{}) func() (int8, struct{}, bool) {

	var keys []int8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int8sRev(keys)
	var current = 0
	return func() (int8, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt8, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Int(a map[int16]int) func() (int16, int, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, int, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16IntRev(a map[int16]int) func() (int16, int, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, int, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Int8(a map[int16]int8) func() (int16, int8, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, int8, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Int8Rev(a map[int16]int8) func() (int16, int8, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, int8, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Int16(a map[int16]int16) func() (int16, int16, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, int16, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Int16Rev(a map[int16]int16) func() (int16, int16, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, int16, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Int32(a map[int16]int32) func() (int16, int32, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, int32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Int32Rev(a map[int16]int32) func() (int16, int32, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, int32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Int64(a map[int16]int64) func() (int16, int64, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, int64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Int64Rev(a map[int16]int64) func() (int16, int64, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, int64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Uint8(a map[int16]uint8) func() (int16, uint8, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, uint8, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Uint8Rev(a map[int16]uint8) func() (int16, uint8, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, uint8, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Uint16(a map[int16]uint16) func() (int16, uint16, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, uint16, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Uint16Rev(a map[int16]uint16) func() (int16, uint16, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, uint16, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Uint32(a map[int16]uint32) func() (int16, uint32, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, uint32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Uint32Rev(a map[int16]uint32) func() (int16, uint32, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, uint32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Uint64(a map[int16]uint64) func() (int16, uint64, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, uint64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Uint64Rev(a map[int16]uint64) func() (int16, uint64, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, uint64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16String(a map[int16]string) func() (int16, string, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, string, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16StringRev(a map[int16]string) func() (int16, string, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, string, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Float32(a map[int16]float32) func() (int16, float32, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, float32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Float32Rev(a map[int16]float32) func() (int16, float32, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, float32, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Float64(a map[int16]float64) func() (int16, float64, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, float64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16Float64Rev(a map[int16]float64) func() (int16, float64, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, float64, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Interface(a map[int16]interface{}) func() (int16, interface{}, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16InterfaceRev(a map[int16]interface{}) func() (int16, interface{}, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Bool(a map[int16]bool) func() (int16, bool, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, bool, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16BoolRev(a map[int16]bool) func() (int16, bool, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, bool, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int16Struct(a map[int16]struct{}) func() (int16, struct{}, bool) {
	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16s(keys)
	var current = 0
	return func() (int16, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int16StructRev(a map[int16]struct{}) func() (int16, struct{}, bool) {

	var keys []int16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int16sRev(keys)
	var current = 0
	return func() (int16, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt16, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Int(a map[int32]int) func() (int32, int, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, int, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32IntRev(a map[int32]int) func() (int32, int, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, int, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Int8(a map[int32]int8) func() (int32, int8, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, int8, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Int8Rev(a map[int32]int8) func() (int32, int8, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, int8, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Int16(a map[int32]int16) func() (int32, int16, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, int16, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Int16Rev(a map[int32]int16) func() (int32, int16, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, int16, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Int32(a map[int32]int32) func() (int32, int32, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, int32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Int32Rev(a map[int32]int32) func() (int32, int32, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, int32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Int64(a map[int32]int64) func() (int32, int64, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, int64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Int64Rev(a map[int32]int64) func() (int32, int64, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, int64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Uint8(a map[int32]uint8) func() (int32, uint8, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, uint8, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Uint8Rev(a map[int32]uint8) func() (int32, uint8, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, uint8, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Uint16(a map[int32]uint16) func() (int32, uint16, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, uint16, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Uint16Rev(a map[int32]uint16) func() (int32, uint16, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, uint16, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Uint32(a map[int32]uint32) func() (int32, uint32, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, uint32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Uint32Rev(a map[int32]uint32) func() (int32, uint32, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, uint32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Uint64(a map[int32]uint64) func() (int32, uint64, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, uint64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Uint64Rev(a map[int32]uint64) func() (int32, uint64, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, uint64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32String(a map[int32]string) func() (int32, string, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, string, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32StringRev(a map[int32]string) func() (int32, string, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, string, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Float32(a map[int32]float32) func() (int32, float32, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, float32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Float32Rev(a map[int32]float32) func() (int32, float32, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, float32, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Float64(a map[int32]float64) func() (int32, float64, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, float64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32Float64Rev(a map[int32]float64) func() (int32, float64, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, float64, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Interface(a map[int32]interface{}) func() (int32, interface{}, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32InterfaceRev(a map[int32]interface{}) func() (int32, interface{}, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Bool(a map[int32]bool) func() (int32, bool, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, bool, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32BoolRev(a map[int32]bool) func() (int32, bool, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, bool, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int32Struct(a map[int32]struct{}) func() (int32, struct{}, bool) {
	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32s(keys)
	var current = 0
	return func() (int32, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int32StructRev(a map[int32]struct{}) func() (int32, struct{}, bool) {

	var keys []int32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int32sRev(keys)
	var current = 0
	return func() (int32, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt32, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Int(a map[int64]int) func() (int64, int, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, int, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64IntRev(a map[int64]int) func() (int64, int, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, int, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Int8(a map[int64]int8) func() (int64, int8, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, int8, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Int8Rev(a map[int64]int8) func() (int64, int8, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, int8, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Int16(a map[int64]int16) func() (int64, int16, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, int16, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Int16Rev(a map[int64]int16) func() (int64, int16, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, int16, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Int32(a map[int64]int32) func() (int64, int32, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, int32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Int32Rev(a map[int64]int32) func() (int64, int32, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, int32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Int64(a map[int64]int64) func() (int64, int64, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, int64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Int64Rev(a map[int64]int64) func() (int64, int64, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, int64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Uint8(a map[int64]uint8) func() (int64, uint8, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, uint8, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Uint8Rev(a map[int64]uint8) func() (int64, uint8, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, uint8, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Uint16(a map[int64]uint16) func() (int64, uint16, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, uint16, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Uint16Rev(a map[int64]uint16) func() (int64, uint16, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, uint16, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Uint32(a map[int64]uint32) func() (int64, uint32, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, uint32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Uint32Rev(a map[int64]uint32) func() (int64, uint32, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, uint32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Uint64(a map[int64]uint64) func() (int64, uint64, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, uint64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Uint64Rev(a map[int64]uint64) func() (int64, uint64, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, uint64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64String(a map[int64]string) func() (int64, string, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, string, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64StringRev(a map[int64]string) func() (int64, string, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, string, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Float32(a map[int64]float32) func() (int64, float32, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, float32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Float32Rev(a map[int64]float32) func() (int64, float32, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, float32, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Float64(a map[int64]float64) func() (int64, float64, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, float64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64Float64Rev(a map[int64]float64) func() (int64, float64, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, float64, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Interface(a map[int64]interface{}) func() (int64, interface{}, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64InterfaceRev(a map[int64]interface{}) func() (int64, interface{}, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, interface{}, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Bool(a map[int64]bool) func() (int64, bool, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, bool, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64BoolRev(a map[int64]bool) func() (int64, bool, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, bool, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Int64Struct(a map[int64]struct{}) func() (int64, struct{}, bool) {
	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64s(keys)
	var current = 0
	return func() (int64, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Int64StructRev(a map[int64]struct{}) func() (int64, struct{}, bool) {

	var keys []int64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Int64sRev(keys)
	var current = 0
	return func() (int64, struct{}, bool) {
		if current >= len(keys) {
			return defaultInt64, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Int(a map[uint8]int) func() (uint8, int, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, int, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8IntRev(a map[uint8]int) func() (uint8, int, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, int, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Int8(a map[uint8]int8) func() (uint8, int8, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, int8, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Int8Rev(a map[uint8]int8) func() (uint8, int8, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, int8, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Int16(a map[uint8]int16) func() (uint8, int16, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, int16, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Int16Rev(a map[uint8]int16) func() (uint8, int16, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, int16, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Int32(a map[uint8]int32) func() (uint8, int32, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, int32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Int32Rev(a map[uint8]int32) func() (uint8, int32, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, int32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Int64(a map[uint8]int64) func() (uint8, int64, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, int64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Int64Rev(a map[uint8]int64) func() (uint8, int64, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, int64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Uint8(a map[uint8]uint8) func() (uint8, uint8, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, uint8, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Uint8Rev(a map[uint8]uint8) func() (uint8, uint8, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, uint8, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Uint16(a map[uint8]uint16) func() (uint8, uint16, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, uint16, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Uint16Rev(a map[uint8]uint16) func() (uint8, uint16, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, uint16, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Uint32(a map[uint8]uint32) func() (uint8, uint32, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, uint32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Uint32Rev(a map[uint8]uint32) func() (uint8, uint32, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, uint32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Uint64(a map[uint8]uint64) func() (uint8, uint64, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, uint64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Uint64Rev(a map[uint8]uint64) func() (uint8, uint64, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, uint64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8String(a map[uint8]string) func() (uint8, string, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, string, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8StringRev(a map[uint8]string) func() (uint8, string, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, string, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Float32(a map[uint8]float32) func() (uint8, float32, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, float32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Float32Rev(a map[uint8]float32) func() (uint8, float32, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, float32, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Float64(a map[uint8]float64) func() (uint8, float64, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, float64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8Float64Rev(a map[uint8]float64) func() (uint8, float64, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, float64, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Interface(a map[uint8]interface{}) func() (uint8, interface{}, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8InterfaceRev(a map[uint8]interface{}) func() (uint8, interface{}, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Bool(a map[uint8]bool) func() (uint8, bool, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, bool, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8BoolRev(a map[uint8]bool) func() (uint8, bool, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, bool, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint8Struct(a map[uint8]struct{}) func() (uint8, struct{}, bool) {
	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8s(keys)
	var current = 0
	return func() (uint8, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint8StructRev(a map[uint8]struct{}) func() (uint8, struct{}, bool) {

	var keys []uint8
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint8sRev(keys)
	var current = 0
	return func() (uint8, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint8, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Int(a map[uint16]int) func() (uint16, int, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, int, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16IntRev(a map[uint16]int) func() (uint16, int, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, int, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Int8(a map[uint16]int8) func() (uint16, int8, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, int8, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Int8Rev(a map[uint16]int8) func() (uint16, int8, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, int8, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Int16(a map[uint16]int16) func() (uint16, int16, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, int16, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Int16Rev(a map[uint16]int16) func() (uint16, int16, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, int16, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Int32(a map[uint16]int32) func() (uint16, int32, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, int32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Int32Rev(a map[uint16]int32) func() (uint16, int32, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, int32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Int64(a map[uint16]int64) func() (uint16, int64, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, int64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Int64Rev(a map[uint16]int64) func() (uint16, int64, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, int64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Uint8(a map[uint16]uint8) func() (uint16, uint8, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, uint8, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Uint8Rev(a map[uint16]uint8) func() (uint16, uint8, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, uint8, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Uint16(a map[uint16]uint16) func() (uint16, uint16, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, uint16, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Uint16Rev(a map[uint16]uint16) func() (uint16, uint16, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, uint16, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Uint32(a map[uint16]uint32) func() (uint16, uint32, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, uint32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Uint32Rev(a map[uint16]uint32) func() (uint16, uint32, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, uint32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Uint64(a map[uint16]uint64) func() (uint16, uint64, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, uint64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Uint64Rev(a map[uint16]uint64) func() (uint16, uint64, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, uint64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16String(a map[uint16]string) func() (uint16, string, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, string, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16StringRev(a map[uint16]string) func() (uint16, string, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, string, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Float32(a map[uint16]float32) func() (uint16, float32, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, float32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Float32Rev(a map[uint16]float32) func() (uint16, float32, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, float32, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Float64(a map[uint16]float64) func() (uint16, float64, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, float64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16Float64Rev(a map[uint16]float64) func() (uint16, float64, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, float64, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Interface(a map[uint16]interface{}) func() (uint16, interface{}, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16InterfaceRev(a map[uint16]interface{}) func() (uint16, interface{}, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Bool(a map[uint16]bool) func() (uint16, bool, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, bool, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16BoolRev(a map[uint16]bool) func() (uint16, bool, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, bool, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint16Struct(a map[uint16]struct{}) func() (uint16, struct{}, bool) {
	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16s(keys)
	var current = 0
	return func() (uint16, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint16StructRev(a map[uint16]struct{}) func() (uint16, struct{}, bool) {

	var keys []uint16
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint16sRev(keys)
	var current = 0
	return func() (uint16, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint16, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Int(a map[uint32]int) func() (uint32, int, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, int, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32IntRev(a map[uint32]int) func() (uint32, int, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, int, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Int8(a map[uint32]int8) func() (uint32, int8, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, int8, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Int8Rev(a map[uint32]int8) func() (uint32, int8, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, int8, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Int16(a map[uint32]int16) func() (uint32, int16, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, int16, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Int16Rev(a map[uint32]int16) func() (uint32, int16, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, int16, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Int32(a map[uint32]int32) func() (uint32, int32, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, int32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Int32Rev(a map[uint32]int32) func() (uint32, int32, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, int32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Int64(a map[uint32]int64) func() (uint32, int64, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, int64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Int64Rev(a map[uint32]int64) func() (uint32, int64, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, int64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Uint8(a map[uint32]uint8) func() (uint32, uint8, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, uint8, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Uint8Rev(a map[uint32]uint8) func() (uint32, uint8, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, uint8, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Uint16(a map[uint32]uint16) func() (uint32, uint16, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, uint16, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Uint16Rev(a map[uint32]uint16) func() (uint32, uint16, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, uint16, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Uint32(a map[uint32]uint32) func() (uint32, uint32, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, uint32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Uint32Rev(a map[uint32]uint32) func() (uint32, uint32, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, uint32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Uint64(a map[uint32]uint64) func() (uint32, uint64, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, uint64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Uint64Rev(a map[uint32]uint64) func() (uint32, uint64, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, uint64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32String(a map[uint32]string) func() (uint32, string, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, string, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32StringRev(a map[uint32]string) func() (uint32, string, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, string, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Float32(a map[uint32]float32) func() (uint32, float32, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, float32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Float32Rev(a map[uint32]float32) func() (uint32, float32, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, float32, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Float64(a map[uint32]float64) func() (uint32, float64, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, float64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32Float64Rev(a map[uint32]float64) func() (uint32, float64, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, float64, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Interface(a map[uint32]interface{}) func() (uint32, interface{}, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32InterfaceRev(a map[uint32]interface{}) func() (uint32, interface{}, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Bool(a map[uint32]bool) func() (uint32, bool, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, bool, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32BoolRev(a map[uint32]bool) func() (uint32, bool, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, bool, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint32Struct(a map[uint32]struct{}) func() (uint32, struct{}, bool) {
	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32s(keys)
	var current = 0
	return func() (uint32, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint32StructRev(a map[uint32]struct{}) func() (uint32, struct{}, bool) {

	var keys []uint32
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint32sRev(keys)
	var current = 0
	return func() (uint32, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint32, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Int(a map[uint64]int) func() (uint64, int, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, int, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64IntRev(a map[uint64]int) func() (uint64, int, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, int, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Int8(a map[uint64]int8) func() (uint64, int8, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, int8, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Int8Rev(a map[uint64]int8) func() (uint64, int8, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, int8, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Int16(a map[uint64]int16) func() (uint64, int16, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, int16, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Int16Rev(a map[uint64]int16) func() (uint64, int16, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, int16, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Int32(a map[uint64]int32) func() (uint64, int32, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, int32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Int32Rev(a map[uint64]int32) func() (uint64, int32, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, int32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Int64(a map[uint64]int64) func() (uint64, int64, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, int64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Int64Rev(a map[uint64]int64) func() (uint64, int64, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, int64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Uint8(a map[uint64]uint8) func() (uint64, uint8, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, uint8, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Uint8Rev(a map[uint64]uint8) func() (uint64, uint8, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, uint8, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Uint16(a map[uint64]uint16) func() (uint64, uint16, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, uint16, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Uint16Rev(a map[uint64]uint16) func() (uint64, uint16, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, uint16, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Uint32(a map[uint64]uint32) func() (uint64, uint32, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, uint32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Uint32Rev(a map[uint64]uint32) func() (uint64, uint32, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, uint32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Uint64(a map[uint64]uint64) func() (uint64, uint64, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, uint64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Uint64Rev(a map[uint64]uint64) func() (uint64, uint64, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, uint64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64String(a map[uint64]string) func() (uint64, string, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, string, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64StringRev(a map[uint64]string) func() (uint64, string, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, string, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Float32(a map[uint64]float32) func() (uint64, float32, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, float32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Float32Rev(a map[uint64]float32) func() (uint64, float32, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, float32, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Float64(a map[uint64]float64) func() (uint64, float64, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, float64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64Float64Rev(a map[uint64]float64) func() (uint64, float64, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, float64, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Interface(a map[uint64]interface{}) func() (uint64, interface{}, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64InterfaceRev(a map[uint64]interface{}) func() (uint64, interface{}, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, interface{}, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Bool(a map[uint64]bool) func() (uint64, bool, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, bool, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64BoolRev(a map[uint64]bool) func() (uint64, bool, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, bool, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func Uint64Struct(a map[uint64]struct{}) func() (uint64, struct{}, bool) {
	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64s(keys)
	var current = 0
	return func() (uint64, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func Uint64StructRev(a map[uint64]struct{}) func() (uint64, struct{}, bool) {

	var keys []uint64
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Uint64sRev(keys)
	var current = 0
	return func() (uint64, struct{}, bool) {
		if current >= len(keys) {
			return defaultUint64, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInt(a map[string]int) func() (string, int, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, int, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringIntRev(a map[string]int) func() (string, int, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, int, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInt8(a map[string]int8) func() (string, int8, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, int8, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringInt8Rev(a map[string]int8) func() (string, int8, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, int8, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInt16(a map[string]int16) func() (string, int16, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, int16, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringInt16Rev(a map[string]int16) func() (string, int16, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, int16, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInt32(a map[string]int32) func() (string, int32, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, int32, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringInt32Rev(a map[string]int32) func() (string, int32, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, int32, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInt64(a map[string]int64) func() (string, int64, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, int64, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringInt64Rev(a map[string]int64) func() (string, int64, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, int64, bool) {
		if current >= len(keys) {
			return defaultString, defaultInt64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringUint8(a map[string]uint8) func() (string, uint8, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, uint8, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringUint8Rev(a map[string]uint8) func() (string, uint8, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, uint8, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint8, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringUint16(a map[string]uint16) func() (string, uint16, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, uint16, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringUint16Rev(a map[string]uint16) func() (string, uint16, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, uint16, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint16, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringUint32(a map[string]uint32) func() (string, uint32, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, uint32, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringUint32Rev(a map[string]uint32) func() (string, uint32, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, uint32, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringUint64(a map[string]uint64) func() (string, uint64, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, uint64, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringUint64Rev(a map[string]uint64) func() (string, uint64, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, uint64, bool) {
		if current >= len(keys) {
			return defaultString, defaultUint64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringString(a map[string]string) func() (string, string, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, string, bool) {
		if current >= len(keys) {
			return defaultString, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringStringRev(a map[string]string) func() (string, string, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, string, bool) {
		if current >= len(keys) {
			return defaultString, defaultString, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringFloat32(a map[string]float32) func() (string, float32, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, float32, bool) {
		if current >= len(keys) {
			return defaultString, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringFloat32Rev(a map[string]float32) func() (string, float32, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, float32, bool) {
		if current >= len(keys) {
			return defaultString, defaultFloat32, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringFloat64(a map[string]float64) func() (string, float64, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, float64, bool) {
		if current >= len(keys) {
			return defaultString, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringFloat64Rev(a map[string]float64) func() (string, float64, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, float64, bool) {
		if current >= len(keys) {
			return defaultString, defaultFloat64, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringInterface(a map[string]interface{}) func() (string, interface{}, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, interface{}, bool) {
		if current >= len(keys) {
			return defaultString, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringInterfaceRev(a map[string]interface{}) func() (string, interface{}, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, interface{}, bool) {
		if current >= len(keys) {
			return defaultString, defaultInterface, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringBool(a map[string]bool) func() (string, bool, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, bool, bool) {
		if current >= len(keys) {
			return defaultString, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringBoolRev(a map[string]bool) func() (string, bool, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, bool, bool) {
		if current >= len(keys) {
			return defaultString, defaultBool, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}

func StringStruct(a map[string]struct{}) func() (string, struct{}, bool) {
	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.Strings(keys)
	var current = 0
	return func() (string, struct{}, bool) {
		if current >= len(keys) {
			return defaultString, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}
}

func StringStructRev(a map[string]struct{}) func() (string, struct{}, bool) {

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sorts.StringsRev(keys)
	var current = 0
	return func() (string, struct{}, bool) {
		if current >= len(keys) {
			return defaultString, defaultStruct, false
		}
		k, v := keys[current], a[keys[current]]
		current = current + 1
		return k, v, true
	}

}
