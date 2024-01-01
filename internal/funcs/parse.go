package funcs

import (
	"bufio"
	"io"
	"strings"
)

const assemblyFunction = "STEXT"

func Parse(reader io.Reader) ([]Function, error) {
	var functions []Function

	scanner := bufio.NewScanner(reader)

	var cur Function
	for scanner.Scan() {
		line := scanner.Text()
		f, ok := ParseFunction(line)
		if ok {
			functions = append(functions, cur)
			cur = f
			continue
		}

		ins, ok := ParseInstruction(line)
		if ok {
			cur.Instructions = append(cur.Instructions, ins)
		}
	}
	if !cur.IsEmpty() {
		functions = append(functions, cur)
	}

	return functions, nil
}

func ParseFunction(s string) (Function, bool) {
	if !strings.Contains(s, assemblyFunction) || strings.HasPrefix(s, "\t") {
		return Function{}, false
	}

	name := strings.SplitN(s, " ", 2)[0]
	return Function{
		Name: name,
	}, true
}

func ParseInstruction(s string) (Instruction, bool) {
	if !strings.HasPrefix(s, "\t") {
		return Instruction{}, false
	}
	s = strings.ReplaceAll(s, "\t", " ")
	s = strings.ReplaceAll(s, " ", " ")
	s = strings.TrimSpace(s)

	splitN := strings.SplitN(s, " ", 5)

	if len(splitN) != 5 {
		return Instruction{}, false
	}
	if len(splitN[0]) != 6 {
		return Instruction{}, false
	}
	if len(splitN[1]) != 5 {
		return Instruction{}, false
	}

	line := splitN[2]
	instruction := splitN[3]

	return Instruction{
		Line:     line,
		Name:     instruction,
		Operands: splitN[4],
	}, true
}
