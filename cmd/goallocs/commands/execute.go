package commands

import (
	"fmt"
	"os"
)

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(findCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
