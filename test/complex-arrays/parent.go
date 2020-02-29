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
	"github.com/actgardner/gogen-avro/compiler"
)

  
type Parent struct { 
	
	
		Children []Child
	

}

func DeserializeParent(r io.Reader) (t Parent, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func DeserializeParentFromSchema(r io.Reader, schema string) (t Parent, err error) {
	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err == nil {
		err = vm.Eval(r, deser, &t)
	}
	return
}

func writeParent(r Parent, w io.Writer) error {
	var err error
	
	err = writeArrayChild( r.Children, w)
	if err != nil {
		return err			
	}
	
	return err
}

func (r Parent) Serialize(w io.Writer) error {
	return writeParent(r, w)
}

func (r Parent) Schema() string {
	return "{\"fields\":[{\"default\":[{\"name\":\"record1\"},{\"name\":\"record2\"}],\"name\":\"Children\",\"type\":{\"items\":{\"fields\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"Child\",\"type\":\"record\"},\"type\":\"array\"}}],\"name\":\"Parent\",\"type\":\"record\"}"
}

func (r Parent) SchemaName() string {
	return "Parent"
}

func (_ *Parent) SetBoolean(v bool) { panic("Unsupported operation") }
func (_ *Parent) SetInt(v int32) { panic("Unsupported operation") }
func (_ *Parent) SetLong(v int64) { panic("Unsupported operation") }
func (_ *Parent) SetFloat(v float32) { panic("Unsupported operation") }
func (_ *Parent) SetDouble(v float64) { panic("Unsupported operation") }
func (_ *Parent) SetBytes(v []byte) { panic("Unsupported operation") }
func (_ *Parent) SetString(v string) { panic("Unsupported operation") }
func (_ *Parent) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *Parent) Get(i int) types.Field {
	switch (i) {
	
	case 0:
		
		
			return (*ArrayChild)(&r.Children)
		
	
	default:
		panic("Unknown field index")
	}
}

func (r *Parent) SetDefault(i int) {
	switch (i) { 
	case 0:
		r.Children = make([]Child,2)
r.Children[0] = Child{}
r.Children[0].Name = "record1"

r.Children[1] = Child{}
r.Children[1].Name = "record2"


		
	default:
		panic("Unknown field index")
	}
}

func (r *Parent) Clear(i int) {
	switch (i) { 
	default:
		panic("Non-optional field index")
	}
}

func (_ *Parent) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *Parent) AppendArray() types.Field { panic("Unsupported operation") }
func (_ *Parent) Finalize() { }
