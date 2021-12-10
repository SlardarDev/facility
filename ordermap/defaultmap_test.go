package ordermap

import (
	"log"
	"testing"
)

func TestStringFloat32(t *testing.T) {

	k := map[string]int64{}
	k["a"] = 12
	k["c"] = 12
	k["d"] = 12
	k["e"] = 12
	k["f"] = 12
	k["v"] = 12

	iter := StringInt64Rev(k)
	for k, v, ok := iter(); ok; k, v, ok = iter() {
		log.Printf("%s, %d\n", k, v)
	}
}
