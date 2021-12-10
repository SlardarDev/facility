package set

type IntSet map[int]struct{}

func (s IntSet) Exist(key int) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewIntSet(keys ...int) IntSet {
	s := IntSet{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s IntSet) Add(key int) {
	s[key] = struct{}{}
}

func (s IntSet) AddAll(keys ...int) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s IntSet) Remove(key int) {
	delete(s, key)
}

func (s IntSet) Contains(key int) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s IntSet) ToList() []int {
	var ret []int

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s IntSet) Intersect(s2 IntSet) IntSet {
	newSet := IntSet{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s IntSet) Union(s2 IntSet) IntSet {
	newSet := IntSet{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s IntSet) Minus(s2 IntSet) IntSet {
	newSet := IntSet{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s IntSet) MinusInPlace(s2 IntSet) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s IntSet) IsEmpty() bool {
	return len(s) == 0
}

func (s IntSet) Equal(s2 IntSet) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s IntSet) Size() int {
	return len(s)
}

type Int8Set map[int8]struct{}

func (s Int8Set) Exist(key int8) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewInt8Set(keys ...int8) Int8Set {
	s := Int8Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Int8Set) Add(key int8) {
	s[key] = struct{}{}
}

func (s Int8Set) AddAll(keys ...int8) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Int8Set) Remove(key int8) {
	delete(s, key)
}

func (s Int8Set) Contains(key int8) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Int8Set) ToList() []int8 {
	var ret []int8

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Int8Set) Intersect(s2 Int8Set) Int8Set {
	newSet := Int8Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int8Set) Union(s2 Int8Set) Int8Set {
	newSet := Int8Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Int8Set) Minus(s2 Int8Set) Int8Set {
	newSet := Int8Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int8Set) MinusInPlace(s2 Int8Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Int8Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Int8Set) Equal(s2 Int8Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Int8Set) Size() int {
	return len(s)
}

type Int16Set map[int16]struct{}

func (s Int16Set) Exist(key int16) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewInt16Set(keys ...int16) Int16Set {
	s := Int16Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Int16Set) Add(key int16) {
	s[key] = struct{}{}
}

func (s Int16Set) AddAll(keys ...int16) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Int16Set) Remove(key int16) {
	delete(s, key)
}

func (s Int16Set) Contains(key int16) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Int16Set) ToList() []int16 {
	var ret []int16

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Int16Set) Intersect(s2 Int16Set) Int16Set {
	newSet := Int16Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int16Set) Union(s2 Int16Set) Int16Set {
	newSet := Int16Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Int16Set) Minus(s2 Int16Set) Int16Set {
	newSet := Int16Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int16Set) MinusInPlace(s2 Int16Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Int16Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Int16Set) Equal(s2 Int16Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Int16Set) Size() int {
	return len(s)
}

type Int32Set map[int32]struct{}

func (s Int32Set) Exist(key int32) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewInt32Set(keys ...int32) Int32Set {
	s := Int32Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Int32Set) Add(key int32) {
	s[key] = struct{}{}
}

func (s Int32Set) AddAll(keys ...int32) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Int32Set) Remove(key int32) {
	delete(s, key)
}

func (s Int32Set) Contains(key int32) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Int32Set) ToList() []int32 {
	var ret []int32

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Int32Set) Intersect(s2 Int32Set) Int32Set {
	newSet := Int32Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int32Set) Union(s2 Int32Set) Int32Set {
	newSet := Int32Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Int32Set) Minus(s2 Int32Set) Int32Set {
	newSet := Int32Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int32Set) MinusInPlace(s2 Int32Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Int32Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Int32Set) Equal(s2 Int32Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Int32Set) Size() int {
	return len(s)
}

type Int64Set map[int64]struct{}

func (s Int64Set) Exist(key int64) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewInt64Set(keys ...int64) Int64Set {
	s := Int64Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Int64Set) Add(key int64) {
	s[key] = struct{}{}
}

func (s Int64Set) AddAll(keys ...int64) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Int64Set) Remove(key int64) {
	delete(s, key)
}

func (s Int64Set) Contains(key int64) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Int64Set) ToList() []int64 {
	var ret []int64

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Int64Set) Intersect(s2 Int64Set) Int64Set {
	newSet := Int64Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int64Set) Union(s2 Int64Set) Int64Set {
	newSet := Int64Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Int64Set) Minus(s2 Int64Set) Int64Set {
	newSet := Int64Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Int64Set) MinusInPlace(s2 Int64Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Int64Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Int64Set) Equal(s2 Int64Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Int64Set) Size() int {
	return len(s)
}

type UintSet map[uint]struct{}

