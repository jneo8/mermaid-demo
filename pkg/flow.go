package flow

import (
	"github.com/jneo8/mermaid-demo/pkg/fakedata"
	"github.com/jneo8/mermaid-demo/repository/person"
	"go.uber.org/dig"
)

// Opts ...
type Opts struct {
	dig.In
	N int `name:"n"`
}

// Run ...
func Run(opts Opts, repo person.Repository, generator fakedata.Generator) error {
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
