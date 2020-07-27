package fakedata

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/jneo8/mermaid-demo/entity"
	"time"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano())
}

// Generator ...
type Generator interface {
	NewPerson() entity.Person
}

// NewGenerator ...
func NewGenerator() Generator {
	return &repo{N: 123}
}

type repo struct {
	N int
}

func (r *repo) NewPerson() entity.Person {
	var person entity.Person
	gofakeit.Struct(&person)
	return person
}
