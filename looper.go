package looper

import (
	"reflect"
)

type sliceMapf func(interface{}, interface{}) interface{}
type mapMapf func(interface{}, interface{}) interface{}
type filterf func(interface{}) bool

type I struct {
	IIT
	Data      interface{}
	DataType  reflect.Type
	iterable_ bool
}

type IIT interface {
	Slice(sliceMapf) interface{}
	Map(mapMapf) interface{}
	Find(filterf) interface{}
	Filter(filterf) interface{}
}

func (i I) iterable() bool {
	k := i.DataType.Kind()
	return k == reflect.Array || k == reflect.Slice || k == reflect.Map
}
func (i I) iterator() I {
	i.iterable_ = i.iterable()
	if !i.iterable_ {
		panic("Value is not iterable.")
	}

	return i
}

func (i I) iterateMap(fn mapMapf) interface{} {
	var outGoing map[interface{}]interface{}
	var out reflect.Value

	//i.mustBe(reflect.Map)

	if i.DataType.Kind() == reflect.Map {
		onIterating := reflect.ValueOf(i.Data)
		iter := onIterating.MapRange()

		out = reflect.MakeMap(onIterating.Type())

		outGoing = make(map[interface{}]interface{}, onIterating.Len())
		for iter.Next() {
			if iter.Value().Kind() == reflect.Slice || iter.Value().Kind() == reflect.Array {

				out.SetMapIndex(iter.Key(), iter.Value())
				outGoing[iter.Key().Interface()] = fn(iter.Key().Interface(), iter.Value().Interface())
			}
		}
	}

	return outGoing
}
func (i I) iterateSlice(fn sliceMapf) interface{} {
	var out []interface{}
	if i.DataType.Kind() == reflect.Slice || i.DataType.Kind() == reflect.Array {
		onIterating := reflect.ValueOf(i.Data)

		out = make([]interface{}, onIterating.Len())
		for n := 0; n < onIterating.Len(); n++ {
			out[n] = fn(n, onIterating.Index(n).Interface())
		}
	}

	return out
}
func (i I) iterateFilter(fn filterf) interface{} {
	var out []interface{}
	if i.DataType.Kind() == reflect.Slice || i.DataType.Kind() == reflect.Array {
		onIterating := reflect.ValueOf(i.Data)
		out = make([]interface{}, 0, onIterating.Len())

		for n := 0; n < onIterating.Len(); n++ {
			current := onIterating.Index(n).Interface()
			if fn(current) {
				out = append(out, current)
			}
		}
	}

	return out
}
func (i I) iterateFind(fn filterf) interface{} {
	data := i.iterateFilter(fn)
	ref := reflect.ValueOf(data)
	if ref.Len() == 0 {
		return ""
	}
	return ref.Index(0).Interface()
}

//func (i I) mustBe(t interface{}) {
//	if i.DataType != t {
//		log.Panicf("Unexpected type for %v. Expected: %v Given: %v", i.Data, t, i.DataType)
//	}
//}
