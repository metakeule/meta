package meta

import (
	"fmt"
	"reflect"
	"testing"
)

type S struct {
	A string `hi:"ho"`
	B int
}

func TestGet(t *testing.T) {
	s := S{"hu", 2}

	var r string
	Struct.Get(&s, "A", &r)
	if r != s.A {
		t.Errorf("error in Get(A): got %#v, expected %#v\n", r, s.A)
	}
	var i int
	Struct.Get(s, "B", &i)
	if i != s.B {
		t.Errorf("error in Get(B): got %#v, expected %#v\n", i, s.B)
	}
}

type fieldvalue struct {
	val interface{}
	typ reflect.Type
}

func TestEach(t *testing.T) {
	s := S{"hu", 2}

	res := map[string]fieldvalue{}

	Struct.Each(&s, func(field string, val interface{}) {
		res[field] = fieldvalue{val, reflect.TypeOf(val)}
	})

	if len(res) != 2 {
		t.Errorf("error in len of each (*struct): got %#v, expected %#v\n", len(res), 2)
	}

	fa := res["A"]
	if fa.typ != reflect.TypeOf("") {
		t.Errorf("error in type of field A: is %s should be %s", fa.typ.String(), reflect.TypeOf("").String())
	}

	fav, _ := fa.val.(string)
	if fav != s.A {
		t.Errorf("error in value of field A: is %s should be %s", fav, s.A)
	}

	fb := res["B"]
	if fb.typ != reflect.TypeOf(0) {
		t.Errorf("error in type of field B: is %s should be %s", fb.typ.String(), reflect.TypeOf(0).String())
	}

	fbv, _ := fb.val.(int)
	if fbv != s.B {
		t.Errorf("error in value of field B: is %s should be %s", fbv, s.B)
	}

	res = map[string]fieldvalue{}

	Struct.Each(s, func(field string, val interface{}) {
		res[field] = fieldvalue{val, reflect.TypeOf(val)}
	})

	if len(res) != 2 {
		t.Errorf("error in len of each (struct): got %#v, expected %#v\n", len(res), 2)
	}

	var aTag string
	var aVal string
	Struct.EachTag(s, "hi", func(field reflect.StructField, val reflect.Value, tag string) {
		if field.Name == "A" {
			aTag = tag
			aVal = val.String()
		}
	})

	if aTag != "ho" {
		t.Errorf("wrong tag value: got: %#v, expected: %#v", aTag, "ho")
	}

	if aVal != "hu" {
		t.Errorf("wrong value: got: %#v, expected: %#v", aVal, "hu")
	}
}

func TestTags(t *testing.T) {
	s := S{}
	r := Struct.Tags(&s)

	if len(r) != 1 {
		t.Errorf("error in len of Tags (*struct): got %#v, expected %#v (%#v)\n", len(r), 1, r)
	}

	if r_ := r["A"]; r_ == nil {
		t.Errorf("error in Tag for field A: not found")
	}

	if r_ := r["A"].Get("hi"); r_ != "ho" {
		t.Errorf("error in Tag key hi: got %#v, expected %#v (%#v)\n", r_, "ho", string(*r["A"]))
	}

	r = Struct.Tags(s)

	if len(r) != 1 {
		t.Errorf("error in len of Tags (struct): got %#v, expected %#v (%#v)\n", len(r), 1, r)
	}

	if r_ := Struct.Tag(s, "A").Get("hi"); r_ != "ho" {
		t.Errorf(`error in direct Tag for "A" key hi: got %#v, expected %#v`, r_, "ho")
	}
}

func TestTypes(t *testing.T) {
	s := S{}
	r := Struct.Types(&s)
	has := func(tt reflect.Type) bool {
		for _, ttt := range r {
			if ttt == tt {
				return true
			}
		}
		return false
	}

	if len(r) != 2 {
		t.Errorf("error in len of Types (*struct): got %#v, expected %#v (%#v)\n", len(r), 2, r)
	}

	if !has(reflect.TypeOf("")) {
		t.Errorf("error in Types, missing %#v", reflect.TypeOf("").String())
	}

	if !has(reflect.TypeOf(0)) {
		t.Errorf("error in Types, missing %#v", reflect.TypeOf(0).String())
	}
}

