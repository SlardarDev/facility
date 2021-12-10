package simplettl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readDataFromDoc(path string) []int {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer f.Close()
	checkError(err)

	rd := bufio.NewReader(f)

	data := []int{}

	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		num, _ := strconv.Atoi(strings.Replace(line, "\n", "", -1))
		data = append(data, num)
	}
	return data
}

func countHitTimes(data []int) (ratio float32) {
	tmpTtl := NewSimpleTTLCache("tmpLru")
	var count float32 = 0
	var total float32 = 0
	for _, key := range data {
		total++
		if _, hit := tmpTtl.Get(key); hit {
			count++
		} else {
			tmpTtl.Set(key, key, time.Second*5)
		}
	}
	ratio = count / total
	fmt.Println(ratio)
	return ratio
}

func BenchmarkLockedLRU_Get(b *testing.B) {
	path := "/Users/zhanghuabo/go/src/generateTestData/NormDistributeData"
	testData := readDataFromDoc(path)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			countHitTimes(testData)
		}
	})
}
