package conv

import "time"

func BoolDefault(from interface{}, d bool) bool {
	if v, err := converter.Bool(from); err == nil {
		return v
	}
	return d
}

func DurationDefault(from interface{}, d time.Duration) time.Duration {
	if v, err := converter.Duration(from); err == nil {
		return v
	}
	return d
}

// String will convert the given value to a string, returns the default value
// of "" if a conversion can not be made.
func StringDefault(from interface{}, d string) string {
	if from == nil {
		return d
	}
	if v, err := converter.String(from); err == nil {
		return v
	}
	return d
}

// Time will convert the given value to a time.Time, returns the empty struct
// time.Time{} if a conversion can not be made.
func TimeDefault(from interface{}, d time.Time) time.Time {
	if v, err := converter.Time(from); err == nil {
		return v
	}
	return d
}

// Float32 will convert the given value to a float32, returns the default value
// of 0.0 if a conversion can not be made.
func Float32Default(from interface{}, d float32) float32 {
	if v, err := converter.Float32(from); err == nil {
		return v
	}
	return d
}

// Float64 will convert the given value to a float64, returns the default value
// of 0.0 if a conversion can not be made.
func Float64Default(from interface{}, d float64) float64 {
	if v, err := converter.Float64(from); err == nil {
		return v
	}
	return d
}

// Int will convert the given value to a int, returns the default value of 0 if
// a conversion can not be made.
func IntDefault(from interface{}, d int) int {
	if v, err := converter.Int(from); err == nil {
		return v
	}
	return d
}

// Int8 will convert the given value to a int8, returns the default value of 0
// if a conversion can not be made.
func Int8Default(from interface{}, d int8) int8 {
	if v, err := converter.Int8(from); err == nil {
		return v
	}
	return d
}

// Int16 will convert the given value to a int16, returns the default value of 0
// if a conversion can not be made.
func Int16Default(from interface{}, d int16) int16 {
	if v, err := converter.Int16(from); err == nil {
		return v
	}
	return d
}

// Int32 will convert the given value to a int32, returns the default value of 0
// if a conversion can not be made.
func Int32Default(from interface{}, d int32) int32 {
	if v, err := converter.Int32(from); err == nil {
		return v
	}
	return d
}

// Int64 will convert the given value to a int64, returns the default value of 0
// if a conversion can not be made.
func Int64Default(from interface{}, d int64) int64 {
	if v, err := converter.Int64(from); err == nil {
		return v
	}
	return d
}

// Uint will convert the given value to a uint, returns the default value of 0
// if a conversion can not be made.
func UintDefault(from interface{}, d uint) uint {
	if v, err := converter.Uint(from); err == nil {
		return v
	}
	return d
}

// Uint8 will convert the given value to a uint8, returns the default value of 0
// if a conversion can not be made.
func Uint8Default(from interface{}, d uint8) uint8 {
	if v, err := converter.Uint8(from); err == nil {
		return v
	}
	return d
}

// Uint16 will convert the given value to a uint16, returns the default value of
// 0 if a conversion can not be made.
func Uint16Default(from interface{}, d uint16) uint16 {
	if v, err := converter.Uint16(from); err == nil {
		return v
	}
	return d
}

// Uint32 will convert the given value to a uint32, returns the default value of
// 0 if a conversion can not be made.
func Uint32Default(from interface{}, d uint32) uint32 {
	if v, err := converter.Uint32(from); err == nil {
		return v
	}
	return d
}

// Uint64 will convert the given value to a uint64, returns the default value of
// 0 if a conversion can not be made.
func Uint64Default(from interface{}, d uint64) uint64 {
	if v, err := converter.Uint64(from); err == nil {
		return v
	}
	return d
}
