package ordermap

import (
	"testing"
)

type testStruct struct {
	k int32
	v int64
}

func testInt32Int64() []*testStruct {
	var data []*testStruct = make([]*testStruct, 10)
	data[0] = &testStruct{0, 110}
	data[1] = &testStruct{2, 112}
	data[2] = &testStruct{4, 114}
	data[3] = &testStruct{6, 116}
	data[4] = &testStruct{8, 118}
	data[5] = &testStruct{1, 111}
	data[6] = &testStruct{3, 113}
	data[7] = &testStruct{5, 115}
	data[8] = &testStruct{7, 117}
	data[9] = &testStruct{9, 119}

	return data
}

func TestSetData(t *testing.T) {
	testCase := testInt32Int64()
	//t.Log(testCase)
	om := NewInt32Int64MapAllocSpace(10)
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}

	for _, kvp := range testCase {
		t.Log(kvp)
		om.Set(kvp.k, kvp.v)
	}

	if len(om.store) != len(testCase) {
		t.Error("Failed insert of args:", om.store, testCase)
	}

	for _, kvp := range testCase {
		val, ok := om.Get(kvp.k)
		if ok && kvp.v != val {
			t.Error(kvp.v, val)
		}
	}
	t.Log("1234")
}

func TestDeleteData(t *testing.T) {
	testCase := testInt32Int64()
	testKey1 := testCase[2].k
	testKey2 := testCase[9].k
	//t.Log(testCase)
	om := NewInt32Int64MapAllocSpace(10)
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}

	for _, kvp := range testCase {
		om.Set(kvp.k, kvp.v)
	}

	// Delete key
	om.Delete(testKey1)
	// Assert size equal to "test data size" - 1
	if om.Len() != (len(testCase) - 1) {
		t.Error("mapper size is incorrect")
	}
	if om.Len() != (len(testCase) - 1) {
		t.Error("store size is incorrect")
	}

	// Delete key
	om.Delete(testKey2)
	// Assert size equal to "test data size" - 1
	if om.Len() != (len(testCase) - 2) {
		t.Error("mapper size is incorrect")
	}
	if om.Len() != (len(testCase) - 2) {
		t.Error("store size is incorrect")
	}

	iter := om.IterFunc()
	for val, ok := iter(); ok; val, ok = iter() {
		t.Log(val)
	}
}

func TestInterData(t *testing.T) {
	testCase := testInt32Int64()
	//testKey := testCase[2].k
	//testIdx := 0
	//testExpected := []int64{110, 111, 112, 113, 114, 115, 116, 117, 118, 119}
	//t.Log(testCase)
	om := NewInt32Int64MapAllocSpace(10)
	if om == nil {
		t.Error("Failed to create OrderedMap")
	}

	for _, kvp := range testCase {
		om.Set(kvp.k, kvp.v)
	}

	iterAsc := om.IterFuncAsc()
	for val, ok := iterAsc(); ok; val, ok = iterAsc() {
		//if val != testExpected[idx] {
		//	t.Error(val, testExpected[idx])
		//}
		t.Log(val)
	}

	iter := om.IterFunc()
	idx := 0
	for val, ok := iter(); ok; val, ok = iter() {
		//if val != testExpected[idx] {
		//	t.Error(val, testExpected[idx])
		//}
		idx++
		t.Log(val)
	}

	iterDsc := om.IterFuncDesc()
	for val, ok := iterDsc(); ok; val, ok = iterDsc() {
		//if val != testExpected[idx] {
		//	t.Error(val, testExpected[idx])
		//}
		idx++
		t.Log(val)
	}

	//om.Delete(testCase[testIdx].k)
	//
	//testExpected = append(testExpected[:testIdx], testExpected[testIdx+1:]...)
	//t.Log(testExpected)
	//iterD := om.IterFunc()
	//idxD := 0
	//for val, ok := iterD(); ok; val, ok = iterD() {
	//	if val != testExpected[idxD] {
	//		t.Error(val, testExpected[idxD])
	//	}
	//	idxD++
	//	t.Log(val)
	//}

}
