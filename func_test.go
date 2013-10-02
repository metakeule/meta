package meta

import (
	"reflect"
	// "fmt"
	"testing"
)

func Add2(i *int) { *i = *i + 2 }

func TestCall(t *testing.T) {
	var j = 3
	Add2(&j)

	if j != 5 {
		t.Errorf("Add2 did not work, got %d, expected %d", j, 5)
	}

	var k = 4
	Func.Call(Add2, &k)
	if k != 6 {
		t.Errorf("Call did not work, got %d, expected %d", k, 6)
	}
}

func TestFuncReplace(t *testing.T) {
	gen := func(in []reflect.Value) (out []reflect.Value) {
		v := in[0].Elem().Int() + 5
		in[0].Elem().Set(reflect.ValueOf(v))
		return []reflect.Value{}
	}
	a := func(i *int64) { *i = *i + 3 }
	Func.Replace(&a, gen)

	var j = int64(5)
	a(&j)
	if j != 10 {
		t.Errorf("Replace did not work, got %d, expected %d", j, 10)
	}
}
