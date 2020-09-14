package main

type Assigner func(r io.Reader, f types.Field) error

func assign(r io.Reader, record *primitive.PrimitiveTestRecord) error {
	f := record.Get(0)
	err := assignInt(r, f)
	if err != nil {
		return err
	}

	f = record.Get(1)
	err = assignLong(r, f)
	if err != nil {
		return err
	}

	f = record.Get(2)
	err = assignFloat(r, f)
	if err != nil {
		return err
	}

	f = record.Get(3)
	err = assignDouble(r, f)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	p, err := newProgram()
	if err != nil {
		fmt.Printf("mmap error: %v", err)
	}
	p.callAssigner(reflect.ValueOf(assignBool).Pointer())

	//fn := p.funcPtr()
	//exeFn := *(*Assigner)(fn)

	src := primitive.PrimitiveTestRecord{
		IntField:    1,
		LongField:   2,
		FloatField:  3.2,
		DoubleField: 4.1,
		BoolField:   true,
		StringField: "string",
	}

	var buf bytes.Buffer
	err = src.Serialize(&buf)
	//var target primitive.PrimitiveTestRecord

	//err = exeFn(buf, &target)
	var target primitive.PrimitiveTestRecord
	f := assign(&buf, &target)
	fmt.Printf("Result: %v\n", f)
}