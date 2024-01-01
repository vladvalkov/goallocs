package analyze

import (
	"github.com/samber/lo"
	"github.com/vladvalkov/goallocs/internal/funcs"
)

type AllocationCheck func(function funcs.Function, i int) bool

var AllocatingCalls = []AllocationCheck{
	isNewObject,
	isMakeSlice,
	isMakeMap,
	isMakeMapSmall,
}

func Analyze(functions []funcs.Function) []Allocation {
	var allocations []Allocation

	for _, f := range functions {
		for i, inst := range f.Instructions {
			if lo.ContainsBy[AllocationCheck](AllocatingCalls, func(item AllocationCheck) bool {
				return item(f, i)
			}) {
				allocations = append(allocations, Allocation{
					Line:     inst.Line,
					FuncName: f.Name,
				})
			}
		}
	}

	return allocations
}

type Allocation struct {
	Line     string
	FuncName string
}
