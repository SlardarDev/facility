package fasttime

import (
	"sync/atomic"
	"time"
)

var unixtime int64

func init() {
	go func() {
		atomic.StoreInt64(&unixtime, time.Now().UnixNano())
		cnt := 0
		tick := time.Tick(time.Second)
		for {
			select {
			case <-tick:
				cnt++
				if cnt == 60 {
					cnt = 0
					atomic.StoreInt64(&unixtime, time.Now().UnixNano())
				} else {
					atomic.AddInt64(&unixtime, int64(time.Second))
				}
			}
		}
	}()
}

func Now() time.Time {
	return time.Unix(0, atomic.LoadInt64(&unixtime))
}
