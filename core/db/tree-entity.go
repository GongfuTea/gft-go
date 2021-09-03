package db

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/types"
)

type DbTreeEntity struct {
	*types.TreeEntity `bson:",inline"`
	CreatedAt         time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (e *DbTreeEntity) Init() {
	e.TreeEntity.Init()
	fmt.Printf("init DbEntity, %+v\n", e)
	fmt.Printf("init DbEntity, %+v\n", e.Entity)

	e.CreatedAt = time.Now()
}

func NewDbTreeEntity() *DbTreeEntity {
	return &DbTreeEntity{
		TreeEntity: &types.TreeEntity{},
	}
}
