package meta

import (
	"reflect"
	"testing"
)

type y struct {
	A int
	B int
}

func (ø *y) Sync() { ø.B = ø.A }

func TestCallMethod(t *testing.T) {
	y_ := &y{2, 3}
	CallMethod(y_, "Sync")

	if y_.B != y_.A {
		t.Errorf("error in CallMethod, expected: %d, got %d", y_.A, y_.B)
	}
}

func TestReplace(t *testing.T) {
	a := "hi"
	Replace(&a, "ho")
	if a != "ho" {
		t.Errorf("error in Replace, expected: %#v, got %#v", "ho", a)
	}
}

type ii float32

func TestConvert(t *testing.T) {
	cv := Convert(42, reflect.TypeOf(ii(0.0))).(ii)

	if cv != ii(42.0) {
		t.Errorf("error in Convert, expected: %#v, got %#v", ii(42.0), cv)
	}
}

func TestAssoc(t *testing.T) {
	y1 := &y{A: 5}
	y2 := &y{}
	Assoc(y1, &y2)

	if y2.A != 5 {
		t.Errorf("error in Assoc, did not set target")
	}

	y1.A = 6

	if y2.A != 6 {
		t.Errorf("error in Assoc, target not associated with src")
	}
}

func TestNewPtr(t *testing.T) {
	y1 := NewPtr(reflect.TypeOf(y{})).(**y)

	if reflect.TypeOf(**y1) != reflect.TypeOf(y{}) {
		t.Errorf("NewPtr returns not pointer to a pointer, is of type %T and should be **%T", y1, y{})
	}
}

func TestReferenceTo(t *testing.T) {
	y1 := y{A: 5}
	y2 := ReferenceTo(reflect.ValueOf(y1)).(*y)

	if y2.A != 5 {
		t.Errorf("error in ReferenceTo")
	}
}

/*
func ReferenceTo(val reflect.Value) interface{} {
	ref := reflect.New(val.Type())
	ref.Elem().Set(val)
	return ref.Interface()
}
*/
