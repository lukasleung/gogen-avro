// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     namespace.avsc
 */
package avro

import (
	"io"

	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/container"
	"github.com/actgardner/gogen-avro/v7/vm"
)

func NewUUIDWriter(writer io.Writer, codec container.Codec, recordsPerBlock int64) (*container.Writer, error) {
	str := NewUUID()
	return container.NewWriter(writer, codec, recordsPerBlock, str.Schema())
}

// container reader
type UUIDReader struct {
	r io.Reader
	p *vm.Program
}

func NewUUIDReader(r io.Reader) (*UUIDReader, error) {
	containerReader, err := container.NewReader(r)
	if err != nil {
		return nil, err
	}

	t := NewUUID()
	deser, err := compiler.CompileSchemaBytes([]byte(containerReader.AvroContainerSchema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	return &UUIDReader{
		r: containerReader,
		p: deser,
	}, nil
}

func (r UUIDReader) Read() (*UUID, error) {
	t := NewUUID()
	err := vm.Eval(r.r, r.p, t)
	return t, err
}