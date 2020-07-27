package flag

import (
	"github.com/spf13/cobra"
)

// AddDBFlag ...
func AddDBFlag(cmd *cobra.Command) {
	cmd.Flags().String("db_driver_name", "postgres", "Driver name")
	cmd.Flags().String("db_name", "mermaid_demo", "Database name")
	cmd.Flags().String("db_password", "pwd123", "DB password")
	cmd.Flags().String("db_username", "postgres", "DB password")
	cmd.Flags().String("db_host", "localhost", "DB host")
	cmd.Flags().Int("db_port", 5432, "DB port")
}

// AddFakeDataFlag ...
func AddFakeDataFlag(cmd *cobra.Command) {
	cmd.Flags().Int("n", 3, "Number of data")
	cmd.Flags().Int("min_age", 5, "Min age range")
	cmd.Flags().Int("max_age", 30, "Max age range")
}
