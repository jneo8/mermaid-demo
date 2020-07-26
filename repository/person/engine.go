package person

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"go.uber.org/dig"
)

// Config ...
type Config struct {
	dig.In
	DriverName string `name:"db_driver_name"`
	DBName     string `name:"db_name"` // Database name
	Host       string `name:"db_host"`
	Port       int    `name:"db_port"`
	Password   string `name:"db_password"`
	Username   string `name:"db_username"`
}

// NewEngine return new xorm engine.
func NewEngine(config Config) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine(
		config.DriverName,
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Host,
			config.Port,
			config.Username,
			config.Password,
			config.DBName,
		),
	)
	engine.ShowSQL()

	if err != nil {
		return engine, err
	}
	if err := engine.Ping(); err != nil {
		return engine, err
	}
	return engine, nil
}
