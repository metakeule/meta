package meta

import (
	// "reflect"
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

type ii int

/*
func TestConvert(t *testing.T) {
	cv := Convert(42, reflect.TypeOf(ii(0))).(ii)

	if cv != ii(42) {
		t.Errorf("error in Convert, expected: %#v, got %#v", ii(42), cv)
	}
}
*/
