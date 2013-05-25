package meta

import (
	"reflect"
)

type channel_ struct{ meta }
type float_ struct{ meta }
type bool_ struct{ meta }
type complex_ struct{ meta }
type slice_ struct{ meta }
type array_ struct{ meta }
type string_ struct{ meta }
type struct_ struct{ meta }
type int_ struct{ meta }
type uint_ struct{ meta }
type map_ struct{ meta }
type func_ struct{ meta }
type pointer_ struct{ meta }

//type interface_ struct{ meta }
type nil_ struct{ meta }

//var Interface = interface_{meta{"interface"}}

var Channel = channel_{meta{"channel", func(ø meta, i interface{}) bool { return ø.Is(reflect.Chan, i) }}}
var Slice = slice_{meta{"slice", func(ø meta, i interface{}) bool { return ø.Is(reflect.Slice, i) }}}
var Array = array_{meta{"array", func(ø meta, i interface{}) bool { return ø.Is(reflect.Array, i) }}}
var String = string_{meta{"string", func(ø meta, i interface{}) bool { return ø.Is(reflect.String, i) }}}
var Struct = struct_{meta{"struct", func(ø meta, i interface{}) bool { return ø.Is(reflect.Struct, i) }}}
var Map = map_{meta{"map", func(ø meta, i interface{}) bool { return ø.Is(reflect.Map, i) }}}
var Func = func_{meta{"func", func(ø meta, i interface{}) bool { return ø.Is(reflect.Func, i) }}}
var Nil = nil_{meta{"nil", func(ø meta, i interface{}) bool { return i == nil }}}
var Float = float_{meta{"float", func(ø meta, i interface{}) bool {
	return ø.Is(reflect.Float64, i) ||
		ø.Is(reflect.Float32, i)
}}}
var Bool = bool_{meta{"bool", func(ø meta, i interface{}) bool { return ø.Is(reflect.Bool, i) }}}
var Complex = complex_{meta{"complex", func(ø meta, i interface{}) bool {
	return ø.Is(reflect.Complex128, i) ||
		ø.Is(reflect.Complex64, i)
}}}

var Int = int_{meta{"int", func(ø meta, i interface{}) bool {
	return ø.Is(reflect.Int, i) ||
		ø.Is(reflect.Int8, i) ||
		ø.Is(reflect.Int16, i) ||
		ø.Is(reflect.Int32, i) ||
		ø.Is(reflect.Int64, i)
}}}
var Uint = uint_{meta{"uint", func(ø meta, i interface{}) bool {
	return ø.Is(reflect.Uint, i) ||
		ø.Is(reflect.Uint8, i) ||
		ø.Is(reflect.Uint16, i) ||
		ø.Is(reflect.Uint32, i) ||
		ø.Is(reflect.Uint64, i)
}}}

var Pointer = pointer_{meta{"pointer", func(ø meta, i interface{}) bool {
	return ø.Is(reflect.Ptr, i) ||
		ø.Is(reflect.Uintptr, i) ||
		ø.Is(reflect.UnsafePointer, i)
}}}
