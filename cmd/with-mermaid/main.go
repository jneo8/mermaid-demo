package main

import (
	"github.com/jneo8/mermaid"
	"github.com/jneo8/mermaid-demo/pkg/fakedata"
	"github.com/jneo8/mermaid-demo/repository/person"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"os"
)

func init() {
	cmd.Flags().String("db_driver_name", "postgres", "Driver name")
	cmd.Flags().String("db_name", "mermaid_demo", "Database name")
	cmd.Flags().String("db_password", "pwd123", "DB password")
	cmd.Flags().String("db_username", "postgres", "DB password")
	cmd.Flags().String("db_host", "localhost", "DB host")
	cmd.Flags().Int("db_port", 5432, "DB port")
	cmd.Flags().Int("n", 3, "Number of data")
}

// Opts ...
type opts struct {
	dig.In
	N int `name:"n"`
}

func flow(opts opts, repo person.Repository, generator fakedata.Generator) error {
	if err := repo.SyncTable(); err != nil {
		return err
	}
	for i := 0; i < opts.N; i++ {
		if err := repo.InsertPerson(generator.NewPerson()); err != nil {
			return err
		}
	}
	return nil
}

var cmd = &cobra.Command{
	Short: "Run service",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			person.New,
			person.NewEngine,
			fakedata.NewGenerator,
		}
		return mermaid.Run(cmd, flow, initializers...)
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
