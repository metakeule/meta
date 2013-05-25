package meta

import (
	//"fmt"
	"testing"
)

type s struct{}

func TestCheck(t *testing.T) {
	//var i interface{}
	if !Int.Check(int(0)) {
		t.Error("check for int failed")
	}
	if !Array.Check([1]int{}) {
		t.Error("check for array failed")
	}
	if !Slice.Check([]int{}) {
		t.Error("check for slice failed")
	}
	if !String.Check("") {
		t.Error("check for string failed")
	}
	if !Struct.Check(s{}) {
		t.Error("check for struct failed")
	}
	if !Map.Check(map[int]int{}) {
		t.Error("check for map failed")
	}
	if !Func.Check(TestCheck) {
		t.Error("check for func failed")
	}
	if !Pointer.Check(&s{}) {
		t.Error("check for pointer failed")
	}

	if !Nil.Check(nil) {
		t.Error("check for nil failed")
	}

	if !Bool.Check(false) {
		t.Error("check for bool failed")
	}

	if !Float.Check(23.3) {
		t.Error("check for float failed")
	}

	var c chan int
	if !Channel.Check(c) {
		t.Error("check for channel failed")
	}

	if !Complex.Check(complex(2, -2)) {
		t.Error("check for complex failed")
	}

	/*
		if !Interface.Check(i) {
			t.Error("check for interface failed")
		}
	*/
}
