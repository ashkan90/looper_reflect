package looper

import "reflect"

func NewIterator(d interface{}) IIT {
	return &I{
		Data:      d,
		DataType:  reflect.TypeOf(d),
		iterable_: false,
	}
}

func (i I) Slice(fn sliceMapf) interface{} {
	return i.iterator().iterateSlice(fn)
}
func (i I) Map(fn mapMapf) interface{} {
	return i.iterator().iterateMap(fn)
}
func (i I) Find(fn filterf) interface{} {
	return i.iterator().iterateFind(fn)
}
func (i I) Filter(fn filterf) interface{} {
	return i.iterator().iterateFilter(fn)
}
