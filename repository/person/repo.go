package person

import (
	"github.com/go-xorm/xorm"
	"github.com/jneo8/mermaid-demo/entity"
)

// Repository ...
type Repository interface {
	SyncTable() error
	InsertPerson(entity.Person) error
}

// New Repository.
func New(engine *xorm.Engine) Repository {
	return &repo{
		Engine: engine,
	}
}

type repo struct {
	Engine *xorm.Engine
}

func (r *repo) SyncTable() error {
	if err := r.Engine.Sync(new(entity.Person)); err != nil {
		return err
	}
	return nil
}

func (r *repo) InsertPerson(person entity.Person) error {
	_, err := r.Engine.Insert(person)
	return err
}
