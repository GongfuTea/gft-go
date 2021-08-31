package db

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/types"
)

type DbEntity struct {
	*types.Entity `bson:",inline"`
	CreatedAt     time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (e *DbEntity) Init() {
	e.Entity.Init()
	fmt.Printf("init DbEntity, %+v\n", e)

	e.CreatedAt = time.Now()
}

func NewDbEntity() *DbEntity {
	return &DbEntity{
		Entity: &types.Entity{},
	}
}
