package meta

import (
	"fmt"
	"reflect"
)

// replaces the function with a (generic) function of the same type
// stolen from the example of reflect.MakeFunc
// specific and generic must have the same number of input and output parameters
//only works with go1.1+
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

// calls function ø with vals and returns the returned values as a
// slice of interfaces
func (ø func_) CallAndReturn(f interface{}, vals ...interface{}) []interface{} {
	params := []reflect.Value{}
	for i := range vals {
		if vals[i] != nil {
			params = append(params, reflect.ValueOf(vals[i]))
		}
	}
	res1 := reflect.ValueOf(f).Call(params)
	res2 := make([]interface{}, len(res1))
	for i := 0; i < len(res1); i++ {
		res2[i] = res1[i].Interface()
	}
	return res2
}
