package fakedata

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/jneo8/mermaid-demo/entity"
	"go.uber.org/dig"
	"time"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

// Config ...
type Config struct {
	dig.In
	MaxAge int `name:"max_age"`
	MinAge int `name:"min_age"`
}

// Generator ...
type Generator interface {
	NewPerson() entity.Person
}

// NewGenerator ...
func NewGenerator(config Config) Generator {
	lookups := map[string]gofakeit.Info{
		"age": gofakeit.Info{
			Category: "Person",
			Output:   "int",
			Call: func(m *map[string][]string, info *gofakeit.Info) (interface{}, error) {
				return gofakeit.Number(config.MinAge, config.MaxAge), nil
			},
		},
	}
	for k, v := range lookups {
		gofakeit.AddFuncLookup(k, v)
	}
	return &repo{Config: config}
}

type repo struct {
	Config Config
}

func (r *repo) NewPerson() entity.Person {
	var person entity.Person
	gofakeit.Struct(&person)
	return person
}
