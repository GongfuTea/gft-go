package db

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/types"
)

type IDbEntity interface {
	types.IEntity
	GetCreatedAt() time.Time
}

type DbEntity struct {
	types.Entity `bson:",inline" json:",inline"`
	CreatedAt    time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (e DbEntity) GetCreatedAt() time.Time {
	return e.CreatedAt
}

func (e *DbEntity) Init() {
	e.Entity.Init()

	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}

	fmt.Printf("init DbEntity, %+v\n", e)
}

func NewDbEntity() *DbEntity {
	return &DbEntity{}
}
