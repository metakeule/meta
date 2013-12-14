package meta

import (
	"fmt"
	"reflect"
)

func Panicf(s string, i ...interface{}) {
	panic(fmt.Sprintf(s, i...))
}

type checker interface {
	Check(interface{}) bool
	String() string
}

func Is(c checker, i interface{}) error {
	if !c.Check(i) {
		return fmt.Errorf("%#v is not a %s", i, c)
	}
	return nil
}

func MustBe(c checker, i interface{}) {
	if !c.Check(i) {
		panic(fmt.Sprintf("%#v is not a %s", i, c))
	}
}

func Inspect(i interface{}) (s string) {
	if reflect.TypeOf(i).Kind().String() == "float64" || reflect.TypeOf(i).Kind().String() == "float32" {
		s = fmt.Sprintf("%f (%s)", i, reflect.TypeOf(i))
	} else {
		s = fmt.Sprintf("%#v (%s)", i, reflect.TypeOf(i))
	}
	return
}

// returns a reference to a new empty value based on Type
func New(t reflect.Type) (ptr interface{}) {
	return reflect.New(t).Interface()
}

// returns a reference to a new empty value based on Type of given i
func NewByValue(i interface{}) (ptr interface{}) {
	return reflect.New(reflect.TypeOf(i)).Interface()
}

// returns a reference to a new reference to a new empty value based on Type
func NewPtr(ty reflect.Type) interface{} {
	val := reflect.New(ty)
	ref := reflect.New(val.Type())
	ref.Elem().Set(val)
	return ref.Interface()
}

// returns the underlying type of a reference
func DeReference(ptr interface{}) interface{} {
	return reflect.ValueOf(ptr).Elem().Interface()
}

// Assoc associates targetPtrPtr with srcPtr so that
// targetPtrPtr is a pointer to srcPtr and
// targetPtr and srcPtr are pointing to the same address
func Assoc(srcPtr, targetPtrPtr interface{}) {
	reflect.ValueOf(targetPtrPtr).Elem().Set(reflect.ValueOf(srcPtr))
}

// allows you to compare, like if s == Defaults.String

// extend it to your pleasure
var Defaults = map[reflect.Type]interface{}{
	reflect.TypeOf(""):           "",
	reflect.TypeOf(int(0)):       int(0),
	reflect.TypeOf(int32(0)):     int32(0),
	reflect.TypeOf(int64(0)):     int64(0),
	reflect.TypeOf(true):         false,
	reflect.TypeOf(float32(0.0)): float32(0.0),
	reflect.TypeOf(float64(0.0)): float64(0.0),
}

func IsDefault(i interface{}) bool {
	if d := Defaults[reflect.TypeOf(i)]; d != nil {
		// naive, does it work?? check!
		if d == i {
			return true
		}
	}
	return false
}

func P(i interface{}) {
	fmt.Println(Inspect(i))
	return
}

// calls a method of ø with vals, but doesn't return anything
func CallMethod(ø interface{}, meth string, vals ...interface{}) {
	m := reflect.ValueOf(ø).MethodByName(meth)
	if !m.IsValid() {
		panic("can't find method " + meth + " of " + Inspect(ø))
	}
	params := []reflect.Value{}
	for i := range vals {
		if vals[i] != nil {
			params = append(params, reflect.ValueOf(vals[i]))
		}
	}
	m.Call(params)
}

//	does not work, maybe not possible
/*
func ReplaceMethod(ø interface{}, meth string, generic func([]reflect.Value) []reflect.Value) {
	m := reflect.ValueOf(ø).MethodByName(meth)
	if !m.IsValid() {
		panic("can't find method " + meth + " of " + Inspect(ø))
	}
	v := reflect.MakeFunc(m.Type(), generic)
	m.Set(v)
}
*/

// replaces a value
func Replace(ø interface{}, val interface{}) {
	p := reflect.ValueOf(ø)
	if p.Elem().CanSet() {
		p.Elem().Set(reflect.ValueOf(val))
	} else {
		panic("can't set " + Inspect(ø))
	}
}

// converts a value to type
// only works with go1.1
func Convert(i interface{}, t reflect.Type) (r interface{}) {
	return reflect.ValueOf(i).Convert(t).Interface()
}
