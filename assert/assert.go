package assert

import (
	"math"
	"reflect"
)

func IsNil(someThing interface{}) {
	if someThing == nil {
		return
	}
	if !reflect.ValueOf(someThing).IsNil() {
		panic(newAssertPanic("%+v is not nil", someThing))
	}
}

func IsNotNil(someThing interface{}) {
	if someThing == nil {
		panic(newAssertPanic("%+v is nil", someThing))
	}
	if reflect.ValueOf(someThing).IsNil() {
		panic(newAssertPanic("%+v is nil", someThing))
	}
}

func IsTrue(b bool) {
	if !b {
		panic(newAssertPanic("assert fail, not true"))
	}
}

func IsFalse(b bool) {
	if b {
		panic(newAssertPanic("assert fail, not false"))
	}
}

func ApproximateEqualFloat32(a, b float32) {
	if math.Abs(float64(a-b)) > 1e-5 {
		panic(newAssertPanic("assert fail, %v and %v not approximate equal", a, b))
	}
}

func ApproximateEqualFloat64(a, b float64) {
	if math.Abs(a-b) > 1e-5 {
		panic(newAssertPanic("assert fail, %v and %v not approximate equal", a, b))
	}
}

func IsNilWithMsg(someThing interface{}, msg string) {
	if someThing == nil {
		return
	}
	if !reflect.ValueOf(someThing).IsNil() {
		panic(newAssertPanic(msg))
	}
}

func IsNotNilWithMsg(someThing interface{}, msg string) {
	if someThing == nil {
		panic(newAssertPanic(msg))
	}
	if reflect.ValueOf(someThing).IsNil() {
		panic(newAssertPanic(msg))
	}
}

func IsTrueWithMsg(b bool, msg string) {
	if !b {
		panic(newAssertPanic(msg))
	}
}

func IsFalseWithMsg(b bool, msg string) {
	if b {
		panic(newAssertPanic(msg))
	}
}

func ApproximateEqualFloat32WithMsg(a, b float32, msg string) {
	if math.Abs(float64(a-b)) > 1e-5 {
		panic(newAssertPanic(msg))
	}
}

func ApproximateEqualFloat64WithMsg(a, b float64, msg string) {
	if math.Abs(a-b) > 1e-5 {
		panic(newAssertPanic(msg))
	}
}
