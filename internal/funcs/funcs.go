package funcs

type Function struct {
	Name         string
	Instructions []Instruction
}

func (f Function) IsEmpty() bool {
	return f.Name == ""
}

type Instruction struct {
	Line     string
	Name     string
	Operands string
}
