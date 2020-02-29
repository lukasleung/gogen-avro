// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

// Trace  
type Trace struct { 
	// Trace Identifier
	
		TraceId *UnionNullUUID
	

}

func DeserializeTrace(r io.Reader) (t Trace, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeTraceFromSchema(r io.Reader, schema string) (t Trace, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeTrace(r Trace, w io.Writer) error {
	var err error
	
	err = writeUnionNullUUID( r.TraceId, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r Trace) Serialize(w io.Writer) error {
	return writeTrace(r, w)
}

func (r Trace) Schema() string {
	return "{\"doc\":\"Trace\",\"fields\":[{\"default\":null,\"doc\":\"Trace Identifier\",\"name\":\"traceId\",\"type\":[\"null\",{\"doc\":\"A Universally Unique Identifier, in canonical form in lowercase. Example: de305d54-75b4-431b-adb2-eb6b9e546014\",\"fields\":[{\"default\":\"\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"UUID\",\"namespace\":\"headerworks.datatype\",\"type\":\"record\"}]}],\"name\":\"bodyworks.Trace\",\"type\":\"record\"}"
}

func (r Trace) SchemaName() string {
	return "bodyworks.Trace"
}

func (_ *Trace) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Trace) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Trace) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Trace) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Trace) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Trace) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Trace) SetString(v string) { panic("Unsupported operation") }
func (_ *Trace) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Trace) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.TraceId = &UnionNullUUID{}

		
		
			return r.TraceId
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *Trace) SetDefault(i int) {
	switch (i) { 
	case 0:
		r.TraceId = nil
		
	default:
		panic("Unknown field index")
	}
}

func (r *Trace) Clear(i int) {
	switch (i) { 
	case 0:
		r.TraceId = nil
	default:
		panic("Non-optional field index")
	}
}

func (_ *Trace) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Trace) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Trace) Finalize() { }
