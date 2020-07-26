package main

import (
	"github.com/spf13/cobra"
	"os"
)

var cmd = &cobra.Command{
	Short: "Run service",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
