package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vladvalkov/goallocs"
	"github.com/vladvalkov/goallocs/internal/build"
)

var findCmd = &cobra.Command{
	Use:  "find [FLAGS] <target>",
	Args: cobra.MinimumNArgs(1),
	RunE: runFind,
}

func runFind(cmd *cobra.Command, args []string) error {
	reader, err := build.Build(args[0])
	if err != nil {
		return err
	}
	allocs, err := goallocs.Find(reader, goallocs.Config{})
	if err != nil {
		return err
	}

	for _, v := range allocs {
		fmt.Println(fmt.Sprintf("Allocation in %s at line %s", v.FuncName, v.Line))
	}
	return nil
}
