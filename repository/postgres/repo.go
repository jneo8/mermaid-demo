package postgres

import (
	"github.com/go-xorm"
	"go.uber.org/dig"
)

// Config ...
type Config struct {
	dig.In
	DriverName     string `name:"driver_name"`
	DataSourceName string `name:"data_source_name"`
}

// New return new xorm Engine.
func New(config Config) (*xorm.Engine, error) {
	return xorm.NewEngine(config.DriverName, DriverSourceName)
}