func (s UintSet) Exist(key uint) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewUintSet(keys ...uint) UintSet {
	s := UintSet{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s UintSet) Add(key uint) {
	s[key] = struct{}{}
}

func (s UintSet) AddAll(keys ...uint) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s UintSet) Remove(key uint) {
	delete(s, key)
}

func (s UintSet) Contains(key uint) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s UintSet) ToList() []uint {
	var ret []uint

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s UintSet) Intersect(s2 UintSet) UintSet {
	newSet := UintSet{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s UintSet) Union(s2 UintSet) UintSet {
	newSet := UintSet{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s UintSet) Minus(s2 UintSet) UintSet {
	newSet := UintSet{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s UintSet) MinusInPlace(s2 UintSet) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s UintSet) IsEmpty() bool {
	return len(s) == 0
}

func (s UintSet) Equal(s2 UintSet) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s UintSet) Size() int {
	return len(s)
}

type Uint8Set map[uint8]struct{}

func (s Uint8Set) Exist(key uint8) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewUint8Set(keys ...uint8) Uint8Set {
	s := Uint8Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Uint8Set) Add(key uint8) {
	s[key] = struct{}{}
}

func (s Uint8Set) AddAll(keys ...uint8) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Uint8Set) Remove(key uint8) {
	delete(s, key)
}

func (s Uint8Set) Contains(key uint8) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Uint8Set) ToList() []uint8 {
	var ret []uint8

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Uint8Set) Intersect(s2 Uint8Set) Uint8Set {
	newSet := Uint8Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint8Set) Union(s2 Uint8Set) Uint8Set {
	newSet := Uint8Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Uint8Set) Minus(s2 Uint8Set) Uint8Set {
	newSet := Uint8Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint8Set) MinusInPlace(s2 Uint8Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Uint8Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Uint8Set) Equal(s2 Uint8Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Uint8Set) Size() int {
	return len(s)
}

type Uint16Set map[uint16]struct{}

func (s Uint16Set) Exist(key uint16) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewUint16Set(keys ...uint16) Uint16Set {
	s := Uint16Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Uint16Set) Add(key uint16) {
	s[key] = struct{}{}
}

func (s Uint16Set) AddAll(keys ...uint16) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Uint16Set) Remove(key uint16) {
	delete(s, key)
}

func (s Uint16Set) Contains(key uint16) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Uint16Set) ToList() []uint16 {
	var ret []uint16

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Uint16Set) Intersect(s2 Uint16Set) Uint16Set {
	newSet := Uint16Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint16Set) Union(s2 Uint16Set) Uint16Set {
	newSet := Uint16Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Uint16Set) Minus(s2 Uint16Set) Uint16Set {
	newSet := Uint16Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint16Set) MinusInPlace(s2 Uint16Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Uint16Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Uint16Set) Equal(s2 Uint16Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Uint16Set) Size() int {
	return len(s)
}

type Uint32Set map[uint32]struct{}

func (s Uint32Set) Exist(key uint32) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewUint32Set(keys ...uint32) Uint32Set {
	s := Uint32Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Uint32Set) Add(key uint32) {
	s[key] = struct{}{}
}

func (s Uint32Set) AddAll(keys ...uint32) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Uint32Set) Remove(key uint32) {
	delete(s, key)
}

func (s Uint32Set) Contains(key uint32) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Uint32Set) ToList() []uint32 {
	var ret []uint32

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Uint32Set) Intersect(s2 Uint32Set) Uint32Set {
	newSet := Uint32Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint32Set) Union(s2 Uint32Set) Uint32Set {
	newSet := Uint32Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Uint32Set) Minus(s2 Uint32Set) Uint32Set {
	newSet := Uint32Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint32Set) MinusInPlace(s2 Uint32Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Uint32Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Uint32Set) Equal(s2 Uint32Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Uint32Set) Size() int {
	return len(s)
}

type Uint64Set map[uint64]struct{}

func (s Uint64Set) Exist(key uint64) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewUint64Set(keys ...uint64) Uint64Set {
	s := Uint64Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Uint64Set) Add(key uint64) {
	s[key] = struct{}{}
}

func (s Uint64Set) AddAll(keys ...uint64) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Uint64Set) Remove(key uint64) {
	delete(s, key)
}

func (s Uint64Set) Contains(key uint64) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Uint64Set) ToList() []uint64 {
	var ret []uint64

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Uint64Set) Intersect(s2 Uint64Set) Uint64Set {
	newSet := Uint64Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint64Set) Union(s2 Uint64Set) Uint64Set {
	newSet := Uint64Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Uint64Set) Minus(s2 Uint64Set) Uint64Set {
	newSet := Uint64Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Uint64Set) MinusInPlace(s2 Uint64Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Uint64Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Uint64Set) Equal(s2 Uint64Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Uint64Set) Size() int {
	return len(s)
}

type StringSet map[string]struct{}

