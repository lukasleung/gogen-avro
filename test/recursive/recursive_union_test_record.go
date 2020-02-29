// Code generated by github.com/actgardner/gogen-avro. DO NOT EDIT.
/*
 * SOURCE:
 *     recursive.avsc
 */
package avro

import (
	"io"
	
	"github.com/actgardner/gogen-avro/vm/types"
	"github.com/actgardner/gogen-avro/vm"
	"github.com/actgardner/gogen-avro/compiler"
)

  
type RecursiveUnionTestRecord struct { 
	
	
		RecursiveField *UnionNullRecursiveUnionTestRecord
	

}

func DeserializeRecursiveUnionTestRecord(r io.Reader) (t RecursiveUnionTestRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeRecursiveUnionTestRecordFromSchema(r io.Reader, schema string) (t RecursiveUnionTestRecord, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeRecursiveUnionTestRecord(r RecursiveUnionTestRecord, w io.Writer) error {
	var err error
	
	err = writeUnionNullRecursiveUnionTestRecord( r.RecursiveField, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r RecursiveUnionTestRecord) Serialize(w io.Writer) error {
	return writeRecursiveUnionTestRecord(r, w)
}

func (r RecursiveUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"RecursiveField\",\"type\":[\"null\",\"RecursiveUnionTestRecord\"]}],\"name\":\"RecursiveUnionTestRecord\",\"type\":\"record\"}"
}

func (r RecursiveUnionTestRecord) SchemaName() string {
	return "RecursiveUnionTestRecord"
}

func (_ *RecursiveUnionTestRecord) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetInt(v int32) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetLong(v int64) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetString(v string) { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecursiveUnionTestRecord) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
			r.RecursiveField = &UnionNullRecursiveUnionTestRecord{}

		
		
			return r.RecursiveField
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *RecursiveUnionTestRecord) SetDefault(i int) {
	switch (i) { 
	default:
		panic("Unknown field index")
	}
}

func (r *RecursiveUnionTestRecord) Clear(i int) {
	switch (i) { 
	case 0:
		r.RecursiveField = nil
	default:
		panic("Non-optional field index")
	}
}

func (_ *RecursiveUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) Finalize() { }
