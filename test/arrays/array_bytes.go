// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     arrays.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
)

func writeArrayBytes(r [][]byte, w io.Writer) error {
	err := vm.WriteLong(int64(len(r)),w)
	if err != nil || len(r) == 0 {
		return err
	}
	for _, e := range r {
		err = vm.WriteBytes(e, w)
		if err != nil {
			return err
		}
	}
	return vm.WriteLong(0,w)
}

type ArrayBytes [][]byte

func (_ *ArrayBytes) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetInt(v int32) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetLong(v int64) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetString(v string) { panic("Unsupported operation") }
func (_ *ArrayBytes) SetUnionElem(v int64) { panic("Unsupported operation") }
func (_ *ArrayBytes) Get(i int) types.Field { panic("Unsupported operation") }
func (_ *ArrayBytes) Clear(i int) { panic("Unsupported operation") }
func (_ *ArrayBytes) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *ArrayBytes) Finalize() { }
func (_ *ArrayBytes) SetDefault(i int) { panic("Unsupported operation") }

func (r *ArrayBytes) AppendArray() types.Field {
	var v []byte
	
	*r = append(*r, v)
	
	return (*types.Bytes)(&(*r)[len(*r)-1])
	
}

