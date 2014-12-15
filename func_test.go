package meta

import (
	"bytes"
	"fmt"
	"reflect"
	// "fmt"
	"testing"
)

func Add2(i *int) { *i = *i + 2 }

func Add2a(i int) int { return i + 2 }

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

func TestCallAndReturn(t *testing.T) {
	r := Func.CallAndReturn(Add2a, 4)
	if len(r) != 1 {
		t.Errorf("CallAndReturn should have 1 return, but has %d", len(r))
	}

	i, ok := r[0].(int)
	if !ok {
		t.Errorf("CallAndReturn should return int, but returns %T", r)
	}

	if i != 6 {
		t.Errorf("CallAndReturn did not work, got %d, expected %d", i, 6)
	}
}

func TestSliceArg(t *testing.T) {
	fn := func(a ...interface{}) string {
		var b bytes.Buffer

		for _, aa := range a {
			fmt.Fprintf(&b, "%T-", aa)
		}

		return b.String()
	}

	sl := []int{2, 3, 4}

	r := fn(Func.SliceArg(sl)...)
	exp := "int-int-int-"
	if r != exp {
		t.Errorf("SliceArg not working, expected: %s, got: %s\n", exp, r)
	}
}
