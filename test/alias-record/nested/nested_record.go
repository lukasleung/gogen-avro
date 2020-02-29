// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     nested.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type NestedRecord struct { 
	
	
		StringField string
	

	
	
		BoolField bool
	

	
	
		BytesField []byte
	

}

func DeserializeNestedRecord(r io.Reader) (t NestedRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeNestedRecordFromSchema(r io.Reader, schema string) (t NestedRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeNestedRecord(r NestedRecord, w io.Writer) error {
	var err error
	
	err = vm.WriteString( r.StringField, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteBool( r.BoolField, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteBytes( r.BytesField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r NestedRecord) Serialize(w io.Writer) error {
	return writeNestedRecord(r, w)
}

func (r NestedRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"NestedRecord\",\"type\":\"record\"}"
}

func (r NestedRecord) SchemaName() string {
	return "NestedRecord"
}

func (_ *NestedRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *NestedRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *NestedRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *NestedRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *NestedRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *NestedRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *NestedRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *NestedRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *NestedRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.String)(&r.StringField)
		
	
	case 1:
		
		
			return (*types.Boolean)(&r.BoolField)
		
	
	case 2:
		
		
			return (*types.Bytes)(&r.BytesField)
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *NestedRecord) SetDefault(i int) {
	switch (i) { 
	default:
		panic("Unknown field index")
	}
}

func (r *NestedRecord) Clear(i int) {
	switch (i) { 
	default:
		panic("Non-optional field index")
	}
}

func (_ *NestedRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *NestedRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *NestedRecord) Finalize() { }
