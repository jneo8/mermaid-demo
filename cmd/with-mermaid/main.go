package main

import (
	"github.com/jneo8/mermaid"
	"github.com/jneo8/mermaid-demo/cmd/flag"
	"github.com/jneo8/mermaid-demo/pkg/fakedata"
	"github.com/jneo8/mermaid-demo/pkg/flow"
	"github.com/jneo8/mermaid-demo/repository/person"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	flag.AddDBFlag(cmd)
	flag.AddFakeDataFlag(cmd)
}

var cmd = &cobra.Command{
	Short: "Run service",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			person.New,
			person.NewEngine,
			fakedata.NewGenerator,
		}
		return mermaid.Run(cmd, flow.Run, initializers...)
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
