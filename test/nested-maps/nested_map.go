// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     nested-maps.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type NestedMap struct { 
	
	
		MapOfMaps MapMapArrayString
	

}

func DeserializeNestedMap(r io.Reader) (t NestedMap, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeNestedMapFromSchema(r io.Reader, schema string) (t NestedMap, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeNestedMap(r NestedMap, w io.Writer) error {
	var err error
	
	err = writeMapMapArrayString( r.MapOfMaps, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r NestedMap) Serialize(w io.Writer) error {
	return writeNestedMap(r, w)
}

func (r NestedMap) Schema() string {
	return "{\"fields\":[{\"name\":\"MapOfMaps\",\"type\":{\"type\":\"map\",\"values\":{\"type\":\"map\",\"values\":{\"items\":\"string\",\"type\":\"array\"}}}}],\"name\":\"NestedMap\",\"type\":\"record\"}"
}

func (r NestedMap) SchemaName() string {
	return "NestedMap"
}

func (_ *NestedMap) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *NestedMap) SetInt(v int32) { panic("Unsupported operation") }
func (_ *NestedMap) SetLong(v int64) { panic("Unsupported operation") }
func (_ *NestedMap) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *NestedMap) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *NestedMap) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *NestedMap) SetString(v string) { panic("Unsupported operation") }
func (_ *NestedMap) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedMap) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.MapOfMaps = NewMapMapArrayString()

		
		
			return &r.MapOfMaps
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *NestedMap) SetDefault(i int) {
	switch (i) { 
	default:
		panic("Unknown field index")
	}
}

func (r *NestedMap) Clear(i int) {
	switch (i) { 
	default:
		panic("Non-optional field index")
	}
}

func (_ *NestedMap) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NestedMap) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *NestedMap) Finalize() { }
