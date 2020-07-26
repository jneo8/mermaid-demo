package main

import "fmt"

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Short: "Run service",
	RunE: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	fmt.Println("vim-go")
}
