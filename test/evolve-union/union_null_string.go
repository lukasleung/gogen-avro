// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     defaults.avsc
 */
package avro

import (
	"errors"
	"io"

	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/vm/types"
)


type UnionNullStringTypeEnum int

const (
	UnionNullStringTypeEnumString UnionNullStringTypeEnum = 1
)

type UnionNullString struct { 
	String string

	UnionType UnionNullStringTypeEnum
}

func writeUnionNullString(r *UnionNullString, w io.Writer) error { 
	if r == nil {
		return vm.WriteLong(int64(0), w)
	} 
	if err := vm.WriteLong(int64(r.UnionType), w); err != nil {
		return err
	}
	switch r.UnionType{ 
	case UnionNullStringTypeEnumString:
		return vm.WriteString(r.String, w)
	}
	return errors.New("invalid value for *UnionNullString")
}

func (_ *UnionNullString) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *UnionNullString) SetInt(v int32) { panic("Unsupported operation") }
func (_ *UnionNullString) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *UnionNullString) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *UnionNullString) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *UnionNullString) SetString(v string) { panic("Unsupported operation") }

func (r *UnionNullString) SetLong(v int64) { 
	r.UnionType = (UnionNullStringTypeEnum)(v)
}

func (r *UnionNullString) Get(i int) types.Field {
	switch (i) { 
	case 1:
		
		
		return (*types.String)(&r.String)
		
	
	}
	panic("Unknown field index")
}

func (r *UnionNullString) Clear(i int) { panic("Unsupported operation") }
func (_ *UnionNullString) SetDefault(i int) { panic("Unsupported operation") }
func (_ *UnionNullString) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *UnionNullString) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *UnionNullString) Finalize()  { }
