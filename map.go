package meta

import (
	"fmt"
	"reflect"
)

// puts the values of the map to the fields
// of a struct that have the same name as the keys of the map
func (ø map_) ToStruct(m interface{}, s interface{}) {
	if Nil.Check(m) {
		panic("map is nil")
	}
	if Nil.Check(s) {
		panic("struct is nil")
	}

	ft := ø.FinalType(m)

	if ft.Kind() != reflect.Map {
		panic(fmt.Sprintf("%s is not a map", Inspect(m)))
	}

	if ft.Key().Kind() != reflect.String {
		panic(fmt.Sprintf("map %s should have strings as keys", Inspect(m)))
	}

	if !ø.IsPointerTo(reflect.Struct, s) {
		panic(fmt.Sprintf("%s is no pointer to struct", Inspect(s)))
	}

	sv := ø.FinalValue(s)
	mm := ø.FinalValue(m)
	keys := mm.MapKeys()

	for _, key := range keys {
		k := key.String()
		//vv := reflect.ValueOf(mm.MapIndex(key))
		e := mm.MapIndex(key)
		if mm.MapIndex(key).Type().Kind() == reflect.Interface {
			e = e.Elem()
		}
		sv.FieldByName(k).Set(e)
		//Struct.Set(s, k, mm.MapIndex(key))
	}
}
