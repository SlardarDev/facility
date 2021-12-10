package slice

import "reflect"

func fillEmptySliceForValue(o reflect.Value) {
	if o.IsNil() {
		return
	}
	ele := o.Elem()
	if ele.Kind() != reflect.Struct {
		return
	}
	numField := ele.NumField()
	for i := 0; i < numField; i++ {
		field := ele.Field(i)
		fieldKind := field.Kind()
		if fieldKind == reflect.Struct {
			if field.CanAddr() {
				fillEmptySliceForValue(field.Addr())
			}
		}
		if fieldKind == reflect.Ptr {
			fillEmptySliceForValue(field)
		}
		if fieldKind == reflect.Slice {
			if field.IsNil() {
				field.Set(reflect.MakeSlice(field.Type(), 0, 0))
			}
		}
	}
}

/*
o必须是指针
*/
func FillEmptySlice(o interface{}) {
	if o == nil {
		return
	}
	typ := reflect.TypeOf(o)
	if typ.Kind() != reflect.Ptr {
		return
	}
	v := reflect.ValueOf(o)
	if v.IsNil() {
		return
	}
	fillEmptySliceForValue(v)
}
