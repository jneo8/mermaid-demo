package main

import (
	"github.com/jneo8/mermaid"
	"github.com/jneo8/mermaid-demo/entity"
	"github.com/jneo8/mermaid-demo/repository/person"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	cmd.Flags().String("db_driver_name", "postgres", "Driver name")
	cmd.Flags().String("db_name", "mermaid_demo", "Database name")
	cmd.Flags().String("db_password", "pwd123", "DB password")
	cmd.Flags().String("db_username", "postgres", "DB password")
	cmd.Flags().String("db_host", "localhost", "DB host")
	cmd.Flags().Int("db_port", 5432, "DB port")
}

var cmd = &cobra.Command{
	Short: "Run service",
	RunE: func(cmd *cobra.Command, args []string) error {
		initializers := []interface{}{
			person.New,
			person.NewEngine,
		}
		runable := func(repo person.Repository) error {
			if err := repo.SyncTable(); err != nil {
				return err
			}
			if err := repo.InsertPerson(entity.Person{Name: "James", Phone: "0910xxxxxx"}); err != nil {
				return err
			}
			return nil
		}
		return mermaid.Run(cmd, runable, initializers...)
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
