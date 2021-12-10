package fjson

import (
	"encoding/json"
	"testing"
)

type S struct {
	Ttnet string `json:"ttnet"`
	Cc    string `json:"cc"`
}

/*
goos: darwin
goarch: amd64
pkg: github.com/SlardarDev/facility/fjson
5000000	       277 ns/op
PASS
*/
func BenchmarkLoopsForfjson(b *testing.B) {
	var j = []byte(`{"ttnet": "123", "cc": "mk"}`)
	for i := 0; i < b.N; i++ {
		var s S
		ConfigCompatibleWithStandardLibrary.Unmarshal(j, &s)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/SlardarDev/facility/fjson
10000000	       153 ns/op
PASS
*/
func BenchmarkLoopsParallelForfjson(b *testing.B) {
	var j = []byte(`{"ttnet": "123", "cc": "mk"}`)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var s S
			ConfigCompatibleWithStandardLibrary.Unmarshal(j, &s)
		}
	})
}

/*
goos: darwin
goarch: amd64
pkg: github.com/SlardarDev/facility/fjson
2000000	       759 ns/op
PASS
*/
func BenchmarkLoopsForNormalJson(b *testing.B) {
	var j = []byte(`{"ttnet": "123", "cc": "mk"}`)
	for i := 0; i < b.N; i++ {
		var s S
		json.Unmarshal(j, &s)
	}
}

/*
goos: darwin
goarch: amd64
pkg: github.com/SlardarDev/facility/fjson
3000000	       503 ns/op
PASS
*/
func BenchmarkLoopsParallelForNormalJson(b *testing.B) {
	var j = []byte(`{"ttnet": "123", "cc": "mk"}`)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var s S
			json.Unmarshal(j, &s)
		}
	})
}
