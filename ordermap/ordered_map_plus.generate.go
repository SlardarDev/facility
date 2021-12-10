package ordermap

import "github.com/SlardarDev/facility/sorts"

type IntIntMap struct {
	store   map[int]int
	mapper  []int
	isWrite bool
}

func NewIntIntMap() *IntIntMap {
	om := &IntIntMap{
		store:   make(map[int]int),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntIntMapAllocSpace(size int) *IntIntMap {
	om := &IntIntMap{
		store:   make(map[int]int, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntIntMap) Set(key int, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntIntMap) Get(key int) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntIntMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntIntMap) Len() int {
	return len(om.store)
}

func (om *IntIntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntIntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntIntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntInt8Map struct {
	store   map[int]int8
	mapper  []int
	isWrite bool
}

func NewIntInt8Map() *IntInt8Map {
	om := &IntInt8Map{
		store:   make(map[int]int8),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntInt8MapAllocSpace(size int) *IntInt8Map {
	om := &IntInt8Map{
		store:   make(map[int]int8, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntInt8Map) Set(key int, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntInt8Map) Get(key int) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntInt8Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntInt8Map) Len() int {
	return len(om.store)
}

func (om *IntInt8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntInt8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntInt8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntInt16Map struct {
	store   map[int]int16
	mapper  []int
	isWrite bool
}

func NewIntInt16Map() *IntInt16Map {
	om := &IntInt16Map{
		store:   make(map[int]int16),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntInt16MapAllocSpace(size int) *IntInt16Map {
	om := &IntInt16Map{
		store:   make(map[int]int16, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntInt16Map) Set(key int, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntInt16Map) Get(key int) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntInt16Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntInt16Map) Len() int {
	return len(om.store)
}

func (om *IntInt16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntInt16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntInt16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntInt32Map struct {
	store   map[int]int32
	mapper  []int
	isWrite bool
}

func NewIntInt32Map() *IntInt32Map {
	om := &IntInt32Map{
		store:   make(map[int]int32),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntInt32MapAllocSpace(size int) *IntInt32Map {
	om := &IntInt32Map{
		store:   make(map[int]int32, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntInt32Map) Set(key int, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntInt32Map) Get(key int) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntInt32Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntInt32Map) Len() int {
	return len(om.store)
}

func (om *IntInt32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntInt32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntInt32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntInt64Map struct {
	store   map[int]int64
	mapper  []int
	isWrite bool
}

func NewIntInt64Map() *IntInt64Map {
	om := &IntInt64Map{
		store:   make(map[int]int64),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntInt64MapAllocSpace(size int) *IntInt64Map {
	om := &IntInt64Map{
		store:   make(map[int]int64, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntInt64Map) Set(key int, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntInt64Map) Get(key int) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntInt64Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntInt64Map) Len() int {
	return len(om.store)
}

func (om *IntInt64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntInt64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntInt64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntUint8Map struct {
	store   map[int]uint8
	mapper  []int
	isWrite bool
}

func NewIntUint8Map() *IntUint8Map {
	om := &IntUint8Map{
		store:   make(map[int]uint8),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntUint8MapAllocSpace(size int) *IntUint8Map {
	om := &IntUint8Map{
		store:   make(map[int]uint8, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntUint8Map) Set(key int, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntUint8Map) Get(key int) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntUint8Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntUint8Map) Len() int {
	return len(om.store)
}

func (om *IntUint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntUint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntUint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntUint16Map struct {
	store   map[int]uint16
	mapper  []int
	isWrite bool
}

func NewIntUint16Map() *IntUint16Map {
	om := &IntUint16Map{
		store:   make(map[int]uint16),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntUint16MapAllocSpace(size int) *IntUint16Map {
	om := &IntUint16Map{
		store:   make(map[int]uint16, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntUint16Map) Set(key int, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntUint16Map) Get(key int) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntUint16Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntUint16Map) Len() int {
	return len(om.store)
}

func (om *IntUint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntUint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntUint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntUint32Map struct {
	store   map[int]uint32
	mapper  []int
	isWrite bool
}

func NewIntUint32Map() *IntUint32Map {
	om := &IntUint32Map{
		store:   make(map[int]uint32),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntUint32MapAllocSpace(size int) *IntUint32Map {
	om := &IntUint32Map{
		store:   make(map[int]uint32, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntUint32Map) Set(key int, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntUint32Map) Get(key int) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntUint32Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntUint32Map) Len() int {
	return len(om.store)
}

func (om *IntUint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntUint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntUint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntUint64Map struct {
	store   map[int]uint64
	mapper  []int
	isWrite bool
}

func NewIntUint64Map() *IntUint64Map {
	om := &IntUint64Map{
		store:   make(map[int]uint64),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntUint64MapAllocSpace(size int) *IntUint64Map {
	om := &IntUint64Map{
		store:   make(map[int]uint64, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntUint64Map) Set(key int, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntUint64Map) Get(key int) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntUint64Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntUint64Map) Len() int {
	return len(om.store)
}

func (om *IntUint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntUint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntUint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntStringMap struct {
	store   map[int]string
	mapper  []int
	isWrite bool
}

func NewIntStringMap() *IntStringMap {
	om := &IntStringMap{
		store:   make(map[int]string),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntStringMapAllocSpace(size int) *IntStringMap {
	om := &IntStringMap{
		store:   make(map[int]string, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntStringMap) Set(key int, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntStringMap) Get(key int) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntStringMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntStringMap) Len() int {
	return len(om.store)
}

func (om *IntStringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntStringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntStringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntFloat32Map struct {
	store   map[int]float32
	mapper  []int
	isWrite bool
}

func NewIntFloat32Map() *IntFloat32Map {
	om := &IntFloat32Map{
		store:   make(map[int]float32),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntFloat32MapAllocSpace(size int) *IntFloat32Map {
	om := &IntFloat32Map{
		store:   make(map[int]float32, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntFloat32Map) Set(key int, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntFloat32Map) Get(key int) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntFloat32Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntFloat32Map) Len() int {
	return len(om.store)
}

func (om *IntFloat32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntFloat32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntFloat32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntFloat64Map struct {
	store   map[int]float64
	mapper  []int
	isWrite bool
}

func NewIntFloat64Map() *IntFloat64Map {
	om := &IntFloat64Map{
		store:   make(map[int]float64),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntFloat64MapAllocSpace(size int) *IntFloat64Map {
	om := &IntFloat64Map{
		store:   make(map[int]float64, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntFloat64Map) Set(key int, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntFloat64Map) Get(key int) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntFloat64Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntFloat64Map) Len() int {
	return len(om.store)
}

func (om *IntFloat64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntFloat64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntFloat64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntInterfaceMap struct {
	store   map[int]interface{}
	mapper  []int
	isWrite bool
}

func NewIntInterfaceMap() *IntInterfaceMap {
	om := &IntInterfaceMap{
		store:   make(map[int]interface{}),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntInterfaceMapAllocSpace(size int) *IntInterfaceMap {
	om := &IntInterfaceMap{
		store:   make(map[int]interface{}, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntInterfaceMap) Set(key int, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntInterfaceMap) Get(key int) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntInterfaceMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntInterfaceMap) Len() int {
	return len(om.store)
}

func (om *IntInterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntInterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntInterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntBoolMap struct {
	store   map[int]bool
	mapper  []int
	isWrite bool
}

func NewIntBoolMap() *IntBoolMap {
	om := &IntBoolMap{
		store:   make(map[int]bool),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntBoolMapAllocSpace(size int) *IntBoolMap {
	om := &IntBoolMap{
		store:   make(map[int]bool, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntBoolMap) Set(key int, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntBoolMap) Get(key int) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntBoolMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntBoolMap) Len() int {
	return len(om.store)
}

func (om *IntBoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntBoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntBoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntStructMap struct {
	store   map[int]struct{}
	mapper  []int
	isWrite bool
}

func NewIntStructMap() *IntStructMap {
	om := &IntStructMap{
		store:   make(map[int]struct{}),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntStructMapAllocSpace(size int) *IntStructMap {
	om := &IntStructMap{
		store:   make(map[int]struct{}, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntStructMap) Set(key int, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntStructMap) Get(key int) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntStructMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntStructMap) Len() int {
	return len(om.store)
}

func (om *IntStructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntStructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntStructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceintMap struct {
	store   map[int][]int
	mapper  []int
	isWrite bool
}

func NewIntSliceintMap() *IntSliceintMap {
	om := &IntSliceintMap{
		store:   make(map[int][]int),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceintMapAllocSpace(size int) *IntSliceintMap {
	om := &IntSliceintMap{
		store:   make(map[int][]int, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceintMap) Set(key int, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceintMap) Get(key int) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceintMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceintMap) Len() int {
	return len(om.store)
}

func (om *IntSliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceint8Map struct {
	store   map[int][]int8
	mapper  []int
	isWrite bool
}

func NewIntSliceint8Map() *IntSliceint8Map {
	om := &IntSliceint8Map{
		store:   make(map[int][]int8),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceint8MapAllocSpace(size int) *IntSliceint8Map {
	om := &IntSliceint8Map{
		store:   make(map[int][]int8, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceint8Map) Set(key int, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceint8Map) Get(key int) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceint8Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceint8Map) Len() int {
	return len(om.store)
}

func (om *IntSliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceint16Map struct {
	store   map[int][]int16
	mapper  []int
	isWrite bool
}

func NewIntSliceint16Map() *IntSliceint16Map {
	om := &IntSliceint16Map{
		store:   make(map[int][]int16),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceint16MapAllocSpace(size int) *IntSliceint16Map {
	om := &IntSliceint16Map{
		store:   make(map[int][]int16, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceint16Map) Set(key int, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceint16Map) Get(key int) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceint16Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceint16Map) Len() int {
	return len(om.store)
}

func (om *IntSliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceint32Map struct {
	store   map[int][]int32
	mapper  []int
	isWrite bool
}

func NewIntSliceint32Map() *IntSliceint32Map {
	om := &IntSliceint32Map{
		store:   make(map[int][]int32),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceint32MapAllocSpace(size int) *IntSliceint32Map {
	om := &IntSliceint32Map{
		store:   make(map[int][]int32, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceint32Map) Set(key int, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceint32Map) Get(key int) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceint32Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceint32Map) Len() int {
	return len(om.store)
}

func (om *IntSliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceint64Map struct {
	store   map[int][]int64
	mapper  []int
	isWrite bool
}

func NewIntSliceint64Map() *IntSliceint64Map {
	om := &IntSliceint64Map{
		store:   make(map[int][]int64),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceint64MapAllocSpace(size int) *IntSliceint64Map {
	om := &IntSliceint64Map{
		store:   make(map[int][]int64, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceint64Map) Set(key int, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceint64Map) Get(key int) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceint64Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceint64Map) Len() int {
	return len(om.store)
}

func (om *IntSliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceuint8Map struct {
	store   map[int][]uint8
	mapper  []int
	isWrite bool
}

func NewIntSliceuint8Map() *IntSliceuint8Map {
	om := &IntSliceuint8Map{
		store:   make(map[int][]uint8),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceuint8MapAllocSpace(size int) *IntSliceuint8Map {
	om := &IntSliceuint8Map{
		store:   make(map[int][]uint8, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceuint8Map) Set(key int, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceuint8Map) Get(key int) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceuint8Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceuint8Map) Len() int {
	return len(om.store)
}

func (om *IntSliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceuint16Map struct {
	store   map[int][]uint16
	mapper  []int
	isWrite bool
}

func NewIntSliceuint16Map() *IntSliceuint16Map {
	om := &IntSliceuint16Map{
		store:   make(map[int][]uint16),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceuint16MapAllocSpace(size int) *IntSliceuint16Map {
	om := &IntSliceuint16Map{
		store:   make(map[int][]uint16, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceuint16Map) Set(key int, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceuint16Map) Get(key int) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceuint16Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceuint16Map) Len() int {
	return len(om.store)
}

func (om *IntSliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceuint32Map struct {
	store   map[int][]uint32
	mapper  []int
	isWrite bool
}

func NewIntSliceuint32Map() *IntSliceuint32Map {
	om := &IntSliceuint32Map{
		store:   make(map[int][]uint32),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceuint32MapAllocSpace(size int) *IntSliceuint32Map {
	om := &IntSliceuint32Map{
		store:   make(map[int][]uint32, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceuint32Map) Set(key int, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceuint32Map) Get(key int) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceuint32Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceuint32Map) Len() int {
	return len(om.store)
}

func (om *IntSliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSliceuint64Map struct {
	store   map[int][]uint64
	mapper  []int
	isWrite bool
}

func NewIntSliceuint64Map() *IntSliceuint64Map {
	om := &IntSliceuint64Map{
		store:   make(map[int][]uint64),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSliceuint64MapAllocSpace(size int) *IntSliceuint64Map {
	om := &IntSliceuint64Map{
		store:   make(map[int][]uint64, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSliceuint64Map) Set(key int, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSliceuint64Map) Get(key int) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSliceuint64Map) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSliceuint64Map) Len() int {
	return len(om.store)
}

func (om *IntSliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type IntSlicestringMap struct {
	store   map[int][]string
	mapper  []int
	isWrite bool
}

func NewIntSlicestringMap() *IntSlicestringMap {
	om := &IntSlicestringMap{
		store:   make(map[int][]string),
		mapper:  make([]int, 0),
		isWrite: false,
	}
	return om
}

func NewIntSlicestringMapAllocSpace(size int) *IntSlicestringMap {
	om := &IntSlicestringMap{
		store:   make(map[int][]string, size),
		mapper:  make([]int, 0, size),
		isWrite: false,
	}
	return om
}

func (om *IntSlicestringMap) Set(key int, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *IntSlicestringMap) Get(key int) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *IntSlicestringMap) Delete(key int) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *IntSlicestringMap) Len() int {
	return len(om.store)
}

func (om *IntSlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *IntSlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]int, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Ints(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *IntSlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]int, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.IntsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8IntMap struct {
	store   map[int8]int
	mapper  []int8
	isWrite bool
}

func NewInt8IntMap() *Int8IntMap {
	om := &Int8IntMap{
		store:   make(map[int8]int),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8IntMapAllocSpace(size int) *Int8IntMap {
	om := &Int8IntMap{
		store:   make(map[int8]int, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8IntMap) Set(key int8, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8IntMap) Get(key int8) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8IntMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8IntMap) Len() int {
	return len(om.store)
}

func (om *Int8IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Int8Map struct {
	store   map[int8]int8
	mapper  []int8
	isWrite bool
}

func NewInt8Int8Map() *Int8Int8Map {
	om := &Int8Int8Map{
		store:   make(map[int8]int8),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Int8MapAllocSpace(size int) *Int8Int8Map {
	om := &Int8Int8Map{
		store:   make(map[int8]int8, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Int8Map) Set(key int8, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Int8Map) Get(key int8) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Int8Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Int8Map) Len() int {
	return len(om.store)
}

func (om *Int8Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Int16Map struct {
	store   map[int8]int16
	mapper  []int8
	isWrite bool
}

func NewInt8Int16Map() *Int8Int16Map {
	om := &Int8Int16Map{
		store:   make(map[int8]int16),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Int16MapAllocSpace(size int) *Int8Int16Map {
	om := &Int8Int16Map{
		store:   make(map[int8]int16, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Int16Map) Set(key int8, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Int16Map) Get(key int8) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Int16Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Int16Map) Len() int {
	return len(om.store)
}

func (om *Int8Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Int32Map struct {
	store   map[int8]int32
	mapper  []int8
	isWrite bool
}

func NewInt8Int32Map() *Int8Int32Map {
	om := &Int8Int32Map{
		store:   make(map[int8]int32),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Int32MapAllocSpace(size int) *Int8Int32Map {
	om := &Int8Int32Map{
		store:   make(map[int8]int32, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Int32Map) Set(key int8, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Int32Map) Get(key int8) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Int32Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Int32Map) Len() int {
	return len(om.store)
}

func (om *Int8Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Int64Map struct {
	store   map[int8]int64
	mapper  []int8
	isWrite bool
}

func NewInt8Int64Map() *Int8Int64Map {
	om := &Int8Int64Map{
		store:   make(map[int8]int64),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Int64MapAllocSpace(size int) *Int8Int64Map {
	om := &Int8Int64Map{
		store:   make(map[int8]int64, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Int64Map) Set(key int8, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Int64Map) Get(key int8) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Int64Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Int64Map) Len() int {
	return len(om.store)
}

func (om *Int8Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Uint8Map struct {
	store   map[int8]uint8
	mapper  []int8
	isWrite bool
}

func NewInt8Uint8Map() *Int8Uint8Map {
	om := &Int8Uint8Map{
		store:   make(map[int8]uint8),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Uint8MapAllocSpace(size int) *Int8Uint8Map {
	om := &Int8Uint8Map{
		store:   make(map[int8]uint8, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Uint8Map) Set(key int8, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Uint8Map) Get(key int8) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Uint8Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Uint8Map) Len() int {
	return len(om.store)
}

func (om *Int8Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Uint16Map struct {
	store   map[int8]uint16
	mapper  []int8
	isWrite bool
}

func NewInt8Uint16Map() *Int8Uint16Map {
	om := &Int8Uint16Map{
		store:   make(map[int8]uint16),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Uint16MapAllocSpace(size int) *Int8Uint16Map {
	om := &Int8Uint16Map{
		store:   make(map[int8]uint16, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Uint16Map) Set(key int8, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Uint16Map) Get(key int8) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Uint16Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Uint16Map) Len() int {
	return len(om.store)
}

func (om *Int8Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Uint32Map struct {
	store   map[int8]uint32
	mapper  []int8
	isWrite bool
}

func NewInt8Uint32Map() *Int8Uint32Map {
	om := &Int8Uint32Map{
		store:   make(map[int8]uint32),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Uint32MapAllocSpace(size int) *Int8Uint32Map {
	om := &Int8Uint32Map{
		store:   make(map[int8]uint32, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Uint32Map) Set(key int8, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Uint32Map) Get(key int8) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Uint32Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Uint32Map) Len() int {
	return len(om.store)
}

func (om *Int8Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Uint64Map struct {
	store   map[int8]uint64
	mapper  []int8
	isWrite bool
}

func NewInt8Uint64Map() *Int8Uint64Map {
	om := &Int8Uint64Map{
		store:   make(map[int8]uint64),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Uint64MapAllocSpace(size int) *Int8Uint64Map {
	om := &Int8Uint64Map{
		store:   make(map[int8]uint64, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Uint64Map) Set(key int8, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Uint64Map) Get(key int8) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Uint64Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Uint64Map) Len() int {
	return len(om.store)
}

func (om *Int8Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8StringMap struct {
	store   map[int8]string
	mapper  []int8
	isWrite bool
}

func NewInt8StringMap() *Int8StringMap {
	om := &Int8StringMap{
		store:   make(map[int8]string),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8StringMapAllocSpace(size int) *Int8StringMap {
	om := &Int8StringMap{
		store:   make(map[int8]string, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8StringMap) Set(key int8, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8StringMap) Get(key int8) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8StringMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8StringMap) Len() int {
	return len(om.store)
}

func (om *Int8StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Float32Map struct {
	store   map[int8]float32
	mapper  []int8
	isWrite bool
}

func NewInt8Float32Map() *Int8Float32Map {
	om := &Int8Float32Map{
		store:   make(map[int8]float32),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Float32MapAllocSpace(size int) *Int8Float32Map {
	om := &Int8Float32Map{
		store:   make(map[int8]float32, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Float32Map) Set(key int8, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Float32Map) Get(key int8) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Float32Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Float32Map) Len() int {
	return len(om.store)
}

func (om *Int8Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Float64Map struct {
	store   map[int8]float64
	mapper  []int8
	isWrite bool
}

func NewInt8Float64Map() *Int8Float64Map {
	om := &Int8Float64Map{
		store:   make(map[int8]float64),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Float64MapAllocSpace(size int) *Int8Float64Map {
	om := &Int8Float64Map{
		store:   make(map[int8]float64, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Float64Map) Set(key int8, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Float64Map) Get(key int8) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Float64Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Float64Map) Len() int {
	return len(om.store)
}

func (om *Int8Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8InterfaceMap struct {
	store   map[int8]interface{}
	mapper  []int8
	isWrite bool
}

func NewInt8InterfaceMap() *Int8InterfaceMap {
	om := &Int8InterfaceMap{
		store:   make(map[int8]interface{}),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8InterfaceMapAllocSpace(size int) *Int8InterfaceMap {
	om := &Int8InterfaceMap{
		store:   make(map[int8]interface{}, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8InterfaceMap) Set(key int8, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8InterfaceMap) Get(key int8) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8InterfaceMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Int8InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8BoolMap struct {
	store   map[int8]bool
	mapper  []int8
	isWrite bool
}

func NewInt8BoolMap() *Int8BoolMap {
	om := &Int8BoolMap{
		store:   make(map[int8]bool),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8BoolMapAllocSpace(size int) *Int8BoolMap {
	om := &Int8BoolMap{
		store:   make(map[int8]bool, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8BoolMap) Set(key int8, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8BoolMap) Get(key int8) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8BoolMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8BoolMap) Len() int {
	return len(om.store)
}

func (om *Int8BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8StructMap struct {
	store   map[int8]struct{}
	mapper  []int8
	isWrite bool
}

func NewInt8StructMap() *Int8StructMap {
	om := &Int8StructMap{
		store:   make(map[int8]struct{}),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8StructMapAllocSpace(size int) *Int8StructMap {
	om := &Int8StructMap{
		store:   make(map[int8]struct{}, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8StructMap) Set(key int8, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8StructMap) Get(key int8) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8StructMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8StructMap) Len() int {
	return len(om.store)
}

func (om *Int8StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8SliceintMap struct {
	store   map[int8][]int
	mapper  []int8
	isWrite bool
}

func NewInt8SliceintMap() *Int8SliceintMap {
	om := &Int8SliceintMap{
		store:   make(map[int8][]int),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8SliceintMapAllocSpace(size int) *Int8SliceintMap {
	om := &Int8SliceintMap{
		store:   make(map[int8][]int, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8SliceintMap) Set(key int8, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8SliceintMap) Get(key int8) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8SliceintMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8SliceintMap) Len() int {
	return len(om.store)
}

func (om *Int8SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceint8Map struct {
	store   map[int8][]int8
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceint8Map() *Int8Sliceint8Map {
	om := &Int8Sliceint8Map{
		store:   make(map[int8][]int8),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceint8MapAllocSpace(size int) *Int8Sliceint8Map {
	om := &Int8Sliceint8Map{
		store:   make(map[int8][]int8, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceint8Map) Set(key int8, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceint8Map) Get(key int8) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceint8Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceint16Map struct {
	store   map[int8][]int16
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceint16Map() *Int8Sliceint16Map {
	om := &Int8Sliceint16Map{
		store:   make(map[int8][]int16),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceint16MapAllocSpace(size int) *Int8Sliceint16Map {
	om := &Int8Sliceint16Map{
		store:   make(map[int8][]int16, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceint16Map) Set(key int8, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceint16Map) Get(key int8) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceint16Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceint32Map struct {
	store   map[int8][]int32
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceint32Map() *Int8Sliceint32Map {
	om := &Int8Sliceint32Map{
		store:   make(map[int8][]int32),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceint32MapAllocSpace(size int) *Int8Sliceint32Map {
	om := &Int8Sliceint32Map{
		store:   make(map[int8][]int32, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceint32Map) Set(key int8, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceint32Map) Get(key int8) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceint32Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceint64Map struct {
	store   map[int8][]int64
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceint64Map() *Int8Sliceint64Map {
	om := &Int8Sliceint64Map{
		store:   make(map[int8][]int64),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceint64MapAllocSpace(size int) *Int8Sliceint64Map {
	om := &Int8Sliceint64Map{
		store:   make(map[int8][]int64, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceint64Map) Set(key int8, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceint64Map) Get(key int8) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceint64Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceuint8Map struct {
	store   map[int8][]uint8
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceuint8Map() *Int8Sliceuint8Map {
	om := &Int8Sliceuint8Map{
		store:   make(map[int8][]uint8),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceuint8MapAllocSpace(size int) *Int8Sliceuint8Map {
	om := &Int8Sliceuint8Map{
		store:   make(map[int8][]uint8, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceuint8Map) Set(key int8, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceuint8Map) Get(key int8) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceuint8Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceuint16Map struct {
	store   map[int8][]uint16
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceuint16Map() *Int8Sliceuint16Map {
	om := &Int8Sliceuint16Map{
		store:   make(map[int8][]uint16),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceuint16MapAllocSpace(size int) *Int8Sliceuint16Map {
	om := &Int8Sliceuint16Map{
		store:   make(map[int8][]uint16, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceuint16Map) Set(key int8, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceuint16Map) Get(key int8) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceuint16Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceuint32Map struct {
	store   map[int8][]uint32
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceuint32Map() *Int8Sliceuint32Map {
	om := &Int8Sliceuint32Map{
		store:   make(map[int8][]uint32),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceuint32MapAllocSpace(size int) *Int8Sliceuint32Map {
	om := &Int8Sliceuint32Map{
		store:   make(map[int8][]uint32, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceuint32Map) Set(key int8, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceuint32Map) Get(key int8) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceuint32Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8Sliceuint64Map struct {
	store   map[int8][]uint64
	mapper  []int8
	isWrite bool
}

func NewInt8Sliceuint64Map() *Int8Sliceuint64Map {
	om := &Int8Sliceuint64Map{
		store:   make(map[int8][]uint64),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8Sliceuint64MapAllocSpace(size int) *Int8Sliceuint64Map {
	om := &Int8Sliceuint64Map{
		store:   make(map[int8][]uint64, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8Sliceuint64Map) Set(key int8, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8Sliceuint64Map) Get(key int8) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8Sliceuint64Map) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Int8Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int8SlicestringMap struct {
	store   map[int8][]string
	mapper  []int8
	isWrite bool
}

func NewInt8SlicestringMap() *Int8SlicestringMap {
	om := &Int8SlicestringMap{
		store:   make(map[int8][]string),
		mapper:  make([]int8, 0),
		isWrite: false,
	}
	return om
}

func NewInt8SlicestringMapAllocSpace(size int) *Int8SlicestringMap {
	om := &Int8SlicestringMap{
		store:   make(map[int8][]string, size),
		mapper:  make([]int8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int8SlicestringMap) Set(key int8, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int8SlicestringMap) Get(key int8) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int8SlicestringMap) Delete(key int8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int8SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Int8SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int8SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]int8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int8SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]int8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16IntMap struct {
	store   map[int16]int
	mapper  []int16
	isWrite bool
}

func NewInt16IntMap() *Int16IntMap {
	om := &Int16IntMap{
		store:   make(map[int16]int),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16IntMapAllocSpace(size int) *Int16IntMap {
	om := &Int16IntMap{
		store:   make(map[int16]int, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16IntMap) Set(key int16, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16IntMap) Get(key int16) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16IntMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16IntMap) Len() int {
	return len(om.store)
}

func (om *Int16IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Int8Map struct {
	store   map[int16]int8
	mapper  []int16
	isWrite bool
}

func NewInt16Int8Map() *Int16Int8Map {
	om := &Int16Int8Map{
		store:   make(map[int16]int8),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Int8MapAllocSpace(size int) *Int16Int8Map {
	om := &Int16Int8Map{
		store:   make(map[int16]int8, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Int8Map) Set(key int16, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Int8Map) Get(key int16) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Int8Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Int8Map) Len() int {
	return len(om.store)
}

func (om *Int16Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Int16Map struct {
	store   map[int16]int16
	mapper  []int16
	isWrite bool
}

func NewInt16Int16Map() *Int16Int16Map {
	om := &Int16Int16Map{
		store:   make(map[int16]int16),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Int16MapAllocSpace(size int) *Int16Int16Map {
	om := &Int16Int16Map{
		store:   make(map[int16]int16, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Int16Map) Set(key int16, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Int16Map) Get(key int16) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Int16Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Int16Map) Len() int {
	return len(om.store)
}

func (om *Int16Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Int32Map struct {
	store   map[int16]int32
	mapper  []int16
	isWrite bool
}

func NewInt16Int32Map() *Int16Int32Map {
	om := &Int16Int32Map{
		store:   make(map[int16]int32),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Int32MapAllocSpace(size int) *Int16Int32Map {
	om := &Int16Int32Map{
		store:   make(map[int16]int32, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Int32Map) Set(key int16, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Int32Map) Get(key int16) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Int32Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Int32Map) Len() int {
	return len(om.store)
}

func (om *Int16Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Int64Map struct {
	store   map[int16]int64
	mapper  []int16
	isWrite bool
}

func NewInt16Int64Map() *Int16Int64Map {
	om := &Int16Int64Map{
		store:   make(map[int16]int64),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Int64MapAllocSpace(size int) *Int16Int64Map {
	om := &Int16Int64Map{
		store:   make(map[int16]int64, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Int64Map) Set(key int16, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Int64Map) Get(key int16) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Int64Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Int64Map) Len() int {
	return len(om.store)
}

func (om *Int16Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Uint8Map struct {
	store   map[int16]uint8
	mapper  []int16
	isWrite bool
}

func NewInt16Uint8Map() *Int16Uint8Map {
	om := &Int16Uint8Map{
		store:   make(map[int16]uint8),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Uint8MapAllocSpace(size int) *Int16Uint8Map {
	om := &Int16Uint8Map{
		store:   make(map[int16]uint8, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Uint8Map) Set(key int16, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Uint8Map) Get(key int16) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Uint8Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Uint8Map) Len() int {
	return len(om.store)
}

func (om *Int16Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Uint16Map struct {
	store   map[int16]uint16
	mapper  []int16
	isWrite bool
}

func NewInt16Uint16Map() *Int16Uint16Map {
	om := &Int16Uint16Map{
		store:   make(map[int16]uint16),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Uint16MapAllocSpace(size int) *Int16Uint16Map {
	om := &Int16Uint16Map{
		store:   make(map[int16]uint16, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Uint16Map) Set(key int16, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Uint16Map) Get(key int16) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Uint16Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Uint16Map) Len() int {
	return len(om.store)
}

func (om *Int16Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Uint32Map struct {
	store   map[int16]uint32
	mapper  []int16
	isWrite bool
}

func NewInt16Uint32Map() *Int16Uint32Map {
	om := &Int16Uint32Map{
		store:   make(map[int16]uint32),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Uint32MapAllocSpace(size int) *Int16Uint32Map {
	om := &Int16Uint32Map{
		store:   make(map[int16]uint32, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Uint32Map) Set(key int16, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Uint32Map) Get(key int16) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Uint32Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Uint32Map) Len() int {
	return len(om.store)
}

func (om *Int16Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Uint64Map struct {
	store   map[int16]uint64
	mapper  []int16
	isWrite bool
}

func NewInt16Uint64Map() *Int16Uint64Map {
	om := &Int16Uint64Map{
		store:   make(map[int16]uint64),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Uint64MapAllocSpace(size int) *Int16Uint64Map {
	om := &Int16Uint64Map{
		store:   make(map[int16]uint64, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Uint64Map) Set(key int16, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Uint64Map) Get(key int16) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Uint64Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Uint64Map) Len() int {
	return len(om.store)
}

func (om *Int16Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16StringMap struct {
	store   map[int16]string
	mapper  []int16
	isWrite bool
}

func NewInt16StringMap() *Int16StringMap {
	om := &Int16StringMap{
		store:   make(map[int16]string),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16StringMapAllocSpace(size int) *Int16StringMap {
	om := &Int16StringMap{
		store:   make(map[int16]string, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16StringMap) Set(key int16, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16StringMap) Get(key int16) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16StringMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16StringMap) Len() int {
	return len(om.store)
}

func (om *Int16StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Float32Map struct {
	store   map[int16]float32
	mapper  []int16
	isWrite bool
}

func NewInt16Float32Map() *Int16Float32Map {
	om := &Int16Float32Map{
		store:   make(map[int16]float32),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Float32MapAllocSpace(size int) *Int16Float32Map {
	om := &Int16Float32Map{
		store:   make(map[int16]float32, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Float32Map) Set(key int16, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Float32Map) Get(key int16) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Float32Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Float32Map) Len() int {
	return len(om.store)
}

func (om *Int16Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Float64Map struct {
	store   map[int16]float64
	mapper  []int16
	isWrite bool
}

func NewInt16Float64Map() *Int16Float64Map {
	om := &Int16Float64Map{
		store:   make(map[int16]float64),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Float64MapAllocSpace(size int) *Int16Float64Map {
	om := &Int16Float64Map{
		store:   make(map[int16]float64, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Float64Map) Set(key int16, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Float64Map) Get(key int16) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Float64Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Float64Map) Len() int {
	return len(om.store)
}

func (om *Int16Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16InterfaceMap struct {
	store   map[int16]interface{}
	mapper  []int16
	isWrite bool
}

func NewInt16InterfaceMap() *Int16InterfaceMap {
	om := &Int16InterfaceMap{
		store:   make(map[int16]interface{}),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16InterfaceMapAllocSpace(size int) *Int16InterfaceMap {
	om := &Int16InterfaceMap{
		store:   make(map[int16]interface{}, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16InterfaceMap) Set(key int16, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16InterfaceMap) Get(key int16) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16InterfaceMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Int16InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16BoolMap struct {
	store   map[int16]bool
	mapper  []int16
	isWrite bool
}

func NewInt16BoolMap() *Int16BoolMap {
	om := &Int16BoolMap{
		store:   make(map[int16]bool),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16BoolMapAllocSpace(size int) *Int16BoolMap {
	om := &Int16BoolMap{
		store:   make(map[int16]bool, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16BoolMap) Set(key int16, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16BoolMap) Get(key int16) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16BoolMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16BoolMap) Len() int {
	return len(om.store)
}

func (om *Int16BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16StructMap struct {
	store   map[int16]struct{}
	mapper  []int16
	isWrite bool
}

func NewInt16StructMap() *Int16StructMap {
	om := &Int16StructMap{
		store:   make(map[int16]struct{}),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16StructMapAllocSpace(size int) *Int16StructMap {
	om := &Int16StructMap{
		store:   make(map[int16]struct{}, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16StructMap) Set(key int16, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16StructMap) Get(key int16) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16StructMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16StructMap) Len() int {
	return len(om.store)
}

func (om *Int16StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16SliceintMap struct {
	store   map[int16][]int
	mapper  []int16
	isWrite bool
}

func NewInt16SliceintMap() *Int16SliceintMap {
	om := &Int16SliceintMap{
		store:   make(map[int16][]int),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16SliceintMapAllocSpace(size int) *Int16SliceintMap {
	om := &Int16SliceintMap{
		store:   make(map[int16][]int, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16SliceintMap) Set(key int16, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16SliceintMap) Get(key int16) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16SliceintMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16SliceintMap) Len() int {
	return len(om.store)
}

func (om *Int16SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceint8Map struct {
	store   map[int16][]int8
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceint8Map() *Int16Sliceint8Map {
	om := &Int16Sliceint8Map{
		store:   make(map[int16][]int8),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceint8MapAllocSpace(size int) *Int16Sliceint8Map {
	om := &Int16Sliceint8Map{
		store:   make(map[int16][]int8, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceint8Map) Set(key int16, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceint8Map) Get(key int16) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceint8Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceint16Map struct {
	store   map[int16][]int16
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceint16Map() *Int16Sliceint16Map {
	om := &Int16Sliceint16Map{
		store:   make(map[int16][]int16),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceint16MapAllocSpace(size int) *Int16Sliceint16Map {
	om := &Int16Sliceint16Map{
		store:   make(map[int16][]int16, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceint16Map) Set(key int16, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceint16Map) Get(key int16) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceint16Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceint32Map struct {
	store   map[int16][]int32
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceint32Map() *Int16Sliceint32Map {
	om := &Int16Sliceint32Map{
		store:   make(map[int16][]int32),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceint32MapAllocSpace(size int) *Int16Sliceint32Map {
	om := &Int16Sliceint32Map{
		store:   make(map[int16][]int32, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceint32Map) Set(key int16, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceint32Map) Get(key int16) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceint32Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceint64Map struct {
	store   map[int16][]int64
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceint64Map() *Int16Sliceint64Map {
	om := &Int16Sliceint64Map{
		store:   make(map[int16][]int64),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceint64MapAllocSpace(size int) *Int16Sliceint64Map {
	om := &Int16Sliceint64Map{
		store:   make(map[int16][]int64, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceint64Map) Set(key int16, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceint64Map) Get(key int16) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceint64Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceuint8Map struct {
	store   map[int16][]uint8
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceuint8Map() *Int16Sliceuint8Map {
	om := &Int16Sliceuint8Map{
		store:   make(map[int16][]uint8),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceuint8MapAllocSpace(size int) *Int16Sliceuint8Map {
	om := &Int16Sliceuint8Map{
		store:   make(map[int16][]uint8, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceuint8Map) Set(key int16, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceuint8Map) Get(key int16) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceuint8Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceuint16Map struct {
	store   map[int16][]uint16
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceuint16Map() *Int16Sliceuint16Map {
	om := &Int16Sliceuint16Map{
		store:   make(map[int16][]uint16),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceuint16MapAllocSpace(size int) *Int16Sliceuint16Map {
	om := &Int16Sliceuint16Map{
		store:   make(map[int16][]uint16, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceuint16Map) Set(key int16, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceuint16Map) Get(key int16) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceuint16Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceuint32Map struct {
	store   map[int16][]uint32
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceuint32Map() *Int16Sliceuint32Map {
	om := &Int16Sliceuint32Map{
		store:   make(map[int16][]uint32),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceuint32MapAllocSpace(size int) *Int16Sliceuint32Map {
	om := &Int16Sliceuint32Map{
		store:   make(map[int16][]uint32, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceuint32Map) Set(key int16, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceuint32Map) Get(key int16) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceuint32Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16Sliceuint64Map struct {
	store   map[int16][]uint64
	mapper  []int16
	isWrite bool
}

func NewInt16Sliceuint64Map() *Int16Sliceuint64Map {
	om := &Int16Sliceuint64Map{
		store:   make(map[int16][]uint64),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16Sliceuint64MapAllocSpace(size int) *Int16Sliceuint64Map {
	om := &Int16Sliceuint64Map{
		store:   make(map[int16][]uint64, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16Sliceuint64Map) Set(key int16, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16Sliceuint64Map) Get(key int16) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16Sliceuint64Map) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Int16Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int16SlicestringMap struct {
	store   map[int16][]string
	mapper  []int16
	isWrite bool
}

func NewInt16SlicestringMap() *Int16SlicestringMap {
	om := &Int16SlicestringMap{
		store:   make(map[int16][]string),
		mapper:  make([]int16, 0),
		isWrite: false,
	}
	return om
}

func NewInt16SlicestringMapAllocSpace(size int) *Int16SlicestringMap {
	om := &Int16SlicestringMap{
		store:   make(map[int16][]string, size),
		mapper:  make([]int16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int16SlicestringMap) Set(key int16, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int16SlicestringMap) Get(key int16) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int16SlicestringMap) Delete(key int16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int16SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Int16SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int16SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]int16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int16SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]int16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32IntMap struct {
	store   map[int32]int
	mapper  []int32
	isWrite bool
}

func NewInt32IntMap() *Int32IntMap {
	om := &Int32IntMap{
		store:   make(map[int32]int),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32IntMapAllocSpace(size int) *Int32IntMap {
	om := &Int32IntMap{
		store:   make(map[int32]int, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32IntMap) Set(key int32, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32IntMap) Get(key int32) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32IntMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32IntMap) Len() int {
	return len(om.store)
}

func (om *Int32IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Int8Map struct {
	store   map[int32]int8
	mapper  []int32
	isWrite bool
}

func NewInt32Int8Map() *Int32Int8Map {
	om := &Int32Int8Map{
		store:   make(map[int32]int8),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Int8MapAllocSpace(size int) *Int32Int8Map {
	om := &Int32Int8Map{
		store:   make(map[int32]int8, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Int8Map) Set(key int32, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Int8Map) Get(key int32) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Int8Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Int8Map) Len() int {
	return len(om.store)
}

func (om *Int32Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Int16Map struct {
	store   map[int32]int16
	mapper  []int32
	isWrite bool
}

func NewInt32Int16Map() *Int32Int16Map {
	om := &Int32Int16Map{
		store:   make(map[int32]int16),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Int16MapAllocSpace(size int) *Int32Int16Map {
	om := &Int32Int16Map{
		store:   make(map[int32]int16, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Int16Map) Set(key int32, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Int16Map) Get(key int32) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Int16Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Int16Map) Len() int {
	return len(om.store)
}

func (om *Int32Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Int32Map struct {
	store   map[int32]int32
	mapper  []int32
	isWrite bool
}

func NewInt32Int32Map() *Int32Int32Map {
	om := &Int32Int32Map{
		store:   make(map[int32]int32),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Int32MapAllocSpace(size int) *Int32Int32Map {
	om := &Int32Int32Map{
		store:   make(map[int32]int32, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Int32Map) Set(key int32, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Int32Map) Get(key int32) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Int32Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Int32Map) Len() int {
	return len(om.store)
}

func (om *Int32Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Int64Map struct {
	store   map[int32]int64
	mapper  []int32
	isWrite bool
}

func NewInt32Int64Map() *Int32Int64Map {
	om := &Int32Int64Map{
		store:   make(map[int32]int64),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Int64MapAllocSpace(size int) *Int32Int64Map {
	om := &Int32Int64Map{
		store:   make(map[int32]int64, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Int64Map) Set(key int32, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Int64Map) Get(key int32) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Int64Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Int64Map) Len() int {
	return len(om.store)
}

func (om *Int32Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Uint8Map struct {
	store   map[int32]uint8
	mapper  []int32
	isWrite bool
}

func NewInt32Uint8Map() *Int32Uint8Map {
	om := &Int32Uint8Map{
		store:   make(map[int32]uint8),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Uint8MapAllocSpace(size int) *Int32Uint8Map {
	om := &Int32Uint8Map{
		store:   make(map[int32]uint8, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Uint8Map) Set(key int32, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Uint8Map) Get(key int32) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Uint8Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Uint8Map) Len() int {
	return len(om.store)
}

func (om *Int32Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Uint16Map struct {
	store   map[int32]uint16
	mapper  []int32
	isWrite bool
}

func NewInt32Uint16Map() *Int32Uint16Map {
	om := &Int32Uint16Map{
		store:   make(map[int32]uint16),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Uint16MapAllocSpace(size int) *Int32Uint16Map {
	om := &Int32Uint16Map{
		store:   make(map[int32]uint16, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Uint16Map) Set(key int32, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Uint16Map) Get(key int32) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Uint16Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Uint16Map) Len() int {
	return len(om.store)
}

func (om *Int32Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Uint32Map struct {
	store   map[int32]uint32
	mapper  []int32
	isWrite bool
}

func NewInt32Uint32Map() *Int32Uint32Map {
	om := &Int32Uint32Map{
		store:   make(map[int32]uint32),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Uint32MapAllocSpace(size int) *Int32Uint32Map {
	om := &Int32Uint32Map{
		store:   make(map[int32]uint32, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Uint32Map) Set(key int32, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Uint32Map) Get(key int32) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Uint32Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Uint32Map) Len() int {
	return len(om.store)
}

func (om *Int32Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Uint64Map struct {
	store   map[int32]uint64
	mapper  []int32
	isWrite bool
}

func NewInt32Uint64Map() *Int32Uint64Map {
	om := &Int32Uint64Map{
		store:   make(map[int32]uint64),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Uint64MapAllocSpace(size int) *Int32Uint64Map {
	om := &Int32Uint64Map{
		store:   make(map[int32]uint64, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Uint64Map) Set(key int32, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Uint64Map) Get(key int32) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Uint64Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Uint64Map) Len() int {
	return len(om.store)
}

func (om *Int32Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32StringMap struct {
	store   map[int32]string
	mapper  []int32
	isWrite bool
}

func NewInt32StringMap() *Int32StringMap {
	om := &Int32StringMap{
		store:   make(map[int32]string),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32StringMapAllocSpace(size int) *Int32StringMap {
	om := &Int32StringMap{
		store:   make(map[int32]string, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32StringMap) Set(key int32, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32StringMap) Get(key int32) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32StringMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32StringMap) Len() int {
	return len(om.store)
}

func (om *Int32StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Float32Map struct {
	store   map[int32]float32
	mapper  []int32
	isWrite bool
}

func NewInt32Float32Map() *Int32Float32Map {
	om := &Int32Float32Map{
		store:   make(map[int32]float32),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Float32MapAllocSpace(size int) *Int32Float32Map {
	om := &Int32Float32Map{
		store:   make(map[int32]float32, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Float32Map) Set(key int32, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Float32Map) Get(key int32) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Float32Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Float32Map) Len() int {
	return len(om.store)
}

func (om *Int32Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Float64Map struct {
	store   map[int32]float64
	mapper  []int32
	isWrite bool
}

func NewInt32Float64Map() *Int32Float64Map {
	om := &Int32Float64Map{
		store:   make(map[int32]float64),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Float64MapAllocSpace(size int) *Int32Float64Map {
	om := &Int32Float64Map{
		store:   make(map[int32]float64, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Float64Map) Set(key int32, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Float64Map) Get(key int32) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Float64Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Float64Map) Len() int {
	return len(om.store)
}

func (om *Int32Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32InterfaceMap struct {
	store   map[int32]interface{}
	mapper  []int32
	isWrite bool
}

func NewInt32InterfaceMap() *Int32InterfaceMap {
	om := &Int32InterfaceMap{
		store:   make(map[int32]interface{}),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32InterfaceMapAllocSpace(size int) *Int32InterfaceMap {
	om := &Int32InterfaceMap{
		store:   make(map[int32]interface{}, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32InterfaceMap) Set(key int32, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32InterfaceMap) Get(key int32) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32InterfaceMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Int32InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32BoolMap struct {
	store   map[int32]bool
	mapper  []int32
	isWrite bool
}

func NewInt32BoolMap() *Int32BoolMap {
	om := &Int32BoolMap{
		store:   make(map[int32]bool),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32BoolMapAllocSpace(size int) *Int32BoolMap {
	om := &Int32BoolMap{
		store:   make(map[int32]bool, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32BoolMap) Set(key int32, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32BoolMap) Get(key int32) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32BoolMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32BoolMap) Len() int {
	return len(om.store)
}

func (om *Int32BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32StructMap struct {
	store   map[int32]struct{}
	mapper  []int32
	isWrite bool
}

func NewInt32StructMap() *Int32StructMap {
	om := &Int32StructMap{
		store:   make(map[int32]struct{}),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32StructMapAllocSpace(size int) *Int32StructMap {
	om := &Int32StructMap{
		store:   make(map[int32]struct{}, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32StructMap) Set(key int32, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32StructMap) Get(key int32) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32StructMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32StructMap) Len() int {
	return len(om.store)
}

func (om *Int32StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32SliceintMap struct {
	store   map[int32][]int
	mapper  []int32
	isWrite bool
}

func NewInt32SliceintMap() *Int32SliceintMap {
	om := &Int32SliceintMap{
		store:   make(map[int32][]int),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32SliceintMapAllocSpace(size int) *Int32SliceintMap {
	om := &Int32SliceintMap{
		store:   make(map[int32][]int, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32SliceintMap) Set(key int32, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32SliceintMap) Get(key int32) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32SliceintMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32SliceintMap) Len() int {
	return len(om.store)
}

func (om *Int32SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceint8Map struct {
	store   map[int32][]int8
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceint8Map() *Int32Sliceint8Map {
	om := &Int32Sliceint8Map{
		store:   make(map[int32][]int8),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceint8MapAllocSpace(size int) *Int32Sliceint8Map {
	om := &Int32Sliceint8Map{
		store:   make(map[int32][]int8, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceint8Map) Set(key int32, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceint8Map) Get(key int32) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceint8Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceint16Map struct {
	store   map[int32][]int16
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceint16Map() *Int32Sliceint16Map {
	om := &Int32Sliceint16Map{
		store:   make(map[int32][]int16),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceint16MapAllocSpace(size int) *Int32Sliceint16Map {
	om := &Int32Sliceint16Map{
		store:   make(map[int32][]int16, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceint16Map) Set(key int32, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceint16Map) Get(key int32) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceint16Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceint32Map struct {
	store   map[int32][]int32
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceint32Map() *Int32Sliceint32Map {
	om := &Int32Sliceint32Map{
		store:   make(map[int32][]int32),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceint32MapAllocSpace(size int) *Int32Sliceint32Map {
	om := &Int32Sliceint32Map{
		store:   make(map[int32][]int32, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceint32Map) Set(key int32, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceint32Map) Get(key int32) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceint32Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceint64Map struct {
	store   map[int32][]int64
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceint64Map() *Int32Sliceint64Map {
	om := &Int32Sliceint64Map{
		store:   make(map[int32][]int64),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceint64MapAllocSpace(size int) *Int32Sliceint64Map {
	om := &Int32Sliceint64Map{
		store:   make(map[int32][]int64, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceint64Map) Set(key int32, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceint64Map) Get(key int32) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceint64Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceuint8Map struct {
	store   map[int32][]uint8
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceuint8Map() *Int32Sliceuint8Map {
	om := &Int32Sliceuint8Map{
		store:   make(map[int32][]uint8),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceuint8MapAllocSpace(size int) *Int32Sliceuint8Map {
	om := &Int32Sliceuint8Map{
		store:   make(map[int32][]uint8, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceuint8Map) Set(key int32, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceuint8Map) Get(key int32) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceuint8Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceuint16Map struct {
	store   map[int32][]uint16
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceuint16Map() *Int32Sliceuint16Map {
	om := &Int32Sliceuint16Map{
		store:   make(map[int32][]uint16),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceuint16MapAllocSpace(size int) *Int32Sliceuint16Map {
	om := &Int32Sliceuint16Map{
		store:   make(map[int32][]uint16, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceuint16Map) Set(key int32, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceuint16Map) Get(key int32) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceuint16Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceuint32Map struct {
	store   map[int32][]uint32
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceuint32Map() *Int32Sliceuint32Map {
	om := &Int32Sliceuint32Map{
		store:   make(map[int32][]uint32),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceuint32MapAllocSpace(size int) *Int32Sliceuint32Map {
	om := &Int32Sliceuint32Map{
		store:   make(map[int32][]uint32, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceuint32Map) Set(key int32, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceuint32Map) Get(key int32) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceuint32Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32Sliceuint64Map struct {
	store   map[int32][]uint64
	mapper  []int32
	isWrite bool
}

func NewInt32Sliceuint64Map() *Int32Sliceuint64Map {
	om := &Int32Sliceuint64Map{
		store:   make(map[int32][]uint64),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32Sliceuint64MapAllocSpace(size int) *Int32Sliceuint64Map {
	om := &Int32Sliceuint64Map{
		store:   make(map[int32][]uint64, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32Sliceuint64Map) Set(key int32, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32Sliceuint64Map) Get(key int32) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32Sliceuint64Map) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Int32Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int32SlicestringMap struct {
	store   map[int32][]string
	mapper  []int32
	isWrite bool
}

func NewInt32SlicestringMap() *Int32SlicestringMap {
	om := &Int32SlicestringMap{
		store:   make(map[int32][]string),
		mapper:  make([]int32, 0),
		isWrite: false,
	}
	return om
}

func NewInt32SlicestringMapAllocSpace(size int) *Int32SlicestringMap {
	om := &Int32SlicestringMap{
		store:   make(map[int32][]string, size),
		mapper:  make([]int32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int32SlicestringMap) Set(key int32, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int32SlicestringMap) Get(key int32) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int32SlicestringMap) Delete(key int32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int32SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Int32SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int32SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]int32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int32SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]int32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64IntMap struct {
	store   map[int64]int
	mapper  []int64
	isWrite bool
}

func NewInt64IntMap() *Int64IntMap {
	om := &Int64IntMap{
		store:   make(map[int64]int),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64IntMapAllocSpace(size int) *Int64IntMap {
	om := &Int64IntMap{
		store:   make(map[int64]int, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64IntMap) Set(key int64, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64IntMap) Get(key int64) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64IntMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64IntMap) Len() int {
	return len(om.store)
}

func (om *Int64IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Int8Map struct {
	store   map[int64]int8
	mapper  []int64
	isWrite bool
}

func NewInt64Int8Map() *Int64Int8Map {
	om := &Int64Int8Map{
		store:   make(map[int64]int8),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Int8MapAllocSpace(size int) *Int64Int8Map {
	om := &Int64Int8Map{
		store:   make(map[int64]int8, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Int8Map) Set(key int64, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Int8Map) Get(key int64) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Int8Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Int8Map) Len() int {
	return len(om.store)
}

func (om *Int64Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Int16Map struct {
	store   map[int64]int16
	mapper  []int64
	isWrite bool
}

func NewInt64Int16Map() *Int64Int16Map {
	om := &Int64Int16Map{
		store:   make(map[int64]int16),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Int16MapAllocSpace(size int) *Int64Int16Map {
	om := &Int64Int16Map{
		store:   make(map[int64]int16, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Int16Map) Set(key int64, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Int16Map) Get(key int64) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Int16Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Int16Map) Len() int {
	return len(om.store)
}

func (om *Int64Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Int32Map struct {
	store   map[int64]int32
	mapper  []int64
	isWrite bool
}

func NewInt64Int32Map() *Int64Int32Map {
	om := &Int64Int32Map{
		store:   make(map[int64]int32),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Int32MapAllocSpace(size int) *Int64Int32Map {
	om := &Int64Int32Map{
		store:   make(map[int64]int32, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Int32Map) Set(key int64, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Int32Map) Get(key int64) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Int32Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Int32Map) Len() int {
	return len(om.store)
}

func (om *Int64Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Int64Map struct {
	store   map[int64]int64
	mapper  []int64
	isWrite bool
}

func NewInt64Int64Map() *Int64Int64Map {
	om := &Int64Int64Map{
		store:   make(map[int64]int64),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Int64MapAllocSpace(size int) *Int64Int64Map {
	om := &Int64Int64Map{
		store:   make(map[int64]int64, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Int64Map) Set(key int64, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Int64Map) Get(key int64) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Int64Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Int64Map) Len() int {
	return len(om.store)
}

func (om *Int64Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Uint8Map struct {
	store   map[int64]uint8
	mapper  []int64
	isWrite bool
}

func NewInt64Uint8Map() *Int64Uint8Map {
	om := &Int64Uint8Map{
		store:   make(map[int64]uint8),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Uint8MapAllocSpace(size int) *Int64Uint8Map {
	om := &Int64Uint8Map{
		store:   make(map[int64]uint8, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Uint8Map) Set(key int64, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Uint8Map) Get(key int64) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Uint8Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Uint8Map) Len() int {
	return len(om.store)
}

func (om *Int64Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Uint16Map struct {
	store   map[int64]uint16
	mapper  []int64
	isWrite bool
}

func NewInt64Uint16Map() *Int64Uint16Map {
	om := &Int64Uint16Map{
		store:   make(map[int64]uint16),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Uint16MapAllocSpace(size int) *Int64Uint16Map {
	om := &Int64Uint16Map{
		store:   make(map[int64]uint16, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Uint16Map) Set(key int64, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Uint16Map) Get(key int64) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Uint16Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Uint16Map) Len() int {
	return len(om.store)
}

func (om *Int64Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Uint32Map struct {
	store   map[int64]uint32
	mapper  []int64
	isWrite bool
}

func NewInt64Uint32Map() *Int64Uint32Map {
	om := &Int64Uint32Map{
		store:   make(map[int64]uint32),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Uint32MapAllocSpace(size int) *Int64Uint32Map {
	om := &Int64Uint32Map{
		store:   make(map[int64]uint32, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Uint32Map) Set(key int64, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Uint32Map) Get(key int64) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Uint32Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Uint32Map) Len() int {
	return len(om.store)
}

func (om *Int64Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Uint64Map struct {
	store   map[int64]uint64
	mapper  []int64
	isWrite bool
}

func NewInt64Uint64Map() *Int64Uint64Map {
	om := &Int64Uint64Map{
		store:   make(map[int64]uint64),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Uint64MapAllocSpace(size int) *Int64Uint64Map {
	om := &Int64Uint64Map{
		store:   make(map[int64]uint64, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Uint64Map) Set(key int64, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Uint64Map) Get(key int64) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Uint64Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Uint64Map) Len() int {
	return len(om.store)
}

func (om *Int64Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64StringMap struct {
	store   map[int64]string
	mapper  []int64
	isWrite bool
}

func NewInt64StringMap() *Int64StringMap {
	om := &Int64StringMap{
		store:   make(map[int64]string),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64StringMapAllocSpace(size int) *Int64StringMap {
	om := &Int64StringMap{
		store:   make(map[int64]string, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64StringMap) Set(key int64, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64StringMap) Get(key int64) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64StringMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64StringMap) Len() int {
	return len(om.store)
}

func (om *Int64StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Float32Map struct {
	store   map[int64]float32
	mapper  []int64
	isWrite bool
}

func NewInt64Float32Map() *Int64Float32Map {
	om := &Int64Float32Map{
		store:   make(map[int64]float32),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Float32MapAllocSpace(size int) *Int64Float32Map {
	om := &Int64Float32Map{
		store:   make(map[int64]float32, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Float32Map) Set(key int64, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Float32Map) Get(key int64) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Float32Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Float32Map) Len() int {
	return len(om.store)
}

func (om *Int64Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Float64Map struct {
	store   map[int64]float64
	mapper  []int64
	isWrite bool
}

func NewInt64Float64Map() *Int64Float64Map {
	om := &Int64Float64Map{
		store:   make(map[int64]float64),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Float64MapAllocSpace(size int) *Int64Float64Map {
	om := &Int64Float64Map{
		store:   make(map[int64]float64, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Float64Map) Set(key int64, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Float64Map) Get(key int64) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Float64Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Float64Map) Len() int {
	return len(om.store)
}

func (om *Int64Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64InterfaceMap struct {
	store   map[int64]interface{}
	mapper  []int64
	isWrite bool
}

func NewInt64InterfaceMap() *Int64InterfaceMap {
	om := &Int64InterfaceMap{
		store:   make(map[int64]interface{}),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64InterfaceMapAllocSpace(size int) *Int64InterfaceMap {
	om := &Int64InterfaceMap{
		store:   make(map[int64]interface{}, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64InterfaceMap) Set(key int64, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64InterfaceMap) Get(key int64) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64InterfaceMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Int64InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64BoolMap struct {
	store   map[int64]bool
	mapper  []int64
	isWrite bool
}

func NewInt64BoolMap() *Int64BoolMap {
	om := &Int64BoolMap{
		store:   make(map[int64]bool),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64BoolMapAllocSpace(size int) *Int64BoolMap {
	om := &Int64BoolMap{
		store:   make(map[int64]bool, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64BoolMap) Set(key int64, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64BoolMap) Get(key int64) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64BoolMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64BoolMap) Len() int {
	return len(om.store)
}

func (om *Int64BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64StructMap struct {
	store   map[int64]struct{}
	mapper  []int64
	isWrite bool
}

func NewInt64StructMap() *Int64StructMap {
	om := &Int64StructMap{
		store:   make(map[int64]struct{}),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64StructMapAllocSpace(size int) *Int64StructMap {
	om := &Int64StructMap{
		store:   make(map[int64]struct{}, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64StructMap) Set(key int64, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64StructMap) Get(key int64) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64StructMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64StructMap) Len() int {
	return len(om.store)
}

func (om *Int64StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64SliceintMap struct {
	store   map[int64][]int
	mapper  []int64
	isWrite bool
}

func NewInt64SliceintMap() *Int64SliceintMap {
	om := &Int64SliceintMap{
		store:   make(map[int64][]int),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64SliceintMapAllocSpace(size int) *Int64SliceintMap {
	om := &Int64SliceintMap{
		store:   make(map[int64][]int, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64SliceintMap) Set(key int64, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64SliceintMap) Get(key int64) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64SliceintMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64SliceintMap) Len() int {
	return len(om.store)
}

func (om *Int64SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceint8Map struct {
	store   map[int64][]int8
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceint8Map() *Int64Sliceint8Map {
	om := &Int64Sliceint8Map{
		store:   make(map[int64][]int8),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceint8MapAllocSpace(size int) *Int64Sliceint8Map {
	om := &Int64Sliceint8Map{
		store:   make(map[int64][]int8, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceint8Map) Set(key int64, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceint8Map) Get(key int64) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceint8Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceint16Map struct {
	store   map[int64][]int16
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceint16Map() *Int64Sliceint16Map {
	om := &Int64Sliceint16Map{
		store:   make(map[int64][]int16),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceint16MapAllocSpace(size int) *Int64Sliceint16Map {
	om := &Int64Sliceint16Map{
		store:   make(map[int64][]int16, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceint16Map) Set(key int64, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceint16Map) Get(key int64) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceint16Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceint32Map struct {
	store   map[int64][]int32
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceint32Map() *Int64Sliceint32Map {
	om := &Int64Sliceint32Map{
		store:   make(map[int64][]int32),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceint32MapAllocSpace(size int) *Int64Sliceint32Map {
	om := &Int64Sliceint32Map{
		store:   make(map[int64][]int32, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceint32Map) Set(key int64, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceint32Map) Get(key int64) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceint32Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceint64Map struct {
	store   map[int64][]int64
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceint64Map() *Int64Sliceint64Map {
	om := &Int64Sliceint64Map{
		store:   make(map[int64][]int64),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceint64MapAllocSpace(size int) *Int64Sliceint64Map {
	om := &Int64Sliceint64Map{
		store:   make(map[int64][]int64, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceint64Map) Set(key int64, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceint64Map) Get(key int64) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceint64Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceuint8Map struct {
	store   map[int64][]uint8
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceuint8Map() *Int64Sliceuint8Map {
	om := &Int64Sliceuint8Map{
		store:   make(map[int64][]uint8),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceuint8MapAllocSpace(size int) *Int64Sliceuint8Map {
	om := &Int64Sliceuint8Map{
		store:   make(map[int64][]uint8, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceuint8Map) Set(key int64, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceuint8Map) Get(key int64) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceuint8Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceuint16Map struct {
	store   map[int64][]uint16
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceuint16Map() *Int64Sliceuint16Map {
	om := &Int64Sliceuint16Map{
		store:   make(map[int64][]uint16),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceuint16MapAllocSpace(size int) *Int64Sliceuint16Map {
	om := &Int64Sliceuint16Map{
		store:   make(map[int64][]uint16, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceuint16Map) Set(key int64, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceuint16Map) Get(key int64) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceuint16Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceuint32Map struct {
	store   map[int64][]uint32
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceuint32Map() *Int64Sliceuint32Map {
	om := &Int64Sliceuint32Map{
		store:   make(map[int64][]uint32),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceuint32MapAllocSpace(size int) *Int64Sliceuint32Map {
	om := &Int64Sliceuint32Map{
		store:   make(map[int64][]uint32, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceuint32Map) Set(key int64, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceuint32Map) Get(key int64) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceuint32Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64Sliceuint64Map struct {
	store   map[int64][]uint64
	mapper  []int64
	isWrite bool
}

func NewInt64Sliceuint64Map() *Int64Sliceuint64Map {
	om := &Int64Sliceuint64Map{
		store:   make(map[int64][]uint64),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64Sliceuint64MapAllocSpace(size int) *Int64Sliceuint64Map {
	om := &Int64Sliceuint64Map{
		store:   make(map[int64][]uint64, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64Sliceuint64Map) Set(key int64, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64Sliceuint64Map) Get(key int64) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64Sliceuint64Map) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Int64Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Int64SlicestringMap struct {
	store   map[int64][]string
	mapper  []int64
	isWrite bool
}

func NewInt64SlicestringMap() *Int64SlicestringMap {
	om := &Int64SlicestringMap{
		store:   make(map[int64][]string),
		mapper:  make([]int64, 0),
		isWrite: false,
	}
	return om
}

func NewInt64SlicestringMapAllocSpace(size int) *Int64SlicestringMap {
	om := &Int64SlicestringMap{
		store:   make(map[int64][]string, size),
		mapper:  make([]int64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Int64SlicestringMap) Set(key int64, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Int64SlicestringMap) Get(key int64) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Int64SlicestringMap) Delete(key int64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Int64SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Int64SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Int64SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]int64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Int64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Int64SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]int64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Int64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8IntMap struct {
	store   map[uint8]int
	mapper  []uint8
	isWrite bool
}

func NewUint8IntMap() *Uint8IntMap {
	om := &Uint8IntMap{
		store:   make(map[uint8]int),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8IntMapAllocSpace(size int) *Uint8IntMap {
	om := &Uint8IntMap{
		store:   make(map[uint8]int, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8IntMap) Set(key uint8, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8IntMap) Get(key uint8) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8IntMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8IntMap) Len() int {
	return len(om.store)
}

func (om *Uint8IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Int8Map struct {
	store   map[uint8]int8
	mapper  []uint8
	isWrite bool
}

func NewUint8Int8Map() *Uint8Int8Map {
	om := &Uint8Int8Map{
		store:   make(map[uint8]int8),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Int8MapAllocSpace(size int) *Uint8Int8Map {
	om := &Uint8Int8Map{
		store:   make(map[uint8]int8, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Int8Map) Set(key uint8, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Int8Map) Get(key uint8) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Int8Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Int8Map) Len() int {
	return len(om.store)
}

func (om *Uint8Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Int16Map struct {
	store   map[uint8]int16
	mapper  []uint8
	isWrite bool
}

func NewUint8Int16Map() *Uint8Int16Map {
	om := &Uint8Int16Map{
		store:   make(map[uint8]int16),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Int16MapAllocSpace(size int) *Uint8Int16Map {
	om := &Uint8Int16Map{
		store:   make(map[uint8]int16, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Int16Map) Set(key uint8, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Int16Map) Get(key uint8) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Int16Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Int16Map) Len() int {
	return len(om.store)
}

func (om *Uint8Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Int32Map struct {
	store   map[uint8]int32
	mapper  []uint8
	isWrite bool
}

func NewUint8Int32Map() *Uint8Int32Map {
	om := &Uint8Int32Map{
		store:   make(map[uint8]int32),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Int32MapAllocSpace(size int) *Uint8Int32Map {
	om := &Uint8Int32Map{
		store:   make(map[uint8]int32, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Int32Map) Set(key uint8, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Int32Map) Get(key uint8) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Int32Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Int32Map) Len() int {
	return len(om.store)
}

func (om *Uint8Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Int64Map struct {
	store   map[uint8]int64
	mapper  []uint8
	isWrite bool
}

func NewUint8Int64Map() *Uint8Int64Map {
	om := &Uint8Int64Map{
		store:   make(map[uint8]int64),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Int64MapAllocSpace(size int) *Uint8Int64Map {
	om := &Uint8Int64Map{
		store:   make(map[uint8]int64, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Int64Map) Set(key uint8, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Int64Map) Get(key uint8) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Int64Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Int64Map) Len() int {
	return len(om.store)
}

func (om *Uint8Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Uint8Map struct {
	store   map[uint8]uint8
	mapper  []uint8
	isWrite bool
}

func NewUint8Uint8Map() *Uint8Uint8Map {
	om := &Uint8Uint8Map{
		store:   make(map[uint8]uint8),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Uint8MapAllocSpace(size int) *Uint8Uint8Map {
	om := &Uint8Uint8Map{
		store:   make(map[uint8]uint8, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Uint8Map) Set(key uint8, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Uint8Map) Get(key uint8) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Uint8Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Uint8Map) Len() int {
	return len(om.store)
}

func (om *Uint8Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Uint16Map struct {
	store   map[uint8]uint16
	mapper  []uint8
	isWrite bool
}

func NewUint8Uint16Map() *Uint8Uint16Map {
	om := &Uint8Uint16Map{
		store:   make(map[uint8]uint16),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Uint16MapAllocSpace(size int) *Uint8Uint16Map {
	om := &Uint8Uint16Map{
		store:   make(map[uint8]uint16, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Uint16Map) Set(key uint8, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Uint16Map) Get(key uint8) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Uint16Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Uint16Map) Len() int {
	return len(om.store)
}

func (om *Uint8Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Uint32Map struct {
	store   map[uint8]uint32
	mapper  []uint8
	isWrite bool
}

func NewUint8Uint32Map() *Uint8Uint32Map {
	om := &Uint8Uint32Map{
		store:   make(map[uint8]uint32),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Uint32MapAllocSpace(size int) *Uint8Uint32Map {
	om := &Uint8Uint32Map{
		store:   make(map[uint8]uint32, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Uint32Map) Set(key uint8, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Uint32Map) Get(key uint8) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Uint32Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Uint32Map) Len() int {
	return len(om.store)
}

func (om *Uint8Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Uint64Map struct {
	store   map[uint8]uint64
	mapper  []uint8
	isWrite bool
}

func NewUint8Uint64Map() *Uint8Uint64Map {
	om := &Uint8Uint64Map{
		store:   make(map[uint8]uint64),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Uint64MapAllocSpace(size int) *Uint8Uint64Map {
	om := &Uint8Uint64Map{
		store:   make(map[uint8]uint64, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Uint64Map) Set(key uint8, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Uint64Map) Get(key uint8) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Uint64Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Uint64Map) Len() int {
	return len(om.store)
}

func (om *Uint8Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8StringMap struct {
	store   map[uint8]string
	mapper  []uint8
	isWrite bool
}

func NewUint8StringMap() *Uint8StringMap {
	om := &Uint8StringMap{
		store:   make(map[uint8]string),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8StringMapAllocSpace(size int) *Uint8StringMap {
	om := &Uint8StringMap{
		store:   make(map[uint8]string, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8StringMap) Set(key uint8, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8StringMap) Get(key uint8) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8StringMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8StringMap) Len() int {
	return len(om.store)
}

func (om *Uint8StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Float32Map struct {
	store   map[uint8]float32
	mapper  []uint8
	isWrite bool
}

func NewUint8Float32Map() *Uint8Float32Map {
	om := &Uint8Float32Map{
		store:   make(map[uint8]float32),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Float32MapAllocSpace(size int) *Uint8Float32Map {
	om := &Uint8Float32Map{
		store:   make(map[uint8]float32, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Float32Map) Set(key uint8, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Float32Map) Get(key uint8) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Float32Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Float32Map) Len() int {
	return len(om.store)
}

func (om *Uint8Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Float64Map struct {
	store   map[uint8]float64
	mapper  []uint8
	isWrite bool
}

func NewUint8Float64Map() *Uint8Float64Map {
	om := &Uint8Float64Map{
		store:   make(map[uint8]float64),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Float64MapAllocSpace(size int) *Uint8Float64Map {
	om := &Uint8Float64Map{
		store:   make(map[uint8]float64, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Float64Map) Set(key uint8, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Float64Map) Get(key uint8) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Float64Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Float64Map) Len() int {
	return len(om.store)
}

func (om *Uint8Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8InterfaceMap struct {
	store   map[uint8]interface{}
	mapper  []uint8
	isWrite bool
}

func NewUint8InterfaceMap() *Uint8InterfaceMap {
	om := &Uint8InterfaceMap{
		store:   make(map[uint8]interface{}),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8InterfaceMapAllocSpace(size int) *Uint8InterfaceMap {
	om := &Uint8InterfaceMap{
		store:   make(map[uint8]interface{}, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8InterfaceMap) Set(key uint8, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8InterfaceMap) Get(key uint8) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8InterfaceMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Uint8InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8BoolMap struct {
	store   map[uint8]bool
	mapper  []uint8
	isWrite bool
}

func NewUint8BoolMap() *Uint8BoolMap {
	om := &Uint8BoolMap{
		store:   make(map[uint8]bool),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8BoolMapAllocSpace(size int) *Uint8BoolMap {
	om := &Uint8BoolMap{
		store:   make(map[uint8]bool, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8BoolMap) Set(key uint8, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8BoolMap) Get(key uint8) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8BoolMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8BoolMap) Len() int {
	return len(om.store)
}

func (om *Uint8BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8StructMap struct {
	store   map[uint8]struct{}
	mapper  []uint8
	isWrite bool
}

func NewUint8StructMap() *Uint8StructMap {
	om := &Uint8StructMap{
		store:   make(map[uint8]struct{}),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8StructMapAllocSpace(size int) *Uint8StructMap {
	om := &Uint8StructMap{
		store:   make(map[uint8]struct{}, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8StructMap) Set(key uint8, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8StructMap) Get(key uint8) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8StructMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8StructMap) Len() int {
	return len(om.store)
}

func (om *Uint8StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8SliceintMap struct {
	store   map[uint8][]int
	mapper  []uint8
	isWrite bool
}

func NewUint8SliceintMap() *Uint8SliceintMap {
	om := &Uint8SliceintMap{
		store:   make(map[uint8][]int),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8SliceintMapAllocSpace(size int) *Uint8SliceintMap {
	om := &Uint8SliceintMap{
		store:   make(map[uint8][]int, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8SliceintMap) Set(key uint8, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8SliceintMap) Get(key uint8) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8SliceintMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8SliceintMap) Len() int {
	return len(om.store)
}

func (om *Uint8SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceint8Map struct {
	store   map[uint8][]int8
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceint8Map() *Uint8Sliceint8Map {
	om := &Uint8Sliceint8Map{
		store:   make(map[uint8][]int8),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceint8MapAllocSpace(size int) *Uint8Sliceint8Map {
	om := &Uint8Sliceint8Map{
		store:   make(map[uint8][]int8, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceint8Map) Set(key uint8, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceint8Map) Get(key uint8) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceint8Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceint16Map struct {
	store   map[uint8][]int16
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceint16Map() *Uint8Sliceint16Map {
	om := &Uint8Sliceint16Map{
		store:   make(map[uint8][]int16),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceint16MapAllocSpace(size int) *Uint8Sliceint16Map {
	om := &Uint8Sliceint16Map{
		store:   make(map[uint8][]int16, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceint16Map) Set(key uint8, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceint16Map) Get(key uint8) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceint16Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceint32Map struct {
	store   map[uint8][]int32
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceint32Map() *Uint8Sliceint32Map {
	om := &Uint8Sliceint32Map{
		store:   make(map[uint8][]int32),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceint32MapAllocSpace(size int) *Uint8Sliceint32Map {
	om := &Uint8Sliceint32Map{
		store:   make(map[uint8][]int32, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceint32Map) Set(key uint8, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceint32Map) Get(key uint8) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceint32Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceint64Map struct {
	store   map[uint8][]int64
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceint64Map() *Uint8Sliceint64Map {
	om := &Uint8Sliceint64Map{
		store:   make(map[uint8][]int64),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceint64MapAllocSpace(size int) *Uint8Sliceint64Map {
	om := &Uint8Sliceint64Map{
		store:   make(map[uint8][]int64, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceint64Map) Set(key uint8, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceint64Map) Get(key uint8) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceint64Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceuint8Map struct {
	store   map[uint8][]uint8
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceuint8Map() *Uint8Sliceuint8Map {
	om := &Uint8Sliceuint8Map{
		store:   make(map[uint8][]uint8),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceuint8MapAllocSpace(size int) *Uint8Sliceuint8Map {
	om := &Uint8Sliceuint8Map{
		store:   make(map[uint8][]uint8, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceuint8Map) Set(key uint8, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceuint8Map) Get(key uint8) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceuint8Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceuint16Map struct {
	store   map[uint8][]uint16
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceuint16Map() *Uint8Sliceuint16Map {
	om := &Uint8Sliceuint16Map{
		store:   make(map[uint8][]uint16),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceuint16MapAllocSpace(size int) *Uint8Sliceuint16Map {
	om := &Uint8Sliceuint16Map{
		store:   make(map[uint8][]uint16, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceuint16Map) Set(key uint8, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceuint16Map) Get(key uint8) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceuint16Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceuint32Map struct {
	store   map[uint8][]uint32
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceuint32Map() *Uint8Sliceuint32Map {
	om := &Uint8Sliceuint32Map{
		store:   make(map[uint8][]uint32),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceuint32MapAllocSpace(size int) *Uint8Sliceuint32Map {
	om := &Uint8Sliceuint32Map{
		store:   make(map[uint8][]uint32, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceuint32Map) Set(key uint8, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceuint32Map) Get(key uint8) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceuint32Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8Sliceuint64Map struct {
	store   map[uint8][]uint64
	mapper  []uint8
	isWrite bool
}

func NewUint8Sliceuint64Map() *Uint8Sliceuint64Map {
	om := &Uint8Sliceuint64Map{
		store:   make(map[uint8][]uint64),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8Sliceuint64MapAllocSpace(size int) *Uint8Sliceuint64Map {
	om := &Uint8Sliceuint64Map{
		store:   make(map[uint8][]uint64, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8Sliceuint64Map) Set(key uint8, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8Sliceuint64Map) Get(key uint8) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8Sliceuint64Map) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Uint8Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint8SlicestringMap struct {
	store   map[uint8][]string
	mapper  []uint8
	isWrite bool
}

func NewUint8SlicestringMap() *Uint8SlicestringMap {
	om := &Uint8SlicestringMap{
		store:   make(map[uint8][]string),
		mapper:  make([]uint8, 0),
		isWrite: false,
	}
	return om
}

func NewUint8SlicestringMapAllocSpace(size int) *Uint8SlicestringMap {
	om := &Uint8SlicestringMap{
		store:   make(map[uint8][]string, size),
		mapper:  make([]uint8, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint8SlicestringMap) Set(key uint8, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint8SlicestringMap) Get(key uint8) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint8SlicestringMap) Delete(key uint8) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint8SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Uint8SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint8SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]uint8, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint8s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint8SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]uint8, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint8sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16IntMap struct {
	store   map[uint16]int
	mapper  []uint16
	isWrite bool
}

func NewUint16IntMap() *Uint16IntMap {
	om := &Uint16IntMap{
		store:   make(map[uint16]int),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16IntMapAllocSpace(size int) *Uint16IntMap {
	om := &Uint16IntMap{
		store:   make(map[uint16]int, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16IntMap) Set(key uint16, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16IntMap) Get(key uint16) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16IntMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16IntMap) Len() int {
	return len(om.store)
}

func (om *Uint16IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Int8Map struct {
	store   map[uint16]int8
	mapper  []uint16
	isWrite bool
}

func NewUint16Int8Map() *Uint16Int8Map {
	om := &Uint16Int8Map{
		store:   make(map[uint16]int8),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Int8MapAllocSpace(size int) *Uint16Int8Map {
	om := &Uint16Int8Map{
		store:   make(map[uint16]int8, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Int8Map) Set(key uint16, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Int8Map) Get(key uint16) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Int8Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Int8Map) Len() int {
	return len(om.store)
}

func (om *Uint16Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Int16Map struct {
	store   map[uint16]int16
	mapper  []uint16
	isWrite bool
}

func NewUint16Int16Map() *Uint16Int16Map {
	om := &Uint16Int16Map{
		store:   make(map[uint16]int16),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Int16MapAllocSpace(size int) *Uint16Int16Map {
	om := &Uint16Int16Map{
		store:   make(map[uint16]int16, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Int16Map) Set(key uint16, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Int16Map) Get(key uint16) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Int16Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Int16Map) Len() int {
	return len(om.store)
}

func (om *Uint16Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Int32Map struct {
	store   map[uint16]int32
	mapper  []uint16
	isWrite bool
}

func NewUint16Int32Map() *Uint16Int32Map {
	om := &Uint16Int32Map{
		store:   make(map[uint16]int32),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Int32MapAllocSpace(size int) *Uint16Int32Map {
	om := &Uint16Int32Map{
		store:   make(map[uint16]int32, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Int32Map) Set(key uint16, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Int32Map) Get(key uint16) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Int32Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Int32Map) Len() int {
	return len(om.store)
}

func (om *Uint16Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Int64Map struct {
	store   map[uint16]int64
	mapper  []uint16
	isWrite bool
}

func NewUint16Int64Map() *Uint16Int64Map {
	om := &Uint16Int64Map{
		store:   make(map[uint16]int64),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Int64MapAllocSpace(size int) *Uint16Int64Map {
	om := &Uint16Int64Map{
		store:   make(map[uint16]int64, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Int64Map) Set(key uint16, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Int64Map) Get(key uint16) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Int64Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Int64Map) Len() int {
	return len(om.store)
}

func (om *Uint16Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Uint8Map struct {
	store   map[uint16]uint8
	mapper  []uint16
	isWrite bool
}

func NewUint16Uint8Map() *Uint16Uint8Map {
	om := &Uint16Uint8Map{
		store:   make(map[uint16]uint8),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Uint8MapAllocSpace(size int) *Uint16Uint8Map {
	om := &Uint16Uint8Map{
		store:   make(map[uint16]uint8, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Uint8Map) Set(key uint16, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Uint8Map) Get(key uint16) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Uint8Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Uint8Map) Len() int {
	return len(om.store)
}

func (om *Uint16Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Uint16Map struct {
	store   map[uint16]uint16
	mapper  []uint16
	isWrite bool
}

func NewUint16Uint16Map() *Uint16Uint16Map {
	om := &Uint16Uint16Map{
		store:   make(map[uint16]uint16),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Uint16MapAllocSpace(size int) *Uint16Uint16Map {
	om := &Uint16Uint16Map{
		store:   make(map[uint16]uint16, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Uint16Map) Set(key uint16, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Uint16Map) Get(key uint16) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Uint16Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Uint16Map) Len() int {
	return len(om.store)
}

func (om *Uint16Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Uint32Map struct {
	store   map[uint16]uint32
	mapper  []uint16
	isWrite bool
}

func NewUint16Uint32Map() *Uint16Uint32Map {
	om := &Uint16Uint32Map{
		store:   make(map[uint16]uint32),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Uint32MapAllocSpace(size int) *Uint16Uint32Map {
	om := &Uint16Uint32Map{
		store:   make(map[uint16]uint32, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Uint32Map) Set(key uint16, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Uint32Map) Get(key uint16) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Uint32Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Uint32Map) Len() int {
	return len(om.store)
}

func (om *Uint16Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Uint64Map struct {
	store   map[uint16]uint64
	mapper  []uint16
	isWrite bool
}

func NewUint16Uint64Map() *Uint16Uint64Map {
	om := &Uint16Uint64Map{
		store:   make(map[uint16]uint64),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Uint64MapAllocSpace(size int) *Uint16Uint64Map {
	om := &Uint16Uint64Map{
		store:   make(map[uint16]uint64, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Uint64Map) Set(key uint16, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Uint64Map) Get(key uint16) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Uint64Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Uint64Map) Len() int {
	return len(om.store)
}

func (om *Uint16Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16StringMap struct {
	store   map[uint16]string
	mapper  []uint16
	isWrite bool
}

func NewUint16StringMap() *Uint16StringMap {
	om := &Uint16StringMap{
		store:   make(map[uint16]string),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16StringMapAllocSpace(size int) *Uint16StringMap {
	om := &Uint16StringMap{
		store:   make(map[uint16]string, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16StringMap) Set(key uint16, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16StringMap) Get(key uint16) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16StringMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16StringMap) Len() int {
	return len(om.store)
}

func (om *Uint16StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Float32Map struct {
	store   map[uint16]float32
	mapper  []uint16
	isWrite bool
}

func NewUint16Float32Map() *Uint16Float32Map {
	om := &Uint16Float32Map{
		store:   make(map[uint16]float32),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Float32MapAllocSpace(size int) *Uint16Float32Map {
	om := &Uint16Float32Map{
		store:   make(map[uint16]float32, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Float32Map) Set(key uint16, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Float32Map) Get(key uint16) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Float32Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Float32Map) Len() int {
	return len(om.store)
}

func (om *Uint16Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Float64Map struct {
	store   map[uint16]float64
	mapper  []uint16
	isWrite bool
}

func NewUint16Float64Map() *Uint16Float64Map {
	om := &Uint16Float64Map{
		store:   make(map[uint16]float64),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Float64MapAllocSpace(size int) *Uint16Float64Map {
	om := &Uint16Float64Map{
		store:   make(map[uint16]float64, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Float64Map) Set(key uint16, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Float64Map) Get(key uint16) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Float64Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Float64Map) Len() int {
	return len(om.store)
}

func (om *Uint16Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16InterfaceMap struct {
	store   map[uint16]interface{}
	mapper  []uint16
	isWrite bool
}

func NewUint16InterfaceMap() *Uint16InterfaceMap {
	om := &Uint16InterfaceMap{
		store:   make(map[uint16]interface{}),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16InterfaceMapAllocSpace(size int) *Uint16InterfaceMap {
	om := &Uint16InterfaceMap{
		store:   make(map[uint16]interface{}, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16InterfaceMap) Set(key uint16, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16InterfaceMap) Get(key uint16) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16InterfaceMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Uint16InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16BoolMap struct {
	store   map[uint16]bool
	mapper  []uint16
	isWrite bool
}

func NewUint16BoolMap() *Uint16BoolMap {
	om := &Uint16BoolMap{
		store:   make(map[uint16]bool),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16BoolMapAllocSpace(size int) *Uint16BoolMap {
	om := &Uint16BoolMap{
		store:   make(map[uint16]bool, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16BoolMap) Set(key uint16, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16BoolMap) Get(key uint16) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16BoolMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16BoolMap) Len() int {
	return len(om.store)
}

func (om *Uint16BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16StructMap struct {
	store   map[uint16]struct{}
	mapper  []uint16
	isWrite bool
}

func NewUint16StructMap() *Uint16StructMap {
	om := &Uint16StructMap{
		store:   make(map[uint16]struct{}),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16StructMapAllocSpace(size int) *Uint16StructMap {
	om := &Uint16StructMap{
		store:   make(map[uint16]struct{}, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16StructMap) Set(key uint16, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16StructMap) Get(key uint16) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16StructMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16StructMap) Len() int {
	return len(om.store)
}

func (om *Uint16StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16SliceintMap struct {
	store   map[uint16][]int
	mapper  []uint16
	isWrite bool
}

func NewUint16SliceintMap() *Uint16SliceintMap {
	om := &Uint16SliceintMap{
		store:   make(map[uint16][]int),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16SliceintMapAllocSpace(size int) *Uint16SliceintMap {
	om := &Uint16SliceintMap{
		store:   make(map[uint16][]int, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16SliceintMap) Set(key uint16, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16SliceintMap) Get(key uint16) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16SliceintMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16SliceintMap) Len() int {
	return len(om.store)
}

func (om *Uint16SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceint8Map struct {
	store   map[uint16][]int8
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceint8Map() *Uint16Sliceint8Map {
	om := &Uint16Sliceint8Map{
		store:   make(map[uint16][]int8),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceint8MapAllocSpace(size int) *Uint16Sliceint8Map {
	om := &Uint16Sliceint8Map{
		store:   make(map[uint16][]int8, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceint8Map) Set(key uint16, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceint8Map) Get(key uint16) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceint8Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceint16Map struct {
	store   map[uint16][]int16
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceint16Map() *Uint16Sliceint16Map {
	om := &Uint16Sliceint16Map{
		store:   make(map[uint16][]int16),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceint16MapAllocSpace(size int) *Uint16Sliceint16Map {
	om := &Uint16Sliceint16Map{
		store:   make(map[uint16][]int16, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceint16Map) Set(key uint16, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceint16Map) Get(key uint16) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceint16Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceint32Map struct {
	store   map[uint16][]int32
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceint32Map() *Uint16Sliceint32Map {
	om := &Uint16Sliceint32Map{
		store:   make(map[uint16][]int32),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceint32MapAllocSpace(size int) *Uint16Sliceint32Map {
	om := &Uint16Sliceint32Map{
		store:   make(map[uint16][]int32, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceint32Map) Set(key uint16, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceint32Map) Get(key uint16) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceint32Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceint64Map struct {
	store   map[uint16][]int64
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceint64Map() *Uint16Sliceint64Map {
	om := &Uint16Sliceint64Map{
		store:   make(map[uint16][]int64),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceint64MapAllocSpace(size int) *Uint16Sliceint64Map {
	om := &Uint16Sliceint64Map{
		store:   make(map[uint16][]int64, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceint64Map) Set(key uint16, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceint64Map) Get(key uint16) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceint64Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceuint8Map struct {
	store   map[uint16][]uint8
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceuint8Map() *Uint16Sliceuint8Map {
	om := &Uint16Sliceuint8Map{
		store:   make(map[uint16][]uint8),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceuint8MapAllocSpace(size int) *Uint16Sliceuint8Map {
	om := &Uint16Sliceuint8Map{
		store:   make(map[uint16][]uint8, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceuint8Map) Set(key uint16, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceuint8Map) Get(key uint16) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceuint8Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceuint16Map struct {
	store   map[uint16][]uint16
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceuint16Map() *Uint16Sliceuint16Map {
	om := &Uint16Sliceuint16Map{
		store:   make(map[uint16][]uint16),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceuint16MapAllocSpace(size int) *Uint16Sliceuint16Map {
	om := &Uint16Sliceuint16Map{
		store:   make(map[uint16][]uint16, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceuint16Map) Set(key uint16, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceuint16Map) Get(key uint16) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceuint16Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceuint32Map struct {
	store   map[uint16][]uint32
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceuint32Map() *Uint16Sliceuint32Map {
	om := &Uint16Sliceuint32Map{
		store:   make(map[uint16][]uint32),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceuint32MapAllocSpace(size int) *Uint16Sliceuint32Map {
	om := &Uint16Sliceuint32Map{
		store:   make(map[uint16][]uint32, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceuint32Map) Set(key uint16, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceuint32Map) Get(key uint16) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceuint32Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16Sliceuint64Map struct {
	store   map[uint16][]uint64
	mapper  []uint16
	isWrite bool
}

func NewUint16Sliceuint64Map() *Uint16Sliceuint64Map {
	om := &Uint16Sliceuint64Map{
		store:   make(map[uint16][]uint64),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16Sliceuint64MapAllocSpace(size int) *Uint16Sliceuint64Map {
	om := &Uint16Sliceuint64Map{
		store:   make(map[uint16][]uint64, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16Sliceuint64Map) Set(key uint16, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16Sliceuint64Map) Get(key uint16) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16Sliceuint64Map) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Uint16Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint16SlicestringMap struct {
	store   map[uint16][]string
	mapper  []uint16
	isWrite bool
}

func NewUint16SlicestringMap() *Uint16SlicestringMap {
	om := &Uint16SlicestringMap{
		store:   make(map[uint16][]string),
		mapper:  make([]uint16, 0),
		isWrite: false,
	}
	return om
}

func NewUint16SlicestringMapAllocSpace(size int) *Uint16SlicestringMap {
	om := &Uint16SlicestringMap{
		store:   make(map[uint16][]string, size),
		mapper:  make([]uint16, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint16SlicestringMap) Set(key uint16, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint16SlicestringMap) Get(key uint16) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint16SlicestringMap) Delete(key uint16) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint16SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Uint16SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint16SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]uint16, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint16s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint16SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]uint16, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint16sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32IntMap struct {
	store   map[uint32]int
	mapper  []uint32
	isWrite bool
}

func NewUint32IntMap() *Uint32IntMap {
	om := &Uint32IntMap{
		store:   make(map[uint32]int),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32IntMapAllocSpace(size int) *Uint32IntMap {
	om := &Uint32IntMap{
		store:   make(map[uint32]int, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32IntMap) Set(key uint32, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32IntMap) Get(key uint32) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32IntMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32IntMap) Len() int {
	return len(om.store)
}

func (om *Uint32IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Int8Map struct {
	store   map[uint32]int8
	mapper  []uint32
	isWrite bool
}

func NewUint32Int8Map() *Uint32Int8Map {
	om := &Uint32Int8Map{
		store:   make(map[uint32]int8),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Int8MapAllocSpace(size int) *Uint32Int8Map {
	om := &Uint32Int8Map{
		store:   make(map[uint32]int8, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Int8Map) Set(key uint32, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Int8Map) Get(key uint32) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Int8Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Int8Map) Len() int {
	return len(om.store)
}

func (om *Uint32Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Int16Map struct {
	store   map[uint32]int16
	mapper  []uint32
	isWrite bool
}

func NewUint32Int16Map() *Uint32Int16Map {
	om := &Uint32Int16Map{
		store:   make(map[uint32]int16),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Int16MapAllocSpace(size int) *Uint32Int16Map {
	om := &Uint32Int16Map{
		store:   make(map[uint32]int16, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Int16Map) Set(key uint32, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Int16Map) Get(key uint32) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Int16Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Int16Map) Len() int {
	return len(om.store)
}

func (om *Uint32Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Int32Map struct {
	store   map[uint32]int32
	mapper  []uint32
	isWrite bool
}

func NewUint32Int32Map() *Uint32Int32Map {
	om := &Uint32Int32Map{
		store:   make(map[uint32]int32),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Int32MapAllocSpace(size int) *Uint32Int32Map {
	om := &Uint32Int32Map{
		store:   make(map[uint32]int32, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Int32Map) Set(key uint32, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Int32Map) Get(key uint32) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Int32Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Int32Map) Len() int {
	return len(om.store)
}

func (om *Uint32Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Int64Map struct {
	store   map[uint32]int64
	mapper  []uint32
	isWrite bool
}

func NewUint32Int64Map() *Uint32Int64Map {
	om := &Uint32Int64Map{
		store:   make(map[uint32]int64),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Int64MapAllocSpace(size int) *Uint32Int64Map {
	om := &Uint32Int64Map{
		store:   make(map[uint32]int64, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Int64Map) Set(key uint32, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Int64Map) Get(key uint32) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Int64Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Int64Map) Len() int {
	return len(om.store)
}

func (om *Uint32Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Uint8Map struct {
	store   map[uint32]uint8
	mapper  []uint32
	isWrite bool
}

func NewUint32Uint8Map() *Uint32Uint8Map {
	om := &Uint32Uint8Map{
		store:   make(map[uint32]uint8),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Uint8MapAllocSpace(size int) *Uint32Uint8Map {
	om := &Uint32Uint8Map{
		store:   make(map[uint32]uint8, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Uint8Map) Set(key uint32, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Uint8Map) Get(key uint32) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Uint8Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Uint8Map) Len() int {
	return len(om.store)
}

func (om *Uint32Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Uint16Map struct {
	store   map[uint32]uint16
	mapper  []uint32
	isWrite bool
}

func NewUint32Uint16Map() *Uint32Uint16Map {
	om := &Uint32Uint16Map{
		store:   make(map[uint32]uint16),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Uint16MapAllocSpace(size int) *Uint32Uint16Map {
	om := &Uint32Uint16Map{
		store:   make(map[uint32]uint16, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Uint16Map) Set(key uint32, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Uint16Map) Get(key uint32) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Uint16Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Uint16Map) Len() int {
	return len(om.store)
}

func (om *Uint32Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Uint32Map struct {
	store   map[uint32]uint32
	mapper  []uint32
	isWrite bool
}

func NewUint32Uint32Map() *Uint32Uint32Map {
	om := &Uint32Uint32Map{
		store:   make(map[uint32]uint32),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Uint32MapAllocSpace(size int) *Uint32Uint32Map {
	om := &Uint32Uint32Map{
		store:   make(map[uint32]uint32, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Uint32Map) Set(key uint32, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Uint32Map) Get(key uint32) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Uint32Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Uint32Map) Len() int {
	return len(om.store)
}

func (om *Uint32Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Uint64Map struct {
	store   map[uint32]uint64
	mapper  []uint32
	isWrite bool
}

func NewUint32Uint64Map() *Uint32Uint64Map {
	om := &Uint32Uint64Map{
		store:   make(map[uint32]uint64),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Uint64MapAllocSpace(size int) *Uint32Uint64Map {
	om := &Uint32Uint64Map{
		store:   make(map[uint32]uint64, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Uint64Map) Set(key uint32, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Uint64Map) Get(key uint32) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Uint64Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Uint64Map) Len() int {
	return len(om.store)
}

func (om *Uint32Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32StringMap struct {
	store   map[uint32]string
	mapper  []uint32
	isWrite bool
}

func NewUint32StringMap() *Uint32StringMap {
	om := &Uint32StringMap{
		store:   make(map[uint32]string),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32StringMapAllocSpace(size int) *Uint32StringMap {
	om := &Uint32StringMap{
		store:   make(map[uint32]string, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32StringMap) Set(key uint32, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32StringMap) Get(key uint32) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32StringMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32StringMap) Len() int {
	return len(om.store)
}

func (om *Uint32StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Float32Map struct {
	store   map[uint32]float32
	mapper  []uint32
	isWrite bool
}

func NewUint32Float32Map() *Uint32Float32Map {
	om := &Uint32Float32Map{
		store:   make(map[uint32]float32),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Float32MapAllocSpace(size int) *Uint32Float32Map {
	om := &Uint32Float32Map{
		store:   make(map[uint32]float32, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Float32Map) Set(key uint32, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Float32Map) Get(key uint32) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Float32Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Float32Map) Len() int {
	return len(om.store)
}

func (om *Uint32Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Float64Map struct {
	store   map[uint32]float64
	mapper  []uint32
	isWrite bool
}

func NewUint32Float64Map() *Uint32Float64Map {
	om := &Uint32Float64Map{
		store:   make(map[uint32]float64),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Float64MapAllocSpace(size int) *Uint32Float64Map {
	om := &Uint32Float64Map{
		store:   make(map[uint32]float64, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Float64Map) Set(key uint32, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Float64Map) Get(key uint32) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Float64Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Float64Map) Len() int {
	return len(om.store)
}

func (om *Uint32Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32InterfaceMap struct {
	store   map[uint32]interface{}
	mapper  []uint32
	isWrite bool
}

func NewUint32InterfaceMap() *Uint32InterfaceMap {
	om := &Uint32InterfaceMap{
		store:   make(map[uint32]interface{}),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32InterfaceMapAllocSpace(size int) *Uint32InterfaceMap {
	om := &Uint32InterfaceMap{
		store:   make(map[uint32]interface{}, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32InterfaceMap) Set(key uint32, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32InterfaceMap) Get(key uint32) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32InterfaceMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Uint32InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32BoolMap struct {
	store   map[uint32]bool
	mapper  []uint32
	isWrite bool
}

func NewUint32BoolMap() *Uint32BoolMap {
	om := &Uint32BoolMap{
		store:   make(map[uint32]bool),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32BoolMapAllocSpace(size int) *Uint32BoolMap {
	om := &Uint32BoolMap{
		store:   make(map[uint32]bool, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32BoolMap) Set(key uint32, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32BoolMap) Get(key uint32) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32BoolMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32BoolMap) Len() int {
	return len(om.store)
}

func (om *Uint32BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32StructMap struct {
	store   map[uint32]struct{}
	mapper  []uint32
	isWrite bool
}

func NewUint32StructMap() *Uint32StructMap {
	om := &Uint32StructMap{
		store:   make(map[uint32]struct{}),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32StructMapAllocSpace(size int) *Uint32StructMap {
	om := &Uint32StructMap{
		store:   make(map[uint32]struct{}, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32StructMap) Set(key uint32, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32StructMap) Get(key uint32) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32StructMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32StructMap) Len() int {
	return len(om.store)
}

func (om *Uint32StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32SliceintMap struct {
	store   map[uint32][]int
	mapper  []uint32
	isWrite bool
}

func NewUint32SliceintMap() *Uint32SliceintMap {
	om := &Uint32SliceintMap{
		store:   make(map[uint32][]int),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32SliceintMapAllocSpace(size int) *Uint32SliceintMap {
	om := &Uint32SliceintMap{
		store:   make(map[uint32][]int, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32SliceintMap) Set(key uint32, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32SliceintMap) Get(key uint32) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32SliceintMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32SliceintMap) Len() int {
	return len(om.store)
}

func (om *Uint32SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceint8Map struct {
	store   map[uint32][]int8
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceint8Map() *Uint32Sliceint8Map {
	om := &Uint32Sliceint8Map{
		store:   make(map[uint32][]int8),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceint8MapAllocSpace(size int) *Uint32Sliceint8Map {
	om := &Uint32Sliceint8Map{
		store:   make(map[uint32][]int8, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceint8Map) Set(key uint32, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceint8Map) Get(key uint32) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceint8Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceint16Map struct {
	store   map[uint32][]int16
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceint16Map() *Uint32Sliceint16Map {
	om := &Uint32Sliceint16Map{
		store:   make(map[uint32][]int16),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceint16MapAllocSpace(size int) *Uint32Sliceint16Map {
	om := &Uint32Sliceint16Map{
		store:   make(map[uint32][]int16, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceint16Map) Set(key uint32, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceint16Map) Get(key uint32) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceint16Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceint32Map struct {
	store   map[uint32][]int32
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceint32Map() *Uint32Sliceint32Map {
	om := &Uint32Sliceint32Map{
		store:   make(map[uint32][]int32),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceint32MapAllocSpace(size int) *Uint32Sliceint32Map {
	om := &Uint32Sliceint32Map{
		store:   make(map[uint32][]int32, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceint32Map) Set(key uint32, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceint32Map) Get(key uint32) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceint32Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceint64Map struct {
	store   map[uint32][]int64
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceint64Map() *Uint32Sliceint64Map {
	om := &Uint32Sliceint64Map{
		store:   make(map[uint32][]int64),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceint64MapAllocSpace(size int) *Uint32Sliceint64Map {
	om := &Uint32Sliceint64Map{
		store:   make(map[uint32][]int64, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceint64Map) Set(key uint32, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceint64Map) Get(key uint32) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceint64Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceuint8Map struct {
	store   map[uint32][]uint8
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceuint8Map() *Uint32Sliceuint8Map {
	om := &Uint32Sliceuint8Map{
		store:   make(map[uint32][]uint8),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceuint8MapAllocSpace(size int) *Uint32Sliceuint8Map {
	om := &Uint32Sliceuint8Map{
		store:   make(map[uint32][]uint8, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceuint8Map) Set(key uint32, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceuint8Map) Get(key uint32) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceuint8Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceuint16Map struct {
	store   map[uint32][]uint16
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceuint16Map() *Uint32Sliceuint16Map {
	om := &Uint32Sliceuint16Map{
		store:   make(map[uint32][]uint16),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceuint16MapAllocSpace(size int) *Uint32Sliceuint16Map {
	om := &Uint32Sliceuint16Map{
		store:   make(map[uint32][]uint16, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceuint16Map) Set(key uint32, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceuint16Map) Get(key uint32) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceuint16Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceuint32Map struct {
	store   map[uint32][]uint32
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceuint32Map() *Uint32Sliceuint32Map {
	om := &Uint32Sliceuint32Map{
		store:   make(map[uint32][]uint32),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceuint32MapAllocSpace(size int) *Uint32Sliceuint32Map {
	om := &Uint32Sliceuint32Map{
		store:   make(map[uint32][]uint32, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceuint32Map) Set(key uint32, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceuint32Map) Get(key uint32) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceuint32Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32Sliceuint64Map struct {
	store   map[uint32][]uint64
	mapper  []uint32
	isWrite bool
}

func NewUint32Sliceuint64Map() *Uint32Sliceuint64Map {
	om := &Uint32Sliceuint64Map{
		store:   make(map[uint32][]uint64),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32Sliceuint64MapAllocSpace(size int) *Uint32Sliceuint64Map {
	om := &Uint32Sliceuint64Map{
		store:   make(map[uint32][]uint64, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32Sliceuint64Map) Set(key uint32, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32Sliceuint64Map) Get(key uint32) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32Sliceuint64Map) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Uint32Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint32SlicestringMap struct {
	store   map[uint32][]string
	mapper  []uint32
	isWrite bool
}

func NewUint32SlicestringMap() *Uint32SlicestringMap {
	om := &Uint32SlicestringMap{
		store:   make(map[uint32][]string),
		mapper:  make([]uint32, 0),
		isWrite: false,
	}
	return om
}

func NewUint32SlicestringMapAllocSpace(size int) *Uint32SlicestringMap {
	om := &Uint32SlicestringMap{
		store:   make(map[uint32][]string, size),
		mapper:  make([]uint32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint32SlicestringMap) Set(key uint32, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint32SlicestringMap) Get(key uint32) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint32SlicestringMap) Delete(key uint32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint32SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Uint32SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint32SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]uint32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint32SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]uint32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64IntMap struct {
	store   map[uint64]int
	mapper  []uint64
	isWrite bool
}

func NewUint64IntMap() *Uint64IntMap {
	om := &Uint64IntMap{
		store:   make(map[uint64]int),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64IntMapAllocSpace(size int) *Uint64IntMap {
	om := &Uint64IntMap{
		store:   make(map[uint64]int, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64IntMap) Set(key uint64, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64IntMap) Get(key uint64) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64IntMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64IntMap) Len() int {
	return len(om.store)
}

func (om *Uint64IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Int8Map struct {
	store   map[uint64]int8
	mapper  []uint64
	isWrite bool
}

func NewUint64Int8Map() *Uint64Int8Map {
	om := &Uint64Int8Map{
		store:   make(map[uint64]int8),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Int8MapAllocSpace(size int) *Uint64Int8Map {
	om := &Uint64Int8Map{
		store:   make(map[uint64]int8, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Int8Map) Set(key uint64, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Int8Map) Get(key uint64) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Int8Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Int8Map) Len() int {
	return len(om.store)
}

func (om *Uint64Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Int16Map struct {
	store   map[uint64]int16
	mapper  []uint64
	isWrite bool
}

func NewUint64Int16Map() *Uint64Int16Map {
	om := &Uint64Int16Map{
		store:   make(map[uint64]int16),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Int16MapAllocSpace(size int) *Uint64Int16Map {
	om := &Uint64Int16Map{
		store:   make(map[uint64]int16, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Int16Map) Set(key uint64, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Int16Map) Get(key uint64) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Int16Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Int16Map) Len() int {
	return len(om.store)
}

func (om *Uint64Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Int32Map struct {
	store   map[uint64]int32
	mapper  []uint64
	isWrite bool
}

func NewUint64Int32Map() *Uint64Int32Map {
	om := &Uint64Int32Map{
		store:   make(map[uint64]int32),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Int32MapAllocSpace(size int) *Uint64Int32Map {
	om := &Uint64Int32Map{
		store:   make(map[uint64]int32, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Int32Map) Set(key uint64, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Int32Map) Get(key uint64) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Int32Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Int32Map) Len() int {
	return len(om.store)
}

func (om *Uint64Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Int64Map struct {
	store   map[uint64]int64
	mapper  []uint64
	isWrite bool
}

func NewUint64Int64Map() *Uint64Int64Map {
	om := &Uint64Int64Map{
		store:   make(map[uint64]int64),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Int64MapAllocSpace(size int) *Uint64Int64Map {
	om := &Uint64Int64Map{
		store:   make(map[uint64]int64, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Int64Map) Set(key uint64, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Int64Map) Get(key uint64) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Int64Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Int64Map) Len() int {
	return len(om.store)
}

func (om *Uint64Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Uint8Map struct {
	store   map[uint64]uint8
	mapper  []uint64
	isWrite bool
}

func NewUint64Uint8Map() *Uint64Uint8Map {
	om := &Uint64Uint8Map{
		store:   make(map[uint64]uint8),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Uint8MapAllocSpace(size int) *Uint64Uint8Map {
	om := &Uint64Uint8Map{
		store:   make(map[uint64]uint8, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Uint8Map) Set(key uint64, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Uint8Map) Get(key uint64) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Uint8Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Uint8Map) Len() int {
	return len(om.store)
}

func (om *Uint64Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Uint16Map struct {
	store   map[uint64]uint16
	mapper  []uint64
	isWrite bool
}

func NewUint64Uint16Map() *Uint64Uint16Map {
	om := &Uint64Uint16Map{
		store:   make(map[uint64]uint16),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Uint16MapAllocSpace(size int) *Uint64Uint16Map {
	om := &Uint64Uint16Map{
		store:   make(map[uint64]uint16, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Uint16Map) Set(key uint64, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Uint16Map) Get(key uint64) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Uint16Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Uint16Map) Len() int {
	return len(om.store)
}

func (om *Uint64Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Uint32Map struct {
	store   map[uint64]uint32
	mapper  []uint64
	isWrite bool
}

func NewUint64Uint32Map() *Uint64Uint32Map {
	om := &Uint64Uint32Map{
		store:   make(map[uint64]uint32),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Uint32MapAllocSpace(size int) *Uint64Uint32Map {
	om := &Uint64Uint32Map{
		store:   make(map[uint64]uint32, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Uint32Map) Set(key uint64, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Uint32Map) Get(key uint64) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Uint32Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Uint32Map) Len() int {
	return len(om.store)
}

func (om *Uint64Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Uint64Map struct {
	store   map[uint64]uint64
	mapper  []uint64
	isWrite bool
}

func NewUint64Uint64Map() *Uint64Uint64Map {
	om := &Uint64Uint64Map{
		store:   make(map[uint64]uint64),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Uint64MapAllocSpace(size int) *Uint64Uint64Map {
	om := &Uint64Uint64Map{
		store:   make(map[uint64]uint64, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Uint64Map) Set(key uint64, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Uint64Map) Get(key uint64) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Uint64Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Uint64Map) Len() int {
	return len(om.store)
}

func (om *Uint64Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64StringMap struct {
	store   map[uint64]string
	mapper  []uint64
	isWrite bool
}

func NewUint64StringMap() *Uint64StringMap {
	om := &Uint64StringMap{
		store:   make(map[uint64]string),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64StringMapAllocSpace(size int) *Uint64StringMap {
	om := &Uint64StringMap{
		store:   make(map[uint64]string, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64StringMap) Set(key uint64, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64StringMap) Get(key uint64) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64StringMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64StringMap) Len() int {
	return len(om.store)
}

func (om *Uint64StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Float32Map struct {
	store   map[uint64]float32
	mapper  []uint64
	isWrite bool
}

func NewUint64Float32Map() *Uint64Float32Map {
	om := &Uint64Float32Map{
		store:   make(map[uint64]float32),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Float32MapAllocSpace(size int) *Uint64Float32Map {
	om := &Uint64Float32Map{
		store:   make(map[uint64]float32, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Float32Map) Set(key uint64, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Float32Map) Get(key uint64) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Float32Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Float32Map) Len() int {
	return len(om.store)
}

func (om *Uint64Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Float64Map struct {
	store   map[uint64]float64
	mapper  []uint64
	isWrite bool
}

func NewUint64Float64Map() *Uint64Float64Map {
	om := &Uint64Float64Map{
		store:   make(map[uint64]float64),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Float64MapAllocSpace(size int) *Uint64Float64Map {
	om := &Uint64Float64Map{
		store:   make(map[uint64]float64, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Float64Map) Set(key uint64, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Float64Map) Get(key uint64) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Float64Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Float64Map) Len() int {
	return len(om.store)
}

func (om *Uint64Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64InterfaceMap struct {
	store   map[uint64]interface{}
	mapper  []uint64
	isWrite bool
}

func NewUint64InterfaceMap() *Uint64InterfaceMap {
	om := &Uint64InterfaceMap{
		store:   make(map[uint64]interface{}),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64InterfaceMapAllocSpace(size int) *Uint64InterfaceMap {
	om := &Uint64InterfaceMap{
		store:   make(map[uint64]interface{}, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64InterfaceMap) Set(key uint64, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64InterfaceMap) Get(key uint64) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64InterfaceMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Uint64InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64BoolMap struct {
	store   map[uint64]bool
	mapper  []uint64
	isWrite bool
}

func NewUint64BoolMap() *Uint64BoolMap {
	om := &Uint64BoolMap{
		store:   make(map[uint64]bool),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64BoolMapAllocSpace(size int) *Uint64BoolMap {
	om := &Uint64BoolMap{
		store:   make(map[uint64]bool, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64BoolMap) Set(key uint64, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64BoolMap) Get(key uint64) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64BoolMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64BoolMap) Len() int {
	return len(om.store)
}

func (om *Uint64BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64StructMap struct {
	store   map[uint64]struct{}
	mapper  []uint64
	isWrite bool
}

func NewUint64StructMap() *Uint64StructMap {
	om := &Uint64StructMap{
		store:   make(map[uint64]struct{}),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64StructMapAllocSpace(size int) *Uint64StructMap {
	om := &Uint64StructMap{
		store:   make(map[uint64]struct{}, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64StructMap) Set(key uint64, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64StructMap) Get(key uint64) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64StructMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64StructMap) Len() int {
	return len(om.store)
}

func (om *Uint64StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64SliceintMap struct {
	store   map[uint64][]int
	mapper  []uint64
	isWrite bool
}

func NewUint64SliceintMap() *Uint64SliceintMap {
	om := &Uint64SliceintMap{
		store:   make(map[uint64][]int),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64SliceintMapAllocSpace(size int) *Uint64SliceintMap {
	om := &Uint64SliceintMap{
		store:   make(map[uint64][]int, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64SliceintMap) Set(key uint64, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64SliceintMap) Get(key uint64) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64SliceintMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64SliceintMap) Len() int {
	return len(om.store)
}

func (om *Uint64SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceint8Map struct {
	store   map[uint64][]int8
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceint8Map() *Uint64Sliceint8Map {
	om := &Uint64Sliceint8Map{
		store:   make(map[uint64][]int8),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceint8MapAllocSpace(size int) *Uint64Sliceint8Map {
	om := &Uint64Sliceint8Map{
		store:   make(map[uint64][]int8, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceint8Map) Set(key uint64, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceint8Map) Get(key uint64) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceint8Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceint16Map struct {
	store   map[uint64][]int16
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceint16Map() *Uint64Sliceint16Map {
	om := &Uint64Sliceint16Map{
		store:   make(map[uint64][]int16),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceint16MapAllocSpace(size int) *Uint64Sliceint16Map {
	om := &Uint64Sliceint16Map{
		store:   make(map[uint64][]int16, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceint16Map) Set(key uint64, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceint16Map) Get(key uint64) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceint16Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceint32Map struct {
	store   map[uint64][]int32
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceint32Map() *Uint64Sliceint32Map {
	om := &Uint64Sliceint32Map{
		store:   make(map[uint64][]int32),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceint32MapAllocSpace(size int) *Uint64Sliceint32Map {
	om := &Uint64Sliceint32Map{
		store:   make(map[uint64][]int32, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceint32Map) Set(key uint64, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceint32Map) Get(key uint64) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceint32Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceint64Map struct {
	store   map[uint64][]int64
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceint64Map() *Uint64Sliceint64Map {
	om := &Uint64Sliceint64Map{
		store:   make(map[uint64][]int64),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceint64MapAllocSpace(size int) *Uint64Sliceint64Map {
	om := &Uint64Sliceint64Map{
		store:   make(map[uint64][]int64, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceint64Map) Set(key uint64, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceint64Map) Get(key uint64) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceint64Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceuint8Map struct {
	store   map[uint64][]uint8
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceuint8Map() *Uint64Sliceuint8Map {
	om := &Uint64Sliceuint8Map{
		store:   make(map[uint64][]uint8),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceuint8MapAllocSpace(size int) *Uint64Sliceuint8Map {
	om := &Uint64Sliceuint8Map{
		store:   make(map[uint64][]uint8, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceuint8Map) Set(key uint64, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceuint8Map) Get(key uint64) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceuint8Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceuint16Map struct {
	store   map[uint64][]uint16
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceuint16Map() *Uint64Sliceuint16Map {
	om := &Uint64Sliceuint16Map{
		store:   make(map[uint64][]uint16),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceuint16MapAllocSpace(size int) *Uint64Sliceuint16Map {
	om := &Uint64Sliceuint16Map{
		store:   make(map[uint64][]uint16, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceuint16Map) Set(key uint64, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceuint16Map) Get(key uint64) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceuint16Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceuint32Map struct {
	store   map[uint64][]uint32
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceuint32Map() *Uint64Sliceuint32Map {
	om := &Uint64Sliceuint32Map{
		store:   make(map[uint64][]uint32),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceuint32MapAllocSpace(size int) *Uint64Sliceuint32Map {
	om := &Uint64Sliceuint32Map{
		store:   make(map[uint64][]uint32, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceuint32Map) Set(key uint64, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceuint32Map) Get(key uint64) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceuint32Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64Sliceuint64Map struct {
	store   map[uint64][]uint64
	mapper  []uint64
	isWrite bool
}

func NewUint64Sliceuint64Map() *Uint64Sliceuint64Map {
	om := &Uint64Sliceuint64Map{
		store:   make(map[uint64][]uint64),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64Sliceuint64MapAllocSpace(size int) *Uint64Sliceuint64Map {
	om := &Uint64Sliceuint64Map{
		store:   make(map[uint64][]uint64, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64Sliceuint64Map) Set(key uint64, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64Sliceuint64Map) Get(key uint64) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64Sliceuint64Map) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Uint64Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Uint64SlicestringMap struct {
	store   map[uint64][]string
	mapper  []uint64
	isWrite bool
}

func NewUint64SlicestringMap() *Uint64SlicestringMap {
	om := &Uint64SlicestringMap{
		store:   make(map[uint64][]string),
		mapper:  make([]uint64, 0),
		isWrite: false,
	}
	return om
}

func NewUint64SlicestringMapAllocSpace(size int) *Uint64SlicestringMap {
	om := &Uint64SlicestringMap{
		store:   make(map[uint64][]string, size),
		mapper:  make([]uint64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Uint64SlicestringMap) Set(key uint64, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Uint64SlicestringMap) Get(key uint64) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Uint64SlicestringMap) Delete(key uint64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Uint64SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Uint64SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Uint64SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]uint64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Uint64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Uint64SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]uint64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Uint64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64IntMap struct {
	store   map[float64]int
	mapper  []float64
	isWrite bool
}

func NewFloat64IntMap() *Float64IntMap {
	om := &Float64IntMap{
		store:   make(map[float64]int),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64IntMapAllocSpace(size int) *Float64IntMap {
	om := &Float64IntMap{
		store:   make(map[float64]int, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64IntMap) Set(key float64, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64IntMap) Get(key float64) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64IntMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64IntMap) Len() int {
	return len(om.store)
}

func (om *Float64IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Int8Map struct {
	store   map[float64]int8
	mapper  []float64
	isWrite bool
}

func NewFloat64Int8Map() *Float64Int8Map {
	om := &Float64Int8Map{
		store:   make(map[float64]int8),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Int8MapAllocSpace(size int) *Float64Int8Map {
	om := &Float64Int8Map{
		store:   make(map[float64]int8, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Int8Map) Set(key float64, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Int8Map) Get(key float64) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Int8Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Int8Map) Len() int {
	return len(om.store)
}

func (om *Float64Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Int16Map struct {
	store   map[float64]int16
	mapper  []float64
	isWrite bool
}

func NewFloat64Int16Map() *Float64Int16Map {
	om := &Float64Int16Map{
		store:   make(map[float64]int16),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Int16MapAllocSpace(size int) *Float64Int16Map {
	om := &Float64Int16Map{
		store:   make(map[float64]int16, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Int16Map) Set(key float64, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Int16Map) Get(key float64) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Int16Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Int16Map) Len() int {
	return len(om.store)
}

func (om *Float64Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Int32Map struct {
	store   map[float64]int32
	mapper  []float64
	isWrite bool
}

func NewFloat64Int32Map() *Float64Int32Map {
	om := &Float64Int32Map{
		store:   make(map[float64]int32),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Int32MapAllocSpace(size int) *Float64Int32Map {
	om := &Float64Int32Map{
		store:   make(map[float64]int32, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Int32Map) Set(key float64, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Int32Map) Get(key float64) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Int32Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Int32Map) Len() int {
	return len(om.store)
}

func (om *Float64Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Int64Map struct {
	store   map[float64]int64
	mapper  []float64
	isWrite bool
}

func NewFloat64Int64Map() *Float64Int64Map {
	om := &Float64Int64Map{
		store:   make(map[float64]int64),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Int64MapAllocSpace(size int) *Float64Int64Map {
	om := &Float64Int64Map{
		store:   make(map[float64]int64, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Int64Map) Set(key float64, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Int64Map) Get(key float64) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Int64Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Int64Map) Len() int {
	return len(om.store)
}

func (om *Float64Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Uint8Map struct {
	store   map[float64]uint8
	mapper  []float64
	isWrite bool
}

func NewFloat64Uint8Map() *Float64Uint8Map {
	om := &Float64Uint8Map{
		store:   make(map[float64]uint8),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Uint8MapAllocSpace(size int) *Float64Uint8Map {
	om := &Float64Uint8Map{
		store:   make(map[float64]uint8, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Uint8Map) Set(key float64, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Uint8Map) Get(key float64) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Uint8Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Uint8Map) Len() int {
	return len(om.store)
}

func (om *Float64Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Uint16Map struct {
	store   map[float64]uint16
	mapper  []float64
	isWrite bool
}

func NewFloat64Uint16Map() *Float64Uint16Map {
	om := &Float64Uint16Map{
		store:   make(map[float64]uint16),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Uint16MapAllocSpace(size int) *Float64Uint16Map {
	om := &Float64Uint16Map{
		store:   make(map[float64]uint16, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Uint16Map) Set(key float64, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Uint16Map) Get(key float64) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Uint16Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Uint16Map) Len() int {
	return len(om.store)
}

func (om *Float64Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Uint32Map struct {
	store   map[float64]uint32
	mapper  []float64
	isWrite bool
}

func NewFloat64Uint32Map() *Float64Uint32Map {
	om := &Float64Uint32Map{
		store:   make(map[float64]uint32),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Uint32MapAllocSpace(size int) *Float64Uint32Map {
	om := &Float64Uint32Map{
		store:   make(map[float64]uint32, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Uint32Map) Set(key float64, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Uint32Map) Get(key float64) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Uint32Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Uint32Map) Len() int {
	return len(om.store)
}

func (om *Float64Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Uint64Map struct {
	store   map[float64]uint64
	mapper  []float64
	isWrite bool
}

func NewFloat64Uint64Map() *Float64Uint64Map {
	om := &Float64Uint64Map{
		store:   make(map[float64]uint64),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Uint64MapAllocSpace(size int) *Float64Uint64Map {
	om := &Float64Uint64Map{
		store:   make(map[float64]uint64, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Uint64Map) Set(key float64, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Uint64Map) Get(key float64) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Uint64Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Uint64Map) Len() int {
	return len(om.store)
}

func (om *Float64Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64StringMap struct {
	store   map[float64]string
	mapper  []float64
	isWrite bool
}

func NewFloat64StringMap() *Float64StringMap {
	om := &Float64StringMap{
		store:   make(map[float64]string),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64StringMapAllocSpace(size int) *Float64StringMap {
	om := &Float64StringMap{
		store:   make(map[float64]string, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64StringMap) Set(key float64, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64StringMap) Get(key float64) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64StringMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64StringMap) Len() int {
	return len(om.store)
}

func (om *Float64StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Float32Map struct {
	store   map[float64]float32
	mapper  []float64
	isWrite bool
}

func NewFloat64Float32Map() *Float64Float32Map {
	om := &Float64Float32Map{
		store:   make(map[float64]float32),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Float32MapAllocSpace(size int) *Float64Float32Map {
	om := &Float64Float32Map{
		store:   make(map[float64]float32, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Float32Map) Set(key float64, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Float32Map) Get(key float64) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Float32Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Float32Map) Len() int {
	return len(om.store)
}

func (om *Float64Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Float64Map struct {
	store   map[float64]float64
	mapper  []float64
	isWrite bool
}

func NewFloat64Float64Map() *Float64Float64Map {
	om := &Float64Float64Map{
		store:   make(map[float64]float64),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Float64MapAllocSpace(size int) *Float64Float64Map {
	om := &Float64Float64Map{
		store:   make(map[float64]float64, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Float64Map) Set(key float64, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Float64Map) Get(key float64) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Float64Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Float64Map) Len() int {
	return len(om.store)
}

func (om *Float64Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64InterfaceMap struct {
	store   map[float64]interface{}
	mapper  []float64
	isWrite bool
}

func NewFloat64InterfaceMap() *Float64InterfaceMap {
	om := &Float64InterfaceMap{
		store:   make(map[float64]interface{}),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64InterfaceMapAllocSpace(size int) *Float64InterfaceMap {
	om := &Float64InterfaceMap{
		store:   make(map[float64]interface{}, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64InterfaceMap) Set(key float64, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64InterfaceMap) Get(key float64) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64InterfaceMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Float64InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64BoolMap struct {
	store   map[float64]bool
	mapper  []float64
	isWrite bool
}

func NewFloat64BoolMap() *Float64BoolMap {
	om := &Float64BoolMap{
		store:   make(map[float64]bool),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64BoolMapAllocSpace(size int) *Float64BoolMap {
	om := &Float64BoolMap{
		store:   make(map[float64]bool, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64BoolMap) Set(key float64, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64BoolMap) Get(key float64) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64BoolMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64BoolMap) Len() int {
	return len(om.store)
}

func (om *Float64BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64StructMap struct {
	store   map[float64]struct{}
	mapper  []float64
	isWrite bool
}

func NewFloat64StructMap() *Float64StructMap {
	om := &Float64StructMap{
		store:   make(map[float64]struct{}),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64StructMapAllocSpace(size int) *Float64StructMap {
	om := &Float64StructMap{
		store:   make(map[float64]struct{}, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64StructMap) Set(key float64, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64StructMap) Get(key float64) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64StructMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64StructMap) Len() int {
	return len(om.store)
}

func (om *Float64StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64SliceintMap struct {
	store   map[float64][]int
	mapper  []float64
	isWrite bool
}

func NewFloat64SliceintMap() *Float64SliceintMap {
	om := &Float64SliceintMap{
		store:   make(map[float64][]int),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64SliceintMapAllocSpace(size int) *Float64SliceintMap {
	om := &Float64SliceintMap{
		store:   make(map[float64][]int, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64SliceintMap) Set(key float64, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64SliceintMap) Get(key float64) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64SliceintMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64SliceintMap) Len() int {
	return len(om.store)
}

func (om *Float64SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceint8Map struct {
	store   map[float64][]int8
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceint8Map() *Float64Sliceint8Map {
	om := &Float64Sliceint8Map{
		store:   make(map[float64][]int8),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceint8MapAllocSpace(size int) *Float64Sliceint8Map {
	om := &Float64Sliceint8Map{
		store:   make(map[float64][]int8, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceint8Map) Set(key float64, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceint8Map) Get(key float64) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceint8Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceint16Map struct {
	store   map[float64][]int16
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceint16Map() *Float64Sliceint16Map {
	om := &Float64Sliceint16Map{
		store:   make(map[float64][]int16),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceint16MapAllocSpace(size int) *Float64Sliceint16Map {
	om := &Float64Sliceint16Map{
		store:   make(map[float64][]int16, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceint16Map) Set(key float64, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceint16Map) Get(key float64) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceint16Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceint32Map struct {
	store   map[float64][]int32
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceint32Map() *Float64Sliceint32Map {
	om := &Float64Sliceint32Map{
		store:   make(map[float64][]int32),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceint32MapAllocSpace(size int) *Float64Sliceint32Map {
	om := &Float64Sliceint32Map{
		store:   make(map[float64][]int32, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceint32Map) Set(key float64, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceint32Map) Get(key float64) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceint32Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceint64Map struct {
	store   map[float64][]int64
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceint64Map() *Float64Sliceint64Map {
	om := &Float64Sliceint64Map{
		store:   make(map[float64][]int64),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceint64MapAllocSpace(size int) *Float64Sliceint64Map {
	om := &Float64Sliceint64Map{
		store:   make(map[float64][]int64, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceint64Map) Set(key float64, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceint64Map) Get(key float64) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceint64Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceuint8Map struct {
	store   map[float64][]uint8
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceuint8Map() *Float64Sliceuint8Map {
	om := &Float64Sliceuint8Map{
		store:   make(map[float64][]uint8),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceuint8MapAllocSpace(size int) *Float64Sliceuint8Map {
	om := &Float64Sliceuint8Map{
		store:   make(map[float64][]uint8, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceuint8Map) Set(key float64, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceuint8Map) Get(key float64) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceuint8Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceuint16Map struct {
	store   map[float64][]uint16
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceuint16Map() *Float64Sliceuint16Map {
	om := &Float64Sliceuint16Map{
		store:   make(map[float64][]uint16),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceuint16MapAllocSpace(size int) *Float64Sliceuint16Map {
	om := &Float64Sliceuint16Map{
		store:   make(map[float64][]uint16, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceuint16Map) Set(key float64, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceuint16Map) Get(key float64) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceuint16Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceuint32Map struct {
	store   map[float64][]uint32
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceuint32Map() *Float64Sliceuint32Map {
	om := &Float64Sliceuint32Map{
		store:   make(map[float64][]uint32),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceuint32MapAllocSpace(size int) *Float64Sliceuint32Map {
	om := &Float64Sliceuint32Map{
		store:   make(map[float64][]uint32, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceuint32Map) Set(key float64, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceuint32Map) Get(key float64) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceuint32Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64Sliceuint64Map struct {
	store   map[float64][]uint64
	mapper  []float64
	isWrite bool
}

func NewFloat64Sliceuint64Map() *Float64Sliceuint64Map {
	om := &Float64Sliceuint64Map{
		store:   make(map[float64][]uint64),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64Sliceuint64MapAllocSpace(size int) *Float64Sliceuint64Map {
	om := &Float64Sliceuint64Map{
		store:   make(map[float64][]uint64, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64Sliceuint64Map) Set(key float64, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64Sliceuint64Map) Get(key float64) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64Sliceuint64Map) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Float64Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float64SlicestringMap struct {
	store   map[float64][]string
	mapper  []float64
	isWrite bool
}

func NewFloat64SlicestringMap() *Float64SlicestringMap {
	om := &Float64SlicestringMap{
		store:   make(map[float64][]string),
		mapper:  make([]float64, 0),
		isWrite: false,
	}
	return om
}

func NewFloat64SlicestringMapAllocSpace(size int) *Float64SlicestringMap {
	om := &Float64SlicestringMap{
		store:   make(map[float64][]string, size),
		mapper:  make([]float64, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float64SlicestringMap) Set(key float64, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float64SlicestringMap) Get(key float64) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float64SlicestringMap) Delete(key float64) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float64SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Float64SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float64SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]float64, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float64s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float64SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]float64, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float64sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32IntMap struct {
	store   map[float32]int
	mapper  []float32
	isWrite bool
}

func NewFloat32IntMap() *Float32IntMap {
	om := &Float32IntMap{
		store:   make(map[float32]int),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32IntMapAllocSpace(size int) *Float32IntMap {
	om := &Float32IntMap{
		store:   make(map[float32]int, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32IntMap) Set(key float32, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32IntMap) Get(key float32) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32IntMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32IntMap) Len() int {
	return len(om.store)
}

func (om *Float32IntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32IntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32IntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Int8Map struct {
	store   map[float32]int8
	mapper  []float32
	isWrite bool
}

func NewFloat32Int8Map() *Float32Int8Map {
	om := &Float32Int8Map{
		store:   make(map[float32]int8),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Int8MapAllocSpace(size int) *Float32Int8Map {
	om := &Float32Int8Map{
		store:   make(map[float32]int8, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Int8Map) Set(key float32, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Int8Map) Get(key float32) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Int8Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Int8Map) Len() int {
	return len(om.store)
}

func (om *Float32Int8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Int8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Int8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Int16Map struct {
	store   map[float32]int16
	mapper  []float32
	isWrite bool
}

func NewFloat32Int16Map() *Float32Int16Map {
	om := &Float32Int16Map{
		store:   make(map[float32]int16),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Int16MapAllocSpace(size int) *Float32Int16Map {
	om := &Float32Int16Map{
		store:   make(map[float32]int16, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Int16Map) Set(key float32, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Int16Map) Get(key float32) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Int16Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Int16Map) Len() int {
	return len(om.store)
}

func (om *Float32Int16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Int16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Int16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Int32Map struct {
	store   map[float32]int32
	mapper  []float32
	isWrite bool
}

func NewFloat32Int32Map() *Float32Int32Map {
	om := &Float32Int32Map{
		store:   make(map[float32]int32),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Int32MapAllocSpace(size int) *Float32Int32Map {
	om := &Float32Int32Map{
		store:   make(map[float32]int32, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Int32Map) Set(key float32, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Int32Map) Get(key float32) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Int32Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Int32Map) Len() int {
	return len(om.store)
}

func (om *Float32Int32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Int32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Int32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Int64Map struct {
	store   map[float32]int64
	mapper  []float32
	isWrite bool
}

func NewFloat32Int64Map() *Float32Int64Map {
	om := &Float32Int64Map{
		store:   make(map[float32]int64),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Int64MapAllocSpace(size int) *Float32Int64Map {
	om := &Float32Int64Map{
		store:   make(map[float32]int64, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Int64Map) Set(key float32, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Int64Map) Get(key float32) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Int64Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Int64Map) Len() int {
	return len(om.store)
}

func (om *Float32Int64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Int64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Int64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Uint8Map struct {
	store   map[float32]uint8
	mapper  []float32
	isWrite bool
}

func NewFloat32Uint8Map() *Float32Uint8Map {
	om := &Float32Uint8Map{
		store:   make(map[float32]uint8),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Uint8MapAllocSpace(size int) *Float32Uint8Map {
	om := &Float32Uint8Map{
		store:   make(map[float32]uint8, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Uint8Map) Set(key float32, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Uint8Map) Get(key float32) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Uint8Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Uint8Map) Len() int {
	return len(om.store)
}

func (om *Float32Uint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Uint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Uint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Uint16Map struct {
	store   map[float32]uint16
	mapper  []float32
	isWrite bool
}

func NewFloat32Uint16Map() *Float32Uint16Map {
	om := &Float32Uint16Map{
		store:   make(map[float32]uint16),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Uint16MapAllocSpace(size int) *Float32Uint16Map {
	om := &Float32Uint16Map{
		store:   make(map[float32]uint16, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Uint16Map) Set(key float32, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Uint16Map) Get(key float32) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Uint16Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Uint16Map) Len() int {
	return len(om.store)
}

func (om *Float32Uint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Uint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Uint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Uint32Map struct {
	store   map[float32]uint32
	mapper  []float32
	isWrite bool
}

func NewFloat32Uint32Map() *Float32Uint32Map {
	om := &Float32Uint32Map{
		store:   make(map[float32]uint32),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Uint32MapAllocSpace(size int) *Float32Uint32Map {
	om := &Float32Uint32Map{
		store:   make(map[float32]uint32, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Uint32Map) Set(key float32, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Uint32Map) Get(key float32) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Uint32Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Uint32Map) Len() int {
	return len(om.store)
}

func (om *Float32Uint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Uint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Uint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Uint64Map struct {
	store   map[float32]uint64
	mapper  []float32
	isWrite bool
}

func NewFloat32Uint64Map() *Float32Uint64Map {
	om := &Float32Uint64Map{
		store:   make(map[float32]uint64),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Uint64MapAllocSpace(size int) *Float32Uint64Map {
	om := &Float32Uint64Map{
		store:   make(map[float32]uint64, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Uint64Map) Set(key float32, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Uint64Map) Get(key float32) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Uint64Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Uint64Map) Len() int {
	return len(om.store)
}

func (om *Float32Uint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Uint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Uint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32StringMap struct {
	store   map[float32]string
	mapper  []float32
	isWrite bool
}

func NewFloat32StringMap() *Float32StringMap {
	om := &Float32StringMap{
		store:   make(map[float32]string),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32StringMapAllocSpace(size int) *Float32StringMap {
	om := &Float32StringMap{
		store:   make(map[float32]string, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32StringMap) Set(key float32, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32StringMap) Get(key float32) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32StringMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32StringMap) Len() int {
	return len(om.store)
}

func (om *Float32StringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32StringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32StringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Float32Map struct {
	store   map[float32]float32
	mapper  []float32
	isWrite bool
}

func NewFloat32Float32Map() *Float32Float32Map {
	om := &Float32Float32Map{
		store:   make(map[float32]float32),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Float32MapAllocSpace(size int) *Float32Float32Map {
	om := &Float32Float32Map{
		store:   make(map[float32]float32, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Float32Map) Set(key float32, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Float32Map) Get(key float32) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Float32Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Float32Map) Len() int {
	return len(om.store)
}

func (om *Float32Float32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Float32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Float32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Float64Map struct {
	store   map[float32]float64
	mapper  []float32
	isWrite bool
}

func NewFloat32Float64Map() *Float32Float64Map {
	om := &Float32Float64Map{
		store:   make(map[float32]float64),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Float64MapAllocSpace(size int) *Float32Float64Map {
	om := &Float32Float64Map{
		store:   make(map[float32]float64, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Float64Map) Set(key float32, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Float64Map) Get(key float32) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Float64Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Float64Map) Len() int {
	return len(om.store)
}

func (om *Float32Float64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Float64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Float64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32InterfaceMap struct {
	store   map[float32]interface{}
	mapper  []float32
	isWrite bool
}

func NewFloat32InterfaceMap() *Float32InterfaceMap {
	om := &Float32InterfaceMap{
		store:   make(map[float32]interface{}),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32InterfaceMapAllocSpace(size int) *Float32InterfaceMap {
	om := &Float32InterfaceMap{
		store:   make(map[float32]interface{}, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32InterfaceMap) Set(key float32, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32InterfaceMap) Get(key float32) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32InterfaceMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32InterfaceMap) Len() int {
	return len(om.store)
}

func (om *Float32InterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32InterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32InterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32BoolMap struct {
	store   map[float32]bool
	mapper  []float32
	isWrite bool
}

func NewFloat32BoolMap() *Float32BoolMap {
	om := &Float32BoolMap{
		store:   make(map[float32]bool),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32BoolMapAllocSpace(size int) *Float32BoolMap {
	om := &Float32BoolMap{
		store:   make(map[float32]bool, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32BoolMap) Set(key float32, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32BoolMap) Get(key float32) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32BoolMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32BoolMap) Len() int {
	return len(om.store)
}

func (om *Float32BoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32BoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32BoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32StructMap struct {
	store   map[float32]struct{}
	mapper  []float32
	isWrite bool
}

func NewFloat32StructMap() *Float32StructMap {
	om := &Float32StructMap{
		store:   make(map[float32]struct{}),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32StructMapAllocSpace(size int) *Float32StructMap {
	om := &Float32StructMap{
		store:   make(map[float32]struct{}, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32StructMap) Set(key float32, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32StructMap) Get(key float32) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32StructMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32StructMap) Len() int {
	return len(om.store)
}

func (om *Float32StructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32StructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32StructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32SliceintMap struct {
	store   map[float32][]int
	mapper  []float32
	isWrite bool
}

func NewFloat32SliceintMap() *Float32SliceintMap {
	om := &Float32SliceintMap{
		store:   make(map[float32][]int),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32SliceintMapAllocSpace(size int) *Float32SliceintMap {
	om := &Float32SliceintMap{
		store:   make(map[float32][]int, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32SliceintMap) Set(key float32, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32SliceintMap) Get(key float32) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32SliceintMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32SliceintMap) Len() int {
	return len(om.store)
}

func (om *Float32SliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32SliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32SliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceint8Map struct {
	store   map[float32][]int8
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceint8Map() *Float32Sliceint8Map {
	om := &Float32Sliceint8Map{
		store:   make(map[float32][]int8),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceint8MapAllocSpace(size int) *Float32Sliceint8Map {
	om := &Float32Sliceint8Map{
		store:   make(map[float32][]int8, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceint8Map) Set(key float32, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceint8Map) Get(key float32) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceint8Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceint8Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceint16Map struct {
	store   map[float32][]int16
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceint16Map() *Float32Sliceint16Map {
	om := &Float32Sliceint16Map{
		store:   make(map[float32][]int16),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceint16MapAllocSpace(size int) *Float32Sliceint16Map {
	om := &Float32Sliceint16Map{
		store:   make(map[float32][]int16, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceint16Map) Set(key float32, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceint16Map) Get(key float32) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceint16Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceint16Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceint32Map struct {
	store   map[float32][]int32
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceint32Map() *Float32Sliceint32Map {
	om := &Float32Sliceint32Map{
		store:   make(map[float32][]int32),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceint32MapAllocSpace(size int) *Float32Sliceint32Map {
	om := &Float32Sliceint32Map{
		store:   make(map[float32][]int32, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceint32Map) Set(key float32, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceint32Map) Get(key float32) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceint32Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceint32Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceint64Map struct {
	store   map[float32][]int64
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceint64Map() *Float32Sliceint64Map {
	om := &Float32Sliceint64Map{
		store:   make(map[float32][]int64),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceint64MapAllocSpace(size int) *Float32Sliceint64Map {
	om := &Float32Sliceint64Map{
		store:   make(map[float32][]int64, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceint64Map) Set(key float32, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceint64Map) Get(key float32) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceint64Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceint64Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceuint8Map struct {
	store   map[float32][]uint8
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceuint8Map() *Float32Sliceuint8Map {
	om := &Float32Sliceuint8Map{
		store:   make(map[float32][]uint8),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceuint8MapAllocSpace(size int) *Float32Sliceuint8Map {
	om := &Float32Sliceuint8Map{
		store:   make(map[float32][]uint8, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceuint8Map) Set(key float32, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceuint8Map) Get(key float32) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceuint8Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceuint8Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceuint16Map struct {
	store   map[float32][]uint16
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceuint16Map() *Float32Sliceuint16Map {
	om := &Float32Sliceuint16Map{
		store:   make(map[float32][]uint16),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceuint16MapAllocSpace(size int) *Float32Sliceuint16Map {
	om := &Float32Sliceuint16Map{
		store:   make(map[float32][]uint16, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceuint16Map) Set(key float32, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceuint16Map) Get(key float32) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceuint16Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceuint16Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceuint32Map struct {
	store   map[float32][]uint32
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceuint32Map() *Float32Sliceuint32Map {
	om := &Float32Sliceuint32Map{
		store:   make(map[float32][]uint32),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceuint32MapAllocSpace(size int) *Float32Sliceuint32Map {
	om := &Float32Sliceuint32Map{
		store:   make(map[float32][]uint32, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceuint32Map) Set(key float32, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceuint32Map) Get(key float32) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceuint32Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceuint32Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32Sliceuint64Map struct {
	store   map[float32][]uint64
	mapper  []float32
	isWrite bool
}

func NewFloat32Sliceuint64Map() *Float32Sliceuint64Map {
	om := &Float32Sliceuint64Map{
		store:   make(map[float32][]uint64),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32Sliceuint64MapAllocSpace(size int) *Float32Sliceuint64Map {
	om := &Float32Sliceuint64Map{
		store:   make(map[float32][]uint64, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32Sliceuint64Map) Set(key float32, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32Sliceuint64Map) Get(key float32) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32Sliceuint64Map) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32Sliceuint64Map) Len() int {
	return len(om.store)
}

func (om *Float32Sliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32Sliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type Float32SlicestringMap struct {
	store   map[float32][]string
	mapper  []float32
	isWrite bool
}

func NewFloat32SlicestringMap() *Float32SlicestringMap {
	om := &Float32SlicestringMap{
		store:   make(map[float32][]string),
		mapper:  make([]float32, 0),
		isWrite: false,
	}
	return om
}

func NewFloat32SlicestringMapAllocSpace(size int) *Float32SlicestringMap {
	om := &Float32SlicestringMap{
		store:   make(map[float32][]string, size),
		mapper:  make([]float32, 0, size),
		isWrite: false,
	}
	return om
}

func (om *Float32SlicestringMap) Set(key float32, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *Float32SlicestringMap) Get(key float32) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *Float32SlicestringMap) Delete(key float32) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *Float32SlicestringMap) Len() int {
	return len(om.store)
}

func (om *Float32SlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *Float32SlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]float32, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Float32s(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *Float32SlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]float32, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.Float32sRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringIntMap struct {
	store   map[string]int
	mapper  []string
	isWrite bool
}

func NewStringIntMap() *StringIntMap {
	om := &StringIntMap{
		store:   make(map[string]int),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringIntMapAllocSpace(size int) *StringIntMap {
	om := &StringIntMap{
		store:   make(map[string]int, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringIntMap) Set(key string, value int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringIntMap) Get(key string) (int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringIntMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringIntMap) Len() int {
	return len(om.store)
}

func (om *StringIntMap) IterFunc() func() (int, bool) {
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringIntMap) IterFuncAsc() func() (int, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringIntMap) IterFuncDesc() func() (int, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringInt8Map struct {
	store   map[string]int8
	mapper  []string
	isWrite bool
}

func NewStringInt8Map() *StringInt8Map {
	om := &StringInt8Map{
		store:   make(map[string]int8),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringInt8MapAllocSpace(size int) *StringInt8Map {
	om := &StringInt8Map{
		store:   make(map[string]int8, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringInt8Map) Set(key string, value int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringInt8Map) Get(key string) (int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringInt8Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringInt8Map) Len() int {
	return len(om.store)
}

func (om *StringInt8Map) IterFunc() func() (int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringInt8Map) IterFuncAsc() func() (int8, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringInt8Map) IterFuncDesc() func() (int8, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringInt16Map struct {
	store   map[string]int16
	mapper  []string
	isWrite bool
}

func NewStringInt16Map() *StringInt16Map {
	om := &StringInt16Map{
		store:   make(map[string]int16),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringInt16MapAllocSpace(size int) *StringInt16Map {
	om := &StringInt16Map{
		store:   make(map[string]int16, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringInt16Map) Set(key string, value int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringInt16Map) Get(key string) (int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringInt16Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringInt16Map) Len() int {
	return len(om.store)
}

func (om *StringInt16Map) IterFunc() func() (int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringInt16Map) IterFuncAsc() func() (int16, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringInt16Map) IterFuncDesc() func() (int16, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringInt32Map struct {
	store   map[string]int32
	mapper  []string
	isWrite bool
}

func NewStringInt32Map() *StringInt32Map {
	om := &StringInt32Map{
		store:   make(map[string]int32),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringInt32MapAllocSpace(size int) *StringInt32Map {
	om := &StringInt32Map{
		store:   make(map[string]int32, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringInt32Map) Set(key string, value int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringInt32Map) Get(key string) (int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringInt32Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringInt32Map) Len() int {
	return len(om.store)
}

func (om *StringInt32Map) IterFunc() func() (int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringInt32Map) IterFuncAsc() func() (int32, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringInt32Map) IterFuncDesc() func() (int32, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringInt64Map struct {
	store   map[string]int64
	mapper  []string
	isWrite bool
}

func NewStringInt64Map() *StringInt64Map {
	om := &StringInt64Map{
		store:   make(map[string]int64),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringInt64MapAllocSpace(size int) *StringInt64Map {
	om := &StringInt64Map{
		store:   make(map[string]int64, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringInt64Map) Set(key string, value int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringInt64Map) Get(key string) (int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringInt64Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringInt64Map) Len() int {
	return len(om.store)
}

func (om *StringInt64Map) IterFunc() func() (int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringInt64Map) IterFuncAsc() func() (int64, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringInt64Map) IterFuncDesc() func() (int64, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringUint8Map struct {
	store   map[string]uint8
	mapper  []string
	isWrite bool
}

func NewStringUint8Map() *StringUint8Map {
	om := &StringUint8Map{
		store:   make(map[string]uint8),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringUint8MapAllocSpace(size int) *StringUint8Map {
	om := &StringUint8Map{
		store:   make(map[string]uint8, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringUint8Map) Set(key string, value uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringUint8Map) Get(key string) (uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringUint8Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringUint8Map) Len() int {
	return len(om.store)
}

func (om *StringUint8Map) IterFunc() func() (uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringUint8Map) IterFuncAsc() func() (uint8, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringUint8Map) IterFuncDesc() func() (uint8, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringUint16Map struct {
	store   map[string]uint16
	mapper  []string
	isWrite bool
}

func NewStringUint16Map() *StringUint16Map {
	om := &StringUint16Map{
		store:   make(map[string]uint16),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringUint16MapAllocSpace(size int) *StringUint16Map {
	om := &StringUint16Map{
		store:   make(map[string]uint16, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringUint16Map) Set(key string, value uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringUint16Map) Get(key string) (uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringUint16Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringUint16Map) Len() int {
	return len(om.store)
}

func (om *StringUint16Map) IterFunc() func() (uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringUint16Map) IterFuncAsc() func() (uint16, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringUint16Map) IterFuncDesc() func() (uint16, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringUint32Map struct {
	store   map[string]uint32
	mapper  []string
	isWrite bool
}

func NewStringUint32Map() *StringUint32Map {
	om := &StringUint32Map{
		store:   make(map[string]uint32),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringUint32MapAllocSpace(size int) *StringUint32Map {
	om := &StringUint32Map{
		store:   make(map[string]uint32, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringUint32Map) Set(key string, value uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringUint32Map) Get(key string) (uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringUint32Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringUint32Map) Len() int {
	return len(om.store)
}

func (om *StringUint32Map) IterFunc() func() (uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringUint32Map) IterFuncAsc() func() (uint32, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringUint32Map) IterFuncDesc() func() (uint32, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringUint64Map struct {
	store   map[string]uint64
	mapper  []string
	isWrite bool
}

func NewStringUint64Map() *StringUint64Map {
	om := &StringUint64Map{
		store:   make(map[string]uint64),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringUint64MapAllocSpace(size int) *StringUint64Map {
	om := &StringUint64Map{
		store:   make(map[string]uint64, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringUint64Map) Set(key string, value uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringUint64Map) Get(key string) (uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringUint64Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringUint64Map) Len() int {
	return len(om.store)
}

func (om *StringUint64Map) IterFunc() func() (uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringUint64Map) IterFuncAsc() func() (uint64, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringUint64Map) IterFuncDesc() func() (uint64, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringStringMap struct {
	store   map[string]string
	mapper  []string
	isWrite bool
}

func NewStringStringMap() *StringStringMap {
	om := &StringStringMap{
		store:   make(map[string]string),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringStringMapAllocSpace(size int) *StringStringMap {
	om := &StringStringMap{
		store:   make(map[string]string, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringStringMap) Set(key string, value string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringStringMap) Get(key string) (string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringStringMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringStringMap) Len() int {
	return len(om.store)
}

func (om *StringStringMap) IterFunc() func() (string, bool) {
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringStringMap) IterFuncAsc() func() (string, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringStringMap) IterFuncDesc() func() (string, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringFloat32Map struct {
	store   map[string]float32
	mapper  []string
	isWrite bool
}

func NewStringFloat32Map() *StringFloat32Map {
	om := &StringFloat32Map{
		store:   make(map[string]float32),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringFloat32MapAllocSpace(size int) *StringFloat32Map {
	om := &StringFloat32Map{
		store:   make(map[string]float32, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringFloat32Map) Set(key string, value float32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringFloat32Map) Get(key string) (float32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringFloat32Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringFloat32Map) Len() int {
	return len(om.store)
}

func (om *StringFloat32Map) IterFunc() func() (float32, bool) {
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringFloat32Map) IterFuncAsc() func() (float32, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringFloat32Map) IterFuncDesc() func() (float32, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringFloat64Map struct {
	store   map[string]float64
	mapper  []string
	isWrite bool
}

func NewStringFloat64Map() *StringFloat64Map {
	om := &StringFloat64Map{
		store:   make(map[string]float64),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringFloat64MapAllocSpace(size int) *StringFloat64Map {
	om := &StringFloat64Map{
		store:   make(map[string]float64, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringFloat64Map) Set(key string, value float64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringFloat64Map) Get(key string) (float64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringFloat64Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringFloat64Map) Len() int {
	return len(om.store)
}

func (om *StringFloat64Map) IterFunc() func() (float64, bool) {
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringFloat64Map) IterFuncAsc() func() (float64, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringFloat64Map) IterFuncDesc() func() (float64, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val float64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringInterfaceMap struct {
	store   map[string]interface{}
	mapper  []string
	isWrite bool
}

func NewStringInterfaceMap() *StringInterfaceMap {
	om := &StringInterfaceMap{
		store:   make(map[string]interface{}),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringInterfaceMapAllocSpace(size int) *StringInterfaceMap {
	om := &StringInterfaceMap{
		store:   make(map[string]interface{}, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringInterfaceMap) Set(key string, value interface{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringInterfaceMap) Get(key string) (interface{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringInterfaceMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringInterfaceMap) Len() int {
	return len(om.store)
}

func (om *StringInterfaceMap) IterFunc() func() (interface{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringInterfaceMap) IterFuncAsc() func() (interface{}, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringInterfaceMap) IterFuncDesc() func() (interface{}, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val interface{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringBoolMap struct {
	store   map[string]bool
	mapper  []string
	isWrite bool
}

func NewStringBoolMap() *StringBoolMap {
	om := &StringBoolMap{
		store:   make(map[string]bool),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringBoolMapAllocSpace(size int) *StringBoolMap {
	om := &StringBoolMap{
		store:   make(map[string]bool, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringBoolMap) Set(key string, value bool) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringBoolMap) Get(key string) (bool, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringBoolMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringBoolMap) Len() int {
	return len(om.store)
}

func (om *StringBoolMap) IterFunc() func() (bool, bool) {
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringBoolMap) IterFuncAsc() func() (bool, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringBoolMap) IterFuncDesc() func() (bool, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val bool, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringStructMap struct {
	store   map[string]struct{}
	mapper  []string
	isWrite bool
}

func NewStringStructMap() *StringStructMap {
	om := &StringStructMap{
		store:   make(map[string]struct{}),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringStructMapAllocSpace(size int) *StringStructMap {
	om := &StringStructMap{
		store:   make(map[string]struct{}, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringStructMap) Set(key string, value struct{}) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringStructMap) Get(key string) (struct{}, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringStructMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringStructMap) Len() int {
	return len(om.store)
}

func (om *StringStructMap) IterFunc() func() (struct{}, bool) {
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringStructMap) IterFuncAsc() func() (struct{}, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringStructMap) IterFuncDesc() func() (struct{}, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val struct{}, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceintMap struct {
	store   map[string][]int
	mapper  []string
	isWrite bool
}

func NewStringSliceintMap() *StringSliceintMap {
	om := &StringSliceintMap{
		store:   make(map[string][]int),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceintMapAllocSpace(size int) *StringSliceintMap {
	om := &StringSliceintMap{
		store:   make(map[string][]int, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceintMap) Set(key string, value []int) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceintMap) Get(key string) ([]int, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceintMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceintMap) Len() int {
	return len(om.store)
}

func (om *StringSliceintMap) IterFunc() func() ([]int, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceintMap) IterFuncAsc() func() ([]int, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceintMap) IterFuncDesc() func() ([]int, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceint8Map struct {
	store   map[string][]int8
	mapper  []string
	isWrite bool
}

func NewStringSliceint8Map() *StringSliceint8Map {
	om := &StringSliceint8Map{
		store:   make(map[string][]int8),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceint8MapAllocSpace(size int) *StringSliceint8Map {
	om := &StringSliceint8Map{
		store:   make(map[string][]int8, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceint8Map) Set(key string, value []int8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceint8Map) Get(key string) ([]int8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceint8Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceint8Map) Len() int {
	return len(om.store)
}

func (om *StringSliceint8Map) IterFunc() func() ([]int8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceint8Map) IterFuncAsc() func() ([]int8, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceint8Map) IterFuncDesc() func() ([]int8, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceint16Map struct {
	store   map[string][]int16
	mapper  []string
	isWrite bool
}

func NewStringSliceint16Map() *StringSliceint16Map {
	om := &StringSliceint16Map{
		store:   make(map[string][]int16),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceint16MapAllocSpace(size int) *StringSliceint16Map {
	om := &StringSliceint16Map{
		store:   make(map[string][]int16, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceint16Map) Set(key string, value []int16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceint16Map) Get(key string) ([]int16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceint16Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceint16Map) Len() int {
	return len(om.store)
}

func (om *StringSliceint16Map) IterFunc() func() ([]int16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceint16Map) IterFuncAsc() func() ([]int16, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceint16Map) IterFuncDesc() func() ([]int16, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceint32Map struct {
	store   map[string][]int32
	mapper  []string
	isWrite bool
}

func NewStringSliceint32Map() *StringSliceint32Map {
	om := &StringSliceint32Map{
		store:   make(map[string][]int32),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceint32MapAllocSpace(size int) *StringSliceint32Map {
	om := &StringSliceint32Map{
		store:   make(map[string][]int32, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceint32Map) Set(key string, value []int32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceint32Map) Get(key string) ([]int32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceint32Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceint32Map) Len() int {
	return len(om.store)
}

func (om *StringSliceint32Map) IterFunc() func() ([]int32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceint32Map) IterFuncAsc() func() ([]int32, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceint32Map) IterFuncDesc() func() ([]int32, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceint64Map struct {
	store   map[string][]int64
	mapper  []string
	isWrite bool
}

func NewStringSliceint64Map() *StringSliceint64Map {
	om := &StringSliceint64Map{
		store:   make(map[string][]int64),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceint64MapAllocSpace(size int) *StringSliceint64Map {
	om := &StringSliceint64Map{
		store:   make(map[string][]int64, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceint64Map) Set(key string, value []int64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceint64Map) Get(key string) ([]int64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceint64Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceint64Map) Len() int {
	return len(om.store)
}

func (om *StringSliceint64Map) IterFunc() func() ([]int64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceint64Map) IterFuncAsc() func() ([]int64, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceint64Map) IterFuncDesc() func() ([]int64, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []int64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceuint8Map struct {
	store   map[string][]uint8
	mapper  []string
	isWrite bool
}

func NewStringSliceuint8Map() *StringSliceuint8Map {
	om := &StringSliceuint8Map{
		store:   make(map[string][]uint8),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceuint8MapAllocSpace(size int) *StringSliceuint8Map {
	om := &StringSliceuint8Map{
		store:   make(map[string][]uint8, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceuint8Map) Set(key string, value []uint8) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceuint8Map) Get(key string) ([]uint8, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceuint8Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceuint8Map) Len() int {
	return len(om.store)
}

func (om *StringSliceuint8Map) IterFunc() func() ([]uint8, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint8Map) IterFuncAsc() func() ([]uint8, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint8Map) IterFuncDesc() func() ([]uint8, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint8, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceuint16Map struct {
	store   map[string][]uint16
	mapper  []string
	isWrite bool
}

func NewStringSliceuint16Map() *StringSliceuint16Map {
	om := &StringSliceuint16Map{
		store:   make(map[string][]uint16),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceuint16MapAllocSpace(size int) *StringSliceuint16Map {
	om := &StringSliceuint16Map{
		store:   make(map[string][]uint16, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceuint16Map) Set(key string, value []uint16) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceuint16Map) Get(key string) ([]uint16, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceuint16Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceuint16Map) Len() int {
	return len(om.store)
}

func (om *StringSliceuint16Map) IterFunc() func() ([]uint16, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint16Map) IterFuncAsc() func() ([]uint16, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint16Map) IterFuncDesc() func() ([]uint16, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint16, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceuint32Map struct {
	store   map[string][]uint32
	mapper  []string
	isWrite bool
}

func NewStringSliceuint32Map() *StringSliceuint32Map {
	om := &StringSliceuint32Map{
		store:   make(map[string][]uint32),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceuint32MapAllocSpace(size int) *StringSliceuint32Map {
	om := &StringSliceuint32Map{
		store:   make(map[string][]uint32, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceuint32Map) Set(key string, value []uint32) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceuint32Map) Get(key string) ([]uint32, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceuint32Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceuint32Map) Len() int {
	return len(om.store)
}

func (om *StringSliceuint32Map) IterFunc() func() ([]uint32, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint32Map) IterFuncAsc() func() ([]uint32, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint32Map) IterFuncDesc() func() ([]uint32, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint32, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSliceuint64Map struct {
	store   map[string][]uint64
	mapper  []string
	isWrite bool
}

func NewStringSliceuint64Map() *StringSliceuint64Map {
	om := &StringSliceuint64Map{
		store:   make(map[string][]uint64),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSliceuint64MapAllocSpace(size int) *StringSliceuint64Map {
	om := &StringSliceuint64Map{
		store:   make(map[string][]uint64, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSliceuint64Map) Set(key string, value []uint64) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSliceuint64Map) Get(key string) ([]uint64, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSliceuint64Map) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSliceuint64Map) Len() int {
	return len(om.store)
}

func (om *StringSliceuint64Map) IterFunc() func() ([]uint64, bool) {
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint64Map) IterFuncAsc() func() ([]uint64, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSliceuint64Map) IterFuncDesc() func() ([]uint64, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []uint64, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}

type StringSlicestringMap struct {
	store   map[string][]string
	mapper  []string
	isWrite bool
}

func NewStringSlicestringMap() *StringSlicestringMap {
	om := &StringSlicestringMap{
		store:   make(map[string][]string),
		mapper:  make([]string, 0),
		isWrite: false,
	}
	return om
}

func NewStringSlicestringMapAllocSpace(size int) *StringSlicestringMap {
	om := &StringSlicestringMap{
		store:   make(map[string][]string, size),
		mapper:  make([]string, 0, size),
		isWrite: false,
	}
	return om
}

func (om *StringSlicestringMap) Set(key string, value []string) {
	if _, ok := om.store[key]; !ok {
		om.mapper = append(om.mapper, key)
	}
	om.store[key] = value
	om.isWrite = true
}

func (om *StringSlicestringMap) Get(key string) ([]string, bool) {
	val, ok := om.store[key]
	return val, ok
}

func (om *StringSlicestringMap) Delete(key string) {
	if _, ok := om.store[key]; ok {
		delete(om.store, key)
		om.isWrite = true

		idx := -1
		for i, k := range om.mapper {
			if k == key {
				idx = i
				break
			}
		}
		if idx >= 0 {
			om.mapper = append(om.mapper[:idx], om.mapper[idx+1:]...)
		}
	}
}

func (om *StringSlicestringMap) Len() int {
	return len(om.store)
}

func (om *StringSlicestringMap) IterFunc() func() ([]string, bool) {
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(om.mapper) || om.isWrite {
			return
		}
		val, ok = om.store[om.mapper[index]]
		index++
		return
	}
}

func (om *StringSlicestringMap) IterFuncAsc() func() ([]string, bool) {
	ascMapper := make([]string, len(om.mapper))
	copy(ascMapper, om.mapper)
	sorts.Strings(ascMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(ascMapper) || om.isWrite {
			return
		}
		val, ok = om.store[ascMapper[index]]
		index++
		return
	}
}

func (om *StringSlicestringMap) IterFuncDesc() func() ([]string, bool) {
	dscMapper := make([]string, len(om.mapper))
	copy(dscMapper, om.mapper)
	sorts.StringsRev(dscMapper)
	index := 0
	om.isWrite = false
	return func() (val []string, ok bool) {
		if index >= len(dscMapper) || om.isWrite {
			return
		}
		val, ok = om.store[dscMapper[index]]
		index++
		return
	}
}
