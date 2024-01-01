package goallocs

import (
	"github.com/samber/lo"
	"github.com/vladvalkov/goallocs/internal/funcs"
	"github.com/vladvalkov/goallocs/internal/funcs/analyze"
	"io"
)

type Config struct {
	TargetFunctions []string
}

type Allocation struct {
	Line     string
	FuncName string
}

func Find(assemblyReader io.Reader, config Config) ([]Allocation, error) {
	functions, err := funcs.Parse(assemblyReader)
	if err != nil {
		return nil, err
	}
	return lo.Map(analyze.Analyze(functions), func(item analyze.Allocation, index int) Allocation {
		return Allocation(item)
	}), nil
}
