package conv

import "time"

func TimeDurationPtr(arg time.Duration) *time.Duration {
	return &arg
}

func ByteSlicePtr(s []byte) *[]byte {
	return &s
}