func (s StringSet) Exist(key string) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewStringSet(keys ...string) StringSet {
	s := StringSet{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s StringSet) Add(key string) {
	s[key] = struct{}{}
}

func (s StringSet) AddAll(keys ...string) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s StringSet) Remove(key string) {
	delete(s, key)
}

func (s StringSet) Contains(key string) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s StringSet) ToList() []string {
	var ret []string

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s StringSet) Intersect(s2 StringSet) StringSet {
	newSet := StringSet{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s StringSet) Union(s2 StringSet) StringSet {
	newSet := StringSet{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s StringSet) Minus(s2 StringSet) StringSet {
	newSet := StringSet{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s StringSet) MinusInPlace(s2 StringSet) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s StringSet) IsEmpty() bool {
	return len(s) == 0
}

func (s StringSet) Equal(s2 StringSet) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s StringSet) Size() int {
	return len(s)
}

type Float64Set map[float64]struct{}

func (s Float64Set) Exist(key float64) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewFloat64Set(keys ...float64) Float64Set {
	s := Float64Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Float64Set) Add(key float64) {
	s[key] = struct{}{}
}

func (s Float64Set) AddAll(keys ...float64) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Float64Set) Remove(key float64) {
	delete(s, key)
}

func (s Float64Set) Contains(key float64) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Float64Set) ToList() []float64 {
	var ret []float64

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Float64Set) Intersect(s2 Float64Set) Float64Set {
	newSet := Float64Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Float64Set) Union(s2 Float64Set) Float64Set {
	newSet := Float64Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Float64Set) Minus(s2 Float64Set) Float64Set {
	newSet := Float64Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Float64Set) MinusInPlace(s2 Float64Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Float64Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Float64Set) Equal(s2 Float64Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Float64Set) Size() int {
	return len(s)
}

type Float32Set map[float32]struct{}

func (s Float32Set) Exist(key float32) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewFloat32Set(keys ...float32) Float32Set {
	s := Float32Set{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s Float32Set) Add(key float32) {
	s[key] = struct{}{}
}

func (s Float32Set) AddAll(keys ...float32) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s Float32Set) Remove(key float32) {
	delete(s, key)
}

func (s Float32Set) Contains(key float32) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s Float32Set) ToList() []float32 {
	var ret []float32

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s Float32Set) Intersect(s2 Float32Set) Float32Set {
	newSet := Float32Set{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Float32Set) Union(s2 Float32Set) Float32Set {
	newSet := Float32Set{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s Float32Set) Minus(s2 Float32Set) Float32Set {
	newSet := Float32Set{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s Float32Set) MinusInPlace(s2 Float32Set) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s Float32Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Float32Set) Equal(s2 Float32Set) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s Float32Set) Size() int {
	return len(s)
}

type BoolSet map[bool]struct{}

func (s BoolSet) Exist(key bool) bool {
	if s == nil {
		return false
	}
	_, ok := s[key]
	return ok
}

func NewBoolSet(keys ...bool) BoolSet {
	s := BoolSet{}
	for _, k := range keys {
		s[k] = struct{}{}
	}
	return s
}

func (s BoolSet) Add(key bool) {
	s[key] = struct{}{}
}

func (s BoolSet) AddAll(keys ...bool) {
	if len(keys) > 0 {
		for _, key := range keys {
			s[key] = struct{}{}
		}
	}

}

func (s BoolSet) Remove(key bool) {
	delete(s, key)
}

func (s BoolSet) Contains(key bool) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s BoolSet) ToList() []bool {
	var ret []bool

	for k := range s {
		ret = append(ret, k)
	}
	return ret
}

func (s BoolSet) Intersect(s2 BoolSet) BoolSet {
	newSet := BoolSet{}
	for ele := range s {
		if s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s BoolSet) Union(s2 BoolSet) BoolSet {
	newSet := BoolSet{}
	for ele := range s {
		newSet.Add(ele)
	}
	for ele := range s2 {
		newSet.Add(ele)
	}
	return newSet
}

func (s BoolSet) Minus(s2 BoolSet) BoolSet {
	newSet := BoolSet{}
	for ele := range s {
		if !s2.Contains(ele) {
			newSet.Add(ele)
		}
	}
	return newSet
}

func (s BoolSet) MinusInPlace(s2 BoolSet) {
	for ele := range s2 {
		s.Remove(ele)
	}
}

func (s BoolSet) IsEmpty() bool {
	return len(s) == 0
}

func (s BoolSet) Equal(s2 BoolSet) bool {
	if len(s) != len(s2) {
		return false
	}

	for ele := range s {
		if !s2.Contains(ele) {
			return false
		}
	}

	return true

}

func (s BoolSet) Size() int {
	return len(s)
}
