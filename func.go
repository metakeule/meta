package meta

import (
	"fmt"
	"reflect"
)

// replaces the function with a (generic) function of the same type
// stolen from the example of reflect.MakeFunc
// specific and generic must have the same number of input and output parameters
func (ø func_) Replace(specific interface{}, generic func([]reflect.Value) []reflect.Value) {
	if !ø.IsPointerTo(reflect.Func, specific) {
		panic(fmt.Sprintf("%s must be a pointer to a function", Inspect(specific)))
	}

	fn := reflect.ValueOf(specific).Elem()
	v := reflect.MakeFunc(fn.Type(), generic)
	fn.Set(v)
}

// calls function ø with vals, but doesn't return anything
func (ø func_) Call(f interface{}, vals ...interface{}) {
	params := []reflect.Value{}
	for i := range vals {
		if vals[i] != nil {
			params = append(params, reflect.ValueOf(vals[i]))
		}
	}
	reflect.ValueOf(f).Call(params)
}
