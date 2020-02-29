// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
package avro

import (
	"io"
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeMapUnionNullStringInt(r MapUnionNullStringInt, w io.Writer) error {
	err := vm.WriteLong(int64(len(r.M)), w)
	if err != nil || len(r.M) == 0 {
		return err
	}
	for k, e := range r.M {
		err = vm.WriteString(k, w)
		if err != nil {
			return err
		}
		err = writeUnionNullStringInt(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0, w)
}

type MapUnionNullStringInt struct {
	keys []string
	values []*UnionNullStringInt
	M map[string]*UnionNullStringInt
}

func NewMapUnionNullStringInt() MapUnionNullStringInt{
	return MapUnionNullStringInt{ M: make(map[string]*UnionNullStringInt) }
}

func (_ *MapUnionNullStringInt) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetInt(v int32) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetLong(v int64) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetString(v string) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *MapUnionNullStringInt) SetDefault(i int) { panic("Unsupported operation") }
func (r *MapUnionNullStringInt) Finalize() {
	for i := range r.keys {
		r.M[r.keys[i]] = r.values[i]
	}
	r.keys = nil
	r.values = nil
}

func (r *MapUnionNullStringInt) AppendMap(key string) types.Field { 
	r.keys = append(r.keys, key)
	var v *UnionNullStringInt
	
		v = &UnionNullStringInt{}

	
	r.values = append(r.values, v)
	
	return r.values[len(r.values)-1]
	
}

func (r *MapUnionNullStringInt) Clear(i int) { 
	pos := len(r.keys) - 1
	switch {
	case pos < 0:
		panic("Map already empty")
	default:
		r.values[pos] = nil
	}
}

func (_ *MapUnionNullStringInt) AppendArray() types.Field { panic("Unsupported operation") }