func TestSet(t *testing.T) {
	s := &S{"hi", 24}
	Struct.Set(s, "A", "ho")
	Struct.Set(s, "B", 42)

	if s.A != "ho" {
		t.Errorf(`error in Set("A"),  got %#v, expected %#v`, s.A, "ho")
	}

	if s.B != 42 {
		t.Errorf(`error in Set("B"),  got %#v, expected %#v`, s.B, 42)
	}
}

func TestDispatch(t *testing.T) {
	var i int
	m := map[reflect.Type]func(field string, val interface{}) error{
		reflect.TypeOf(""): func(field string, val interface{}) error {
			if field == "A" {
				return fmt.Errorf(val.(string))
			}
			return nil
		},
		reflect.TypeOf(0): func(field string, val interface{}) error {
			if field == "B" {
				i = val.(int)
			}
			return nil
		},
	}
	s := S{"ho", 42}

	fn := Struct.Dispatch(&s, m)
	errs := fn(&s)

	if i != 42 {
		t.Errorf(`error in Dispatch for field "B",  got %#v, expected %#v`, i, 42)
	}

	if len(errs) != 1 {
		t.Errorf(`error in Dispatch len of errs,  got %#v, expected %#v`, len(errs), 1)
	}

	if errs["A"].Error() != "ho" {
		t.Errorf(`error in Dispatch errs,  got %#v, expected %#v`, errs["A"], "ho")
	}
}

type fS struct {
	F   func(int, int) (int, int)
	Sum int64
}

type ifS interface {
	F(int int) (int, int)
}

func TestSetFunc(t *testing.T) {
	fs := &fS{}

	fn := func(stru reflect.Value, in []reflect.Value) (out []reflect.Value) {
		sum := in[0].Int() + in[1].Int()
		stru.FieldByName("Sum").Set(reflect.ValueOf(sum))
		return []reflect.Value{in[1], in[0]}
	}
	Struct.SetFunc(fs, "F", fn)

	a, b := fs.F(1, 2)

	if a != 2 {
		t.Errorf(`error in SetFunc,  got %#v, expected %#v`, a, 2)
	}

	if b != 1 {
		t.Errorf(`error in SetFunc,  got %#v, expected %#v`, b, 1)
	}

	if fs.Sum != 3 {
		t.Errorf(`error in SetFunc(Sum),  got %#v, expected %#v`, fs.Sum, 3)
	}

	fs2 := &fS{}
	// has to be set newly for every instance
	Struct.SetFunc(fs2, "F", fn)

	a, b = fs2.F(4, 3)

	if a != 3 {
		t.Errorf(`error in SetFunc(new),  got %#v, expected %#v`, a, 3)
	}

	if b != 4 {
		t.Errorf(`error in SetFunc(new),  got %#v, expected %#v`, b, 4)
	}

	// does not work, due to limitations in go, see
	// https://groups.google.com/d/msg/golang-nuts/Hx0XWpV0HgE/-VVoZwpFALUJ
	// var ifS_ ifS
	// ifS_ = fs2

	// a, b = ifS_.F(6, 5)

	// if a != 5 {
	// 	t.Errorf(`error in SetFunc(new),  got %#v, expected %#v`, a, 5)
	// }

	// if b != 6 {
	// 	t.Errorf(`error in SetFunc(new),  got %#v, expected %#v`, b, 6)
	// }

}

type a1 struct{ B b1 }
type a2 struct{ B b2 }
type b1 struct{ n int }
type b2 struct{ i int }

func poly(a interface{}, field string, b interface{}) {
	Struct.PolySet(a, field, b)
}

func TestPolySet(t *testing.T) {
	a1_ := &a1{}
	a2_ := &a2{}
	b1_ := b1{5}
	b2_ := b2{6}

	poly(a1_, "B", b1_)
	poly(a2_, "B", b2_)

	if a1_.B.n != 5 {
		t.Errorf(`error in PolySet(a1_.B.n),  got %#v, expected %#v`, a1_.B.n, 5)
	}

	if a2_.B.i != 6 {
		t.Errorf(`error in PolySet(a2_.B.i),  got %#v, expected %#v`, a2_.B.i, 6)
	}
}

//func (ø struct_) PolySet(a interface{}, field string, b interface{}) {
