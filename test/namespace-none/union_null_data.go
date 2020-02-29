// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"errors"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullDataTypeEnum int

const (
	UnionNullDataTypeEnumData UnionNullDataTypeEnum = 1
)

type UnionNullData struct { 
	Data Data

	UnionType UnionNullDataTypeEnum
}

func writeUnionNullData(r *UnionNullData, w io.Writer) error { 
	if r == nil {
		return vm.WriteLong(int64(0), w)
	} 
	if err := vm.WriteLong(int64(r.UnionType), w); err != nil {
		return err
	}
	switch r.UnionType{ 
	case UnionNullDataTypeEnumData:
		return writeData(r.Data, w)
	}
	return errors.New("invalid value for *UnionNullData")
}

func (_ *UnionNullData) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullData) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullData) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullData) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullData) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullData) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullData) SetLong(v int64) { 
	r.UnionType = (UnionNullDataTypeEnum)(v)
}

func (r *UnionNullData) Get(i int) types.Field {
	switch (i) { 
	case 1:
		
		r.Data = Data{}
		
		
		return &r.Data
		
	
	}
	panic("Unknown field index")
}

func (r *UnionNullData) Clear(i int) { panic("Unsupported operation") }
func (_ *UnionNullData) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullData) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullData) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullData) Finalize()  { }
