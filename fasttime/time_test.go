package fasttime

import (
	"github.com/SlardarDev/facility/assert"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkNow1(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			time.Now()
		}
	})
}

func BenchmarkNow2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Now()
		}
	})
}

func TestTimeGap(t *testing.T) {

	for i := 0; i < 10; i++ {
		n := rand.Intn(2000)
		n = n + 1
		time.Sleep(time.Duration(n) * time.Millisecond)
		assert.EqualInt(int(Now().Sub(time.Now()).Seconds()),0)
		t.Logf("%d", int(Now().Sub(time.Now()).Seconds()))
	}

}
