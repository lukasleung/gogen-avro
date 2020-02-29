// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     primitives.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type PrimitiveTestRecord struct { 
	
	
		IntField int32
	

	
	
		LongField int64
	

	
	
		FloatField float32
	

	
	
		DoubleField float64
	

	
	
		StringField string
	

	
	
		BoolField bool
	

	
	
		BytesField []byte
	

}

func DeserializePrimitiveTestRecord(r io.Reader) (t PrimitiveTestRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializePrimitiveTestRecordFromSchema(r io.Reader, schema string) (t PrimitiveTestRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writePrimitiveTestRecord(r PrimitiveTestRecord, w io.Writer) error {
	var err error
	
	err = vm.WriteInt( r.IntField, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteLong( r.LongField, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteFloat( r.FloatField, w)
	if err != nil {
		return err			
	}
	
	err = vm.WriteDouble( r.DoubleField, w)
	if err != nil {
		return err			
	}
	
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

func (r PrimitiveTestRecord) Serialize(w io.Writer) error {
	return writePrimitiveTestRecord(r, w)
}

func (r PrimitiveTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"IntField\",\"type\":\"int\"},{\"name\":\"LongField\",\"type\":\"long\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"DoubleField\",\"type\":\"double\"},{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"BoolField\",\"type\":\"boolean\"},{\"name\":\"BytesField\",\"type\":\"bytes\"}],\"name\":\"PrimitiveTestRecord\",\"type\":\"record\"}"
}

func (r PrimitiveTestRecord) SchemaName() string {
	return "PrimitiveTestRecord"
}

func (_ *PrimitiveTestRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *PrimitiveTestRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*types.Int)(&r.IntField)
		
	
	case 1:
		
		
			return (*types.Long)(&r.LongField)
		
	
	case 2:
		
		
			return (*types.Float)(&r.FloatField)
		
	
	case 3:
		
		
			return (*types.Double)(&r.DoubleField)
		
	
	case 4:
		
		
			return (*types.String)(&r.StringField)
		
	
	case 5:
		
		
			return (*types.Boolean)(&r.BoolField)
		
	
	case 6:
		
		
			return (*types.Bytes)(&r.BytesField)
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *PrimitiveTestRecord) SetDefault(i int) {
	switch (i) { 
	default:
		panic("Unknown field index")
	}
}

func (r *PrimitiveTestRecord) Clear(i int) {
	switch (i) { 
	default:
		panic("Non-optional field index")
	}
}

func (_ *PrimitiveTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *PrimitiveTestRecord) Finalize() { }
