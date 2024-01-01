package analyze

import (
	"github.com/vladvalkov/goallocs/internal/funcs"
	"strings"
)

func isNewObject(function funcs.Function, i int) bool {
	instruction := function.Instructions[i]
	return instruction.Name == "CALL" && strings.Contains(instruction.Operands, "runtime.newobject")
}

func isMakeSlice(function funcs.Function, i int) bool {
	instruction := function.Instructions[i]
	return instruction.Name == "CALL" && strings.Contains(instruction.Operands, "runtime.makeslice")
}

func isMakeMap(function funcs.Function, i int) bool {
	instruction := function.Instructions[i]
	isCallingMakeMap := instruction.Name == "CALL" && strings.Contains(instruction.Operands, "runtime.makemap")
	if !isCallingMakeMap {
		return false
	}
	return i >= 2 && function.Instructions[i-2].Operands == "ZR, R2"
}

func isMakeMapSmall(function funcs.Function, i int) bool {
	instruction := function.Instructions[i]
	return instruction.Name == "CALL" && strings.Contains(instruction.Operands, "runtime.makemap_small")
}
