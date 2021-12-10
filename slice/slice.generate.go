package slice

import "github.com/SlardarDev/facility/conv"

func IntToInt8(s []int) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToInt8Must(s []int) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToInt16(s []int) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToInt16Must(s []int) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToInt32(s []int) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToInt32Must(s []int) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToInt64(s []int) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToInt64Must(s []int) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToUint(s []int) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToUintMust(s []int) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToUint8(s []int) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToUint8Must(s []int) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToUint16(s []int) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToUint16Must(s []int) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToUint32(s []int) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToUint32Must(s []int) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToUint64(s []int) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToUint64Must(s []int) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToString(s []int) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToStringMust(s []int) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToFloat32(s []int) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToFloat32Must(s []int) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToFloat64(s []int) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToFloat64Must(s []int) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToBool(s []int) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func IntToBoolMust(s []int) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func IntToInterface(s []int) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Int8ToInt(s []int8) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToIntMust(s []int8) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToInt16(s []int8) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToInt16Must(s []int8) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToInt32(s []int8) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToInt32Must(s []int8) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToInt64(s []int8) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToInt64Must(s []int8) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToUint(s []int8) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToUintMust(s []int8) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToUint8(s []int8) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToUint8Must(s []int8) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToUint16(s []int8) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToUint16Must(s []int8) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToUint32(s []int8) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToUint32Must(s []int8) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToUint64(s []int8) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToUint64Must(s []int8) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToString(s []int8) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToStringMust(s []int8) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToFloat32(s []int8) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToFloat32Must(s []int8) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToFloat64(s []int8) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToFloat64Must(s []int8) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToBool(s []int8) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int8ToBoolMust(s []int8) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int8ToInterface(s []int8) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Int16ToInt(s []int16) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToIntMust(s []int16) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToInt8(s []int16) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToInt8Must(s []int16) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToInt32(s []int16) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToInt32Must(s []int16) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToInt64(s []int16) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToInt64Must(s []int16) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToUint(s []int16) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToUintMust(s []int16) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToUint8(s []int16) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToUint8Must(s []int16) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToUint16(s []int16) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToUint16Must(s []int16) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToUint32(s []int16) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToUint32Must(s []int16) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToUint64(s []int16) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToUint64Must(s []int16) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToString(s []int16) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToStringMust(s []int16) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToFloat32(s []int16) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToFloat32Must(s []int16) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToFloat64(s []int16) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToFloat64Must(s []int16) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToBool(s []int16) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int16ToBoolMust(s []int16) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int16ToInterface(s []int16) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Int32ToInt(s []int32) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToIntMust(s []int32) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToInt8(s []int32) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToInt8Must(s []int32) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToInt16(s []int32) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToInt16Must(s []int32) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToInt64(s []int32) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToInt64Must(s []int32) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToUint(s []int32) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToUintMust(s []int32) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToUint8(s []int32) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToUint8Must(s []int32) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToUint16(s []int32) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToUint16Must(s []int32) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToUint32(s []int32) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToUint32Must(s []int32) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToUint64(s []int32) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToUint64Must(s []int32) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToString(s []int32) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToStringMust(s []int32) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToFloat32(s []int32) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToFloat32Must(s []int32) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToFloat64(s []int32) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToFloat64Must(s []int32) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToBool(s []int32) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int32ToBoolMust(s []int32) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int32ToInterface(s []int32) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Int64ToInt(s []int64) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToIntMust(s []int64) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToInt8(s []int64) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToInt8Must(s []int64) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToInt16(s []int64) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToInt16Must(s []int64) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToInt32(s []int64) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToInt32Must(s []int64) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToUint(s []int64) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToUintMust(s []int64) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToUint8(s []int64) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToUint8Must(s []int64) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToUint16(s []int64) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToUint16Must(s []int64) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToUint32(s []int64) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToUint32Must(s []int64) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToUint64(s []int64) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToUint64Must(s []int64) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToString(s []int64) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToStringMust(s []int64) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToFloat32(s []int64) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToFloat32Must(s []int64) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToFloat64(s []int64) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToFloat64Must(s []int64) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToBool(s []int64) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Int64ToBoolMust(s []int64) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Int64ToInterface(s []int64) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func UintToInt(s []uint) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToIntMust(s []uint) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToInt8(s []uint) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToInt8Must(s []uint) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToInt16(s []uint) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToInt16Must(s []uint) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToInt32(s []uint) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToInt32Must(s []uint) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToInt64(s []uint) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToInt64Must(s []uint) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToUint8(s []uint) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToUint8Must(s []uint) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToUint16(s []uint) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToUint16Must(s []uint) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToUint32(s []uint) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToUint32Must(s []uint) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToUint64(s []uint) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToUint64Must(s []uint) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToString(s []uint) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToStringMust(s []uint) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToFloat32(s []uint) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToFloat32Must(s []uint) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToFloat64(s []uint) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToFloat64Must(s []uint) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToBool(s []uint) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func UintToBoolMust(s []uint) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func UintToInterface(s []uint) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Uint8ToInt(s []uint8) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToIntMust(s []uint8) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToInt8(s []uint8) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToInt8Must(s []uint8) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToInt16(s []uint8) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToInt16Must(s []uint8) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToInt32(s []uint8) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToInt32Must(s []uint8) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToInt64(s []uint8) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToInt64Must(s []uint8) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToUint(s []uint8) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToUintMust(s []uint8) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToUint16(s []uint8) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToUint16Must(s []uint8) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToUint32(s []uint8) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToUint32Must(s []uint8) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToUint64(s []uint8) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToUint64Must(s []uint8) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToString(s []uint8) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToStringMust(s []uint8) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToFloat32(s []uint8) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToFloat32Must(s []uint8) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToFloat64(s []uint8) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToFloat64Must(s []uint8) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToBool(s []uint8) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint8ToBoolMust(s []uint8) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint8ToInterface(s []uint8) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Uint16ToInt(s []uint16) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToIntMust(s []uint16) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToInt8(s []uint16) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToInt8Must(s []uint16) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToInt16(s []uint16) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToInt16Must(s []uint16) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToInt32(s []uint16) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToInt32Must(s []uint16) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToInt64(s []uint16) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToInt64Must(s []uint16) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToUint(s []uint16) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToUintMust(s []uint16) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToUint8(s []uint16) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToUint8Must(s []uint16) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToUint32(s []uint16) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToUint32Must(s []uint16) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToUint64(s []uint16) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToUint64Must(s []uint16) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToString(s []uint16) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToStringMust(s []uint16) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToFloat32(s []uint16) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToFloat32Must(s []uint16) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToFloat64(s []uint16) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToFloat64Must(s []uint16) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToBool(s []uint16) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint16ToBoolMust(s []uint16) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint16ToInterface(s []uint16) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Uint32ToInt(s []uint32) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToIntMust(s []uint32) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToInt8(s []uint32) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToInt8Must(s []uint32) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToInt16(s []uint32) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToInt16Must(s []uint32) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToInt32(s []uint32) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToInt32Must(s []uint32) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToInt64(s []uint32) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToInt64Must(s []uint32) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToUint(s []uint32) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToUintMust(s []uint32) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToUint8(s []uint32) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToUint8Must(s []uint32) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToUint16(s []uint32) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToUint16Must(s []uint32) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToUint64(s []uint32) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToUint64Must(s []uint32) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToString(s []uint32) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToStringMust(s []uint32) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToFloat32(s []uint32) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToFloat32Must(s []uint32) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToFloat64(s []uint32) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToFloat64Must(s []uint32) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToBool(s []uint32) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint32ToBoolMust(s []uint32) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint32ToInterface(s []uint32) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Uint64ToInt(s []uint64) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToIntMust(s []uint64) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToInt8(s []uint64) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToInt8Must(s []uint64) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToInt16(s []uint64) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToInt16Must(s []uint64) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToInt32(s []uint64) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToInt32Must(s []uint64) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToInt64(s []uint64) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToInt64Must(s []uint64) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToUint(s []uint64) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToUintMust(s []uint64) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToUint8(s []uint64) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToUint8Must(s []uint64) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToUint16(s []uint64) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToUint16Must(s []uint64) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToUint32(s []uint64) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToUint32Must(s []uint64) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToString(s []uint64) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToStringMust(s []uint64) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToFloat32(s []uint64) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToFloat32Must(s []uint64) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToFloat64(s []uint64) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToFloat64Must(s []uint64) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToBool(s []uint64) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Uint64ToBoolMust(s []uint64) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Uint64ToInterface(s []uint64) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func StringToInt(s []string) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToIntMust(s []string) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToInt8(s []string) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToInt8Must(s []string) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToInt16(s []string) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToInt16Must(s []string) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToInt32(s []string) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToInt32Must(s []string) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToInt64(s []string) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToInt64Must(s []string) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToUint(s []string) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToUintMust(s []string) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToUint8(s []string) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToUint8Must(s []string) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToUint16(s []string) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToUint16Must(s []string) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToUint32(s []string) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToUint32Must(s []string) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToUint64(s []string) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToUint64Must(s []string) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToFloat32(s []string) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToFloat32Must(s []string) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToFloat64(s []string) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToFloat64Must(s []string) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToBool(s []string) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func StringToBoolMust(s []string) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func StringToInterface(s []string) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Float32ToInt(s []float32) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToIntMust(s []float32) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToInt8(s []float32) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToInt8Must(s []float32) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToInt16(s []float32) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToInt16Must(s []float32) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToInt32(s []float32) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToInt32Must(s []float32) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToInt64(s []float32) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToInt64Must(s []float32) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToUint(s []float32) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToUintMust(s []float32) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToUint8(s []float32) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToUint8Must(s []float32) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToUint16(s []float32) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToUint16Must(s []float32) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToUint32(s []float32) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToUint32Must(s []float32) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToUint64(s []float32) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToUint64Must(s []float32) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToString(s []float32) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToStringMust(s []float32) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToFloat64(s []float32) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToFloat64Must(s []float32) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToBool(s []float32) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float32ToBoolMust(s []float32) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float32ToInterface(s []float32) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func Float64ToInt(s []float64) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToIntMust(s []float64) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToInt8(s []float64) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToInt8Must(s []float64) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToInt16(s []float64) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToInt16Must(s []float64) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToInt32(s []float64) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToInt32Must(s []float64) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToInt64(s []float64) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToInt64Must(s []float64) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToUint(s []float64) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToUintMust(s []float64) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToUint8(s []float64) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToUint8Must(s []float64) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToUint16(s []float64) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToUint16Must(s []float64) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToUint32(s []float64) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToUint32Must(s []float64) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToUint64(s []float64) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToUint64Must(s []float64) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToString(s []float64) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToStringMust(s []float64) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToFloat32(s []float64) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToFloat32Must(s []float64) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToBool(s []float64) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func Float64ToBoolMust(s []float64) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func Float64ToInterface(s []float64) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func BoolToInt(s []bool) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToIntMust(s []bool) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToInt8(s []bool) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToInt8Must(s []bool) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToInt16(s []bool) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToInt16Must(s []bool) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToInt32(s []bool) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToInt32Must(s []bool) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToInt64(s []bool) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToInt64Must(s []bool) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToUint(s []bool) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToUintMust(s []bool) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToUint8(s []bool) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToUint8Must(s []bool) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToUint16(s []bool) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToUint16Must(s []bool) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToUint32(s []bool) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToUint32Must(s []bool) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToUint64(s []bool) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToUint64Must(s []bool) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToString(s []bool) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToStringMust(s []bool) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToFloat32(s []bool) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToFloat32Must(s []bool) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToFloat64(s []bool) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func BoolToFloat64Must(s []bool) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func BoolToInterface(s []bool) []interface{} {
	ret := make([]interface{}, 0, len(s))
	for _, ele := range s {
		ret = append(ret, ele)
	}
	return ret
}

