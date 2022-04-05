package types

import (
	"fmt"

	"github.com/google/uuid"
)

type IEntity interface {
	ID() string
	Init()
}

type Entity struct {
	Id string `bson:"_id,omitempty" json:"id,omitempty"`
}

func (e Entity) ID() string {
	return e.Id
}

func (e *Entity) Init() {

	e.Id = uuid.NewString()

	fmt.Printf("init Entity, %+v\n", e)

}
