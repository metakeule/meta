package meta

import (
	"testing"
)

type x struct {
	A string
	B int
	C bool
}

func TestToStruct(t *testing.T) {
	m1 := map[string]interface{}{
		"A": "hiho",
		"B": 42,
	}

	m2 := map[string]bool{
		"C": true,
	}

	var x_ = &x{}
	Map.ToStruct(m1, x_)
	Map.ToStruct(m2, x_)

	if x_.A != "hiho" {
		t.Errorf("error in ToStruct(A): got %#v, expected %#v\n", x_.A, "hiho")
	}
	if x_.B != 42 {
		t.Errorf("error in ToStruct(B): got %#v, expected %#v\n", x_.B, 42)
	}
	if !x_.C {
		t.Errorf("error in ToStruct(C): got %#v, expected %#v\n", x_.C, true)
	}
}
