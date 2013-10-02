package meta

import (
	"fmt"
	"reflect"
)

func (ø struct_) Field(s interface{}, field string) (f reflect.Value) {
	fv := ø.FinalValue(s)
	f = fv.FieldByName(field)
	return
}

// get an attribute of a struct into the target
func (ø struct_) Get(s interface{}, field string, t interface{}) {
	if Nil.Check(s) {
		return
	}
	fv := ø.FinalValue(s)

	if !(fv.Type().Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}

	if !ø.Is(reflect.Ptr, t) {
		Panicf("%s is not a pointer", Inspect(t))
	}

	p := fv.FieldByName(field)
	reflect.ValueOf(t).Elem().Set(p)
}

// can be used to get the raw structfield and value and interate through them
func (ø struct_) EachRaw(s interface{}, fn func(field reflect.StructField, val reflect.Value)) {
	if Nil.Check(s) {
		return
	}
	ft := ø.FinalType(s)
	fv := ø.FinalValue(s)
	if !(fv.Type().Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}

	elem := ft.NumField()
	for i := 0; i < elem; i++ {
		fn(ft.Field(i), fv.Field(i))
	}
	return
}

// can be used to get attribute names and values, for
// tags run Tags / Tag
// use reflect.TypeOf() to get the type of the val
func (ø struct_) Each(s interface{}, fn func(field string, val interface{})) {
	if Nil.Check(s) {
		return
	}
	ft := ø.FinalType(s)
	fv := ø.FinalValue(s)
	if !(fv.Type().Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}

	elem := ft.NumField()
	for i := 0; i < elem; i++ {
		f := ft.Field(i)
		v := fv.Field(i)
		//fn(f.Type, f.Name, v.Interface())
		fn(f.Name, v.Interface())
	}
	return
}

// returns all struct tags
func (ø struct_) Tags(s interface{}) (tags map[string]*reflect.StructTag) {
	tags = map[string]*reflect.StructTag{}
	if Nil.Check(s) {
		return
	}
	ft := ø.FinalType(s)
	if !(ft.Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}
	elem := ft.NumField()
	for i := 0; i < elem; i++ {
		f := ft.Field(i)
		if string(f.Tag) != "" {
			tags[f.Name] = &f.Tag
		}
	}
	return
}

// returns a struct tag for a field
func (ø struct_) Tag(s interface{}, field string) *reflect.StructTag {
	if Nil.Check(s) {
		return nil
	}

	ft := ø.FinalType(s)
	if !(ft.Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}
	f, exists := ft.FieldByName(field)
	if !exists {
		panic(fmt.Sprintf("field %s does not exist in %s", field, Inspect(s)))
	}
	return &f.Tag
}

// first checks, if there is a function for each Type of the struct fields
// if not, it panics. This panic should happen at initialization time.
// therefor Dispatch returns a function that may be called later
// this later function must be given a value of the same type,
// otherwise a runtime panic happens
func (ø struct_) Dispatch(s interface{}, m map[reflect.Type]func(field string, val interface{}) error) func(interface{}) map[string]error {
	ft := ø.FinalType(s)
	if !(ft.Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}
	for _, t := range ø.Types(s) {
		if m[t] == nil {
			panic(fmt.Sprintf("can't find handler for type %s in map %v", t, m))
		}
	}

	f := func(s_ interface{}) map[string]error {
		if reflect.TypeOf(s) != reflect.TypeOf(s_) {
			panic(
				fmt.Errorf(
					"definition type (%s) does not match type of passed object (%s)",
					reflect.TypeOf(s),
					reflect.TypeOf(s_),
				))
		}
		errs := map[string]error{}
		inner := func(fie string, va interface{}) {
			e := m[reflect.TypeOf(va)](fie, va)
			if e != nil {
				errs[fie] = e
			}
		}
		ø.Each(s_, inner)
		return errs
	}
	return f
}

func (ø struct_) Types(s interface{}) (t []reflect.Type) {
	t = []reflect.Type{}
	if Nil.Check(s) {
		return
	}
	ft := ø.FinalType(s)
	if !(ft.Kind() == reflect.Struct) {
		Panicf("%s is not a struct / pointer to a struct", Inspect(s))
	}
	has := func(tt reflect.Type) bool {
		for _, ttt := range t {
			if ttt == tt {
				return true
			}
		}
		return false
	}
	elem := ft.NumField()
	for i := 0; i < elem; i++ {
		f := ft.Field(i)
		if !has(f.Type) {
			t = append(t, f.Type)
		}
	}
	return
}

// set an attribute of a struct
func (ø struct_) Set(s interface{}, field string, val interface{}) {
	if Nil.Check(s) {
		return
	}
	if Nil.Check(val) {
		panic("setting to nil is not allowed")
	}
	if !ø.IsPointerTo(reflect.Struct, s) {
		panic(fmt.Sprintf("is no pointer to struct: %s", Inspect(s)))
	}

	p := ø.FinalValue(s).FieldByName(field)

	if p.Type().Kind() != reflect.TypeOf(val).Kind() {
		panic(fmt.Sprintf("field %s has type %s but %#v is not a compatible type", field, p.Type(), Inspect(val)))
	}

	if p.CanSet() {
		p.Set(reflect.ValueOf(val))
	} else {
		panic("can't set field " + field + " of " + Inspect(s))
	}
}

// sets an field that is a function

//only works with go1.1
func (ø struct_) SetFunc(s interface{}, field string, f func(reflect.Value, []reflect.Value) []reflect.Value) {
	if Nil.Check(s) {
		return
	}
	if !ø.IsPointerTo(reflect.Struct, s) {
		panic(fmt.Sprintf("is no pointer to struct: %s", Inspect(s)))
	}

	fv := ø.FinalValue(s)
	p := fv.FieldByName(field)

	generic := func(in []reflect.Value) []reflect.Value { return f(fv, in) }
	v := reflect.MakeFunc(p.Type(), generic)
	p.Set(v)
}

// sets a field on a struct a with value b
// a and b are casted to their internal values so that
// different mixtures of a's and b's are possible
func (ø struct_) PolySet(a interface{}, field string, b interface{}) {
	x := reflect.ValueOf(a).Interface()
	y := reflect.ValueOf(b).Interface()
	ø.Set(x, field, y)
}
