package db

import (
	"fmt"
	"time"

	"github.com/GongfuTea/gft-go/types"
)

type IDbTreeEntity interface {
	types.ITreeEntity
	GetCreatedAt() time.Time
}

type DbTreeEntity struct {
	types.TreeEntity `bson:",inline" json:",inline"`
	CreatedAt        time.Time `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
}

func (e *DbTreeEntity) Init() {
	e.TreeEntity.Init()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = time.Now()
	}

	fmt.Printf("init DbTreeEntity, %+v\n", e)
	fmt.Printf("init DbTreeEntity, %+v\n", e.Entity)

}

func (e DbTreeEntity) GetCreatedAt() time.Time {
	return e.CreatedAt
}

// func NewDbTreeEntity() *DbTreeEntity {
// 	return &DbTreeEntity{
// 		TreeEntity: &types.TreeEntity{
// 			Entity: &types.Entity{},
// 		},
// 	}
// }
