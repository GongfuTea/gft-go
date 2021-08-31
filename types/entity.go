package types

import (
	"fmt"

	"github.com/google/uuid"
)

type IEntity interface {
	ID() string
	IsNew() bool
	Init()
}

type Entity struct {
	Id string `bson:"_id,omitempty" json:"id,omitempty"`
}

func (e Entity) ID() string {
	return e.Id
}

func (e Entity) IsNew() bool {
	isNew := e.Id == ""
	return isNew
}

func (e *Entity) Init() {
	fmt.Printf("init Entity, %+v\n", e)

	e.Id = uuid.NewString()
}
