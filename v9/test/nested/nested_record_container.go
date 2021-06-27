// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCE:
 *     schema.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/container"
	"github.com/actgardner/gogen-avro/v9/vm"
)

func NewNestedRecordWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewNestedRecord()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type NestedRecordReader struct {
	r io.Reader
	p *vm.Program
}

func NewNestedRecordReader(r io.Reader) (*NestedRecordReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewNestedRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &NestedRecordReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r NestedRecordReader) Read() (NestedRecord, error) {
	t := NewNestedRecord()
	err := vm.Eval(r.r, r.p, &t)
	return t, err
}