package main

import (
	"github.com/jneo8/mermaid-demo/cmd/flag"
	"github.com/jneo8/mermaid-demo/pkg/fakedata"
	"github.com/jneo8/mermaid-demo/pkg/flow"
	"github.com/jneo8/mermaid-demo/repository/person"
	log "github.com/sirupsen/logrus"
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
		driverName, _ := cmd.Flags().GetString("db_driver_name")
		dbName, _ := cmd.Flags().GetString("db_name")
		host, _ := cmd.Flags().GetString("db_host")
		port, _ := cmd.Flags().GetInt("db_port")
		username, _ := cmd.Flags().GetString("db_username")
		password, _ := cmd.Flags().GetString("db_password")
		maxAge, _ := cmd.Flags().GetInt("max_age")
		minAge, _ := cmd.Flags().GetInt("min_age")
		n, _ := cmd.Flags().GetInt("n")

		engine, err := person.NewEngine(
			person.Config{
				DriverName: driverName,
				DBName:     dbName,
				Host:       host,
				Port:       port,
				Username:   username,
				Password:   password,
			},
		)
		if err != nil {
			return err
		}
		repository := person.New(engine)
		generator := fakedata.NewGenerator(
			fakedata.Config{
				MaxAge: maxAge,
				MinAge: minAge,
			},
		)
		return flow.Run(log.New(), flow.Opts{N: n}, repository, generator)
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
