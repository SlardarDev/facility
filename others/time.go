package others

import "time"

const (
	TimeFormat_Format1 = "2006-01-02 15:04:05"
)

func TimeUnixNano2Format1(unixNano int64) string {
	t := time.Unix(unixNano/int64(time.Second), unixNano%int64(time.Second))
	return t.Format(TimeFormat_Format1)
}

func TimeFormat1ToTime(format string) (time.Time, error) {
	return time.Parse(TimeFormat_Format1, format)
}

func TimeTime2Format1(t time.Time) string {
	return t.Format(TimeFormat_Format1)
}

func UnixNano2Time(unixNano int64) time.Time {
	return time.Unix(unixNano/int64(time.Second), unixNano%int64(time.Second))
}
