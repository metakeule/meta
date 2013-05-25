package meta

import (
	//"fmt"
	"reflect"
)

/*
	This is a collection of Hacks around go's reflect package

	It is organized by "generic" kinds, like Int, String, Struct, Map etc
*/

var Meta = meta{}

type meta struct {
	name    string
	checker func(meta, interface{}) bool
}

func (ø meta) String() string { return ø.name }

func (ø meta) IsPointerTo(k reflect.Kind, i interface{}) bool {
	if ø.Is(reflect.Ptr, i) && reflect.TypeOf(i).Elem().Kind() == k {
		return true
	}
	return false
}

func (ø meta) Check(i interface{}) bool { return ø.checker(ø, i) }

func (ø meta) Is(k reflect.Kind, i interface{}) bool {
	//fmt.Println(ø.Kind(i))
	return ø.Kind(i) == k
}

func (ø meta) HasType(t reflect.Type, i interface{}) bool {
	return reflect.TypeOf(i) == t
}

func (ø meta) Kind(i interface{}) reflect.Kind {
	return reflect.TypeOf(i).Kind()
}

func (ø meta) FinalType(i interface{}) reflect.Type {
	if ø.Is(reflect.Ptr, i) {
		return reflect.TypeOf(i).Elem()
	}
	return reflect.TypeOf(i)
}

func (ø meta) FinalValue(i interface{}) reflect.Value {
	if ø.Is(reflect.Ptr, i) {
		return reflect.ValueOf(i).Elem()
	}
	return reflect.ValueOf(i)
}
