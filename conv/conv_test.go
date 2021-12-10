package conv

import (
	"testing"
)

// 使用assert包有循环引用的问题, 改成if
func TestConv(t *testing.T) {
	i, _ := Int64("64")
	if i != 64 {
		t.Fail()
	}
	i = Int64Default("xxxx", 0)
	if i != 0 {
		t.Fail()
	}
	cc := Int64Ptr(10)
	if *cc != 10 {
		t.Fail()
	}
}