func InterfaceToInt(s []interface{}) ([]int, error) {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToIntMust(s []interface{}) []int {
	ret := make([]int, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToInt8(s []interface{}) ([]int8, error) {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToInt8Must(s []interface{}) []int8 {
	ret := make([]int8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToInt16(s []interface{}) ([]int16, error) {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToInt16Must(s []interface{}) []int16 {
	ret := make([]int16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToInt32(s []interface{}) ([]int32, error) {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToInt32Must(s []interface{}) []int32 {
	ret := make([]int32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToInt64(s []interface{}) ([]int64, error) {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToInt64Must(s []interface{}) []int64 {
	ret := make([]int64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Int64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToUint(s []interface{}) ([]uint, error) {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToUintMust(s []interface{}) []uint {
	ret := make([]uint, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToUint8(s []interface{}) ([]uint8, error) {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToUint8Must(s []interface{}) []uint8 {
	ret := make([]uint8, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint8(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToUint16(s []interface{}) ([]uint16, error) {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToUint16Must(s []interface{}) []uint16 {
	ret := make([]uint16, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint16(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToUint32(s []interface{}) ([]uint32, error) {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToUint32Must(s []interface{}) []uint32 {
	ret := make([]uint32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToUint64(s []interface{}) ([]uint64, error) {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToUint64Must(s []interface{}) []uint64 {
	ret := make([]uint64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Uint64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToString(s []interface{}) ([]string, error) {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToStringMust(s []interface{}) []string {
	ret := make([]string, 0, len(s))
	for _, ele := range s {
		o, e := conv.String(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToFloat32(s []interface{}) ([]float32, error) {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToFloat32Must(s []interface{}) []float32 {
	ret := make([]float32, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float32(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToFloat64(s []interface{}) ([]float64, error) {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToFloat64Must(s []interface{}) []float64 {
	ret := make([]float64, 0, len(s))
	for _, ele := range s {
		o, e := conv.Float64(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}

func InterfaceToBool(s []interface{}) ([]bool, error) {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil, e
		}
		ret = append(ret, o)
	}
	return ret, nil
}

func InterfaceToBoolMust(s []interface{}) []bool {
	ret := make([]bool, 0, len(s))
	for _, ele := range s {
		o, e := conv.Bool(ele)
		if e != nil {
			return nil
		}
		ret = append(ret, o)
	}
	return ret
}
