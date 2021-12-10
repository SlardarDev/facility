package merge

import (
	"fmt"
	"testing"
	"time"
)

func Test_Merge(t *testing.T) {
	merger := NewDataMerger("Test", 2, 10, 1,
		100*time.Millisecond,
		func(data []interface{}) {
			t.Logf("OneResult: %v", data)
			fmt.Printf("OneResult: %v\n", data)
		})

	for i := 0; i < 100; i++ {
		merger.Send(i)
		time.Sleep(1 * time.Millisecond)

		if i%101 == 0 {
			time.Sleep(3300 * time.Millisecond)
		}
	}

	time.Sleep(2 * time.Second)
}
